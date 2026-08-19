package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"mustang/internal"

	"github.com/spf13/cobra"
)

const defaultDBPath = "~/.xm/xm.db"

var (
	dbPathFlag string
	ctx        context.Context
)

func main() {
	ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)

	rootCmd := &cobra.Command{
		Use:   "xm",
		Short: "Duplicate file detector & cleaner (WeChat oriented)",
		Long: `xm scans directories for duplicate files using content-based (MD5) matching,
with built-in support for WeChat storage cleanup. Files are moved to a quarantine
directory instead of being permanently deleted.`,
	}
	rootCmd.PersistentFlags().StringVar(&dbPathFlag, "db", defaultDBPath, "database path")

	rootCmd.AddCommand(newAnalyzeCmd())
	rootCmd.AddCommand(newStatsCmd())
	rootCmd.AddCommand(newDedupCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// ─── analyze ─────────────────────────────────────────────────────────────────

func newAnalyzeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "analyze [directories...]",
		Short: "Scan directories, compute MD5 hashes, build index",
		Long:  "Walk the target directories, compute MD5 hashes for candidate files, and store results in the database. Defaults to the WeChat storage directory.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirs := args
			if len(dirs) == 0 {
				dirs = defaultWeChatDirs()
			}
			for i := range dirs {
				dirs[i] = expandHome(dirs[i])
			}

			db, err := internal.OpenDB(expandHome(dbPathFlag))
			if err != nil {
				return fmt.Errorf("open db: %w", err)
			}
			defer db.Close()

			// Create a new run record
			runID, err := db.NewRun("analyze")
			if err != nil {
				return fmt.Errorf("new run: %w", err)
			}
			db.SetRunRoots(dirs)

			fmt.Println("Root directories:")
			for _, d := range dirs {
				fmt.Printf("  %s\n", d)
			}

			sc := internal.NewScanner(db, runID, dirs)
			sc.Progressf = func(f string, a ...any) { fmt.Printf("\r"+f, a...) }

			fmt.Println("Scanning...")
			result, err := sc.Scan(ctx)
			if err != nil {
				return fmt.Errorf("scan: %w", err)
			}
			fmt.Printf("\rScanned %d files (new: %d, changed: %d, removed: %d, cached: %d)\n",
				result.TotalScanned, result.Inserted, result.Updated, result.Deleted, result.Skipped)

			// Changed files always need MD5 recomputation.
			needMD5 := append([]string(nil), result.ChangedPaths...)

			// New files only need MD5 if they share a size with another file.
			groups, err := db.ComputeGrouped(runID, result.NewPaths)
			if err != nil {
				return fmt.Errorf("group by size: %w", err)
			}
			var dupCandidates int
			for _, files := range groups {
				for _, f := range files {
					needMD5 = append(needMD5, f.Path)
					dupCandidates++
				}
			}

			// Mark unique-size new files.
			for _, p := range result.NewPaths {
				if !containsPath(groups, p) {
					_ = db.MarkUnique(runID, p)
				}
			}

			if len(needMD5) > 0 {
				fmt.Printf("Computing MD5 for %d files (%d size-duplicate candidates)...\n",
					len(needMD5), dupCandidates)
				var done int
				sc.Progressf = func(f string, a ...any) {
					done++
					fmt.Printf("\rComputing MD5: %d/%d", done, len(needMD5))
				}
				if err := sc.ComputeMD5s(ctx, needMD5, nil); err != nil {
					return fmt.Errorf("compute md5: %w", err)
				}
				fmt.Println()
			}

			db.FinishRun()
			fmt.Println("Analyze complete.")
			return nil
		},
	}
}

// ─── stats ───────────────────────────────────────────────────────────────────

func newStatsCmd() *cobra.Command {
	var (
		topN     int
		strategy string
	)
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Show duplicate statistics from the index",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strategy != "oldest" && strategy != "newest" {
				return fmt.Errorf("invalid strategy %q (use oldest or newest)", strategy)
			}

			db, err := internal.OpenDB(expandHome(dbPathFlag))
			if err != nil {
				return fmt.Errorf("open db: %w", err)
			}
			defer db.Close()

			runID := db.GetCurrentRunID()
			if runID == 0 {
				fmt.Println("No runs recorded. Run 'analyze' first.")
				return nil
			}

			s, err := db.Stats(runID)
			if err != nil {
				return fmt.Errorf("stats: %w", err)
			}
			if s.TotalFiles == 0 {
				fmt.Println("Database is empty. Run 'analyze' first.")
				return nil
			}

			roots := db.GetRunRoots()

			// Run metadata
			fmt.Println("Run info:")
			fmt.Printf("  Command  : %s\n", db.GetRunCommand())
			if started := db.GetRunStarted(); !started.IsZero() {
				fmt.Printf("  Started  : %s\n", started.Format("2006-01-02 15:04:05"))
			}
			if finished := db.GetRunFinished(); !finished.IsZero() {
				fmt.Printf("  Finished : %s\n", finished.Format("2006-01-02 15:04:05"))
			}
			if len(roots) > 0 {
				fmt.Println("  Roots    :")
				for _, d := range roots {
					fmt.Printf("    %s\n", d)
				}
			}

			fmt.Printf("Files indexed : %d\n", s.TotalFiles)
			fmt.Printf("Total size    : %s\n", formatSize(s.TotalSize))
			fmt.Printf("Hashed files  : %d\n", s.HashedFiles)
			fmt.Printf("Dup groups    : %d\n", s.DupGroups)
			fmt.Printf("Wasted files  : %d\n", s.WastedFiles)
			fmt.Printf("Wasted space  : %s\n", formatSize(s.WastedBytes))

			// Find duplicates and populate duplicates table
			groups, err := db.FindDuplicates(runID)
			if err != nil {
				return fmt.Errorf("find duplicates: %w", err)
			}

			// Clear and repopulate duplicates table
			if err := db.ClearDuplicates(runID); err != nil {
				return fmt.Errorf("clear duplicates: %w", err)
			}
			if err := db.InsertDuplicates(runID, groups, strategy); err != nil {
				return fmt.Errorf("insert duplicates: %w", err)
			}

			// Top duplicate groups
			if s.DupGroups > 0 && topN > 0 {
				md5s, err := db.TopDuplicateMD5s(runID, topN)
				if err != nil {
					return fmt.Errorf("top dups: %w", err)
				}
				fmt.Printf("\nTop %d duplicate groups (by wasted space):\n", len(md5s))
				for rank, md5 := range md5s {
					files, _ := db.FilesByMD5(runID, md5)
					if len(files) < 2 {
						continue
					}
					g := internal.DuplicateGroup{MD5: md5, Files: files}
					sz := g.Size()
					unique := g.UniqueInodes()
					links := len(files) - unique
					wasted := g.WastedBytes()
					if links > 0 {
						fmt.Printf("  %d. [%s] %s x %d copies (%d unique, %d hard links), wasted %s\n",
							rank+1, md5[:8], formatSize(sz), len(files), unique, links, formatSize(wasted))
					} else {
						fmt.Printf("  %d. [%s] %s x %d copies, wasted %s\n",
							rank+1, md5[:8], formatSize(sz), len(files), formatSize(wasted))
					}
					for _, f := range files {
						fmt.Printf("       %s\n", f.Path) // already relative
					}
				}
			}

			fmt.Printf("\nDuplicate details saved to database. Run 'dedup' to quarantine duplicates.\n")
			return nil
		},
	}
	cmd.Flags().IntVar(&topN, "top", 5, "number of top duplicate groups to display")
	cmd.Flags().StringVar(&strategy, "keep", "oldest", "which copy to keep: oldest|newest")
	return cmd
}

// ─── dedup ───────────────────────────────────────────────────────────────────

func newDedupCmd() *cobra.Command {
	var (
		qDir   string
		dryRun bool
	)

	cmd := &cobra.Command{
		Use:   "dedup",
		Short: "Quarantine duplicate files marked by stats",
		Long:  "Quarantine duplicate files that were identified by 'stats' command. Files marked for removal are moved to the quarantine directory. Use -dry-run to preview changes without moving files.",
		RunE: func(cmd *cobra.Command, args []string) error {
			qDirExpanded := expandHome(qDir)

			db, err := internal.OpenDB(expandHome(dbPathFlag))
			if err != nil {
				return fmt.Errorf("open db: %w", err)
			}
			defer db.Close()

			runID := db.GetCurrentRunID()
			if runID == 0 {
				fmt.Println("No runs recorded. Run 'analyze' first.")
				return nil
			}

			roots := db.GetRunRoots()
			if len(roots) > 0 {
				fmt.Println("Root directories:")
				for _, d := range roots {
					fmt.Printf("  %s\n", d)
				}
			}

			// Read duplicates from the duplicates table (populated by stats command)
			groups, err := db.GetDuplicates(runID)
			if err != nil {
				return fmt.Errorf("get duplicates: %w", err)
			}
			if len(groups) == 0 {
				fmt.Println("No duplicates found. Run 'stats' first to identify duplicates.")
				return nil
			}

			var (
				totalMoved   int
				totalBytes   int64
				quarantinedR []string // relative paths quarantined
			)

			for _, g := range groups {
				// Files with Status=1 are kept, Status=0 are to be removed
				var keep *internal.FileInfo
				var remove []internal.FileInfo
				for i := range g.Files {
					f := &g.Files[i]
					if f.Status == 1 { // keep
						keep = f
					} else { // remove
						remove = append(remove, *f)
					}
				}
				if keep == nil || len(remove) == 0 {
					continue
				}

				fmt.Printf("\n[%s] %s x %d copies (keep: %s)\n",
					g.MD5[:8], formatSize(g.Size()), len(g.Files), keep.Path)
				fmt.Printf("  + %s\n", keep.Path)
				for _, f := range remove {
					// Skip hard links (same inode as the kept file)
					if f.Inode != 0 && f.Inode == keep.Inode {
						fmt.Printf("  ~ %s (hard link, skipped)\n", f.Path)
						continue
					}
					fmt.Printf("  - %s\n", f.Path)
					totalMoved++
					totalBytes += f.Size
					if dryRun {
						continue
					}
					absPath := resolvePath(f.Path, roots)
					dest, err := quarantineFile(absPath, qDirExpanded)
					if err != nil {
						fmt.Fprintf(os.Stderr, "  ! quarantine failed: %v\n", err)
						totalMoved--
						totalBytes -= f.Size
						continue
					}
					quarantinedR = append(quarantinedR, f.Path)
					fmt.Printf("    -> %s\n", dest)
				}
			}

			if !dryRun && len(quarantinedR) > 0 {
				_ = db.DeleteByPaths(runID, quarantinedR)
				_ = db.DeleteDuplicatesByPaths(runID, quarantinedR)
			}

			fmt.Println()
			action := "Moved"
			if dryRun {
				action = "Would move"
			}
			fmt.Printf("%s %d duplicate files, reclaimed %s\n", action, totalMoved, formatSize(totalBytes))
			return nil
		},
	}

	cmd.Flags().StringVar(&qDir, "quarantine", "~/.xm/quarantine", "quarantine directory")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "preview without moving files")

	return cmd
}

// ─── helpers ─────────────────────────────────────────────────────────────────

func resolvePath(relPath string, roots []string) string {
	for _, root := range roots {
		candidate := filepath.Join(root, relPath)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	if len(roots) > 0 {
		return filepath.Join(roots[0], relPath)
	}
	return relPath
}

func quarantineFile(absPath, quarantineDir string) (string, error) {
	absPath, err := filepath.Abs(absPath)
	if err != nil {
		return "", err
	}
	relPath, err := filepath.Rel("/", absPath)
	if err != nil {
		relPath = filepath.Base(absPath)
	}
	destPath := filepath.Join(quarantineDir, relPath)
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return "", err
	}
	if err := os.Rename(absPath, destPath); err != nil {
		if err := copyFile(absPath, destPath); err != nil {
			return "", err
		}
		if err := os.Remove(absPath); err != nil {
			return "", err
		}
	}
	return destPath, nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.ReadFrom(in)
	return err
}

func defaultWeChatDirs() []string {
	home, _ := os.UserHomeDir()
	base := filepath.Join(home, "Library/Containers/com.tencent.xinWeChat/Data")
	return []string{
		filepath.Join(base, "Library/Application Support/com.tencent.xinWeChat"),
		filepath.Join(base, "Documents"),
	}
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}

func formatSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func containsPath(groups map[int64][]internal.FileInfo, target string) bool {
	for _, files := range groups {
		for _, f := range files {
			if f.Path == target {
				return true
			}
		}
	}
	return false
}
