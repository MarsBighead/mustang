package internal

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

// ScanResult holds the outcome of a directory scan.
type ScanResult struct {
	TotalScanned int
	Skipped      int
	Inserted     int
	Updated      int
	Deleted      int
	ChangedPaths []string // relative paths needing MD5 recomputation
	NewPaths     []string // relative paths that are size-duplicate candidates
}

// Scanner walks directories and detects file changes.
type Scanner struct {
	db        *DB
	runID     int64
	roots     []string // absolute root directories scanned
	Workers   int
	Progressf func(string, ...any)
}

func NewScanner(db *DB, runID int64, roots []string) *Scanner {
	return &Scanner{
		db:      db,
		runID:   runID,
		roots:   roots,
		Workers: 4,
	}
}

// toRel converts an absolute path to a relative path using the best matching root.
func (s *Scanner) toRel(absPath string) string {
	for _, root := range s.roots {
		if rel, err := filepath.Rel(root, absPath); err == nil {
			if !strings.HasPrefix(rel, "..") {
				return rel
			}
		}
	}
	return absPath
}

// toAbs converts a relative path back to an absolute path.
func (s *Scanner) toAbs(relPath string) string {
	for _, root := range s.roots {
		candidate := filepath.Join(root, relPath)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	// fallback: try first root
	if len(s.roots) > 0 {
		return filepath.Join(s.roots[0], relPath)
	}
	return relPath
}

// Scan performs an incremental scan of the given directories.
func (s *Scanner) Scan(ctx context.Context) (*ScanResult, error) {
	existing, err := s.db.AllEntries(s.runID)
	if err != nil {
		return nil, fmt.Errorf("load db entries: %w", err)
	}

	existingMap := make(map[string]FileInfo, len(existing))
	for _, e := range existing {
		existingMap[e.Path] = e // key is relative path
	}
	seen := make(map[string]bool) // tracks relative paths seen on disk

	result := &ScanResult{}

	for _, dir := range s.roots {
		err = filepath.WalkDir(dir, func(absPath string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if !d.Type().IsRegular() {
				return nil
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}

			info, err := d.Info()
			if err != nil {
				return nil
			}

			relPath := s.toRel(absPath)
			result.TotalScanned++
			seen[relPath] = true

			if s.Progressf != nil && result.TotalScanned%200 == 0 {
				s.Progressf("Scanned %d files...", result.TotalScanned)
			}

			if ex, ok := existingMap[relPath]; ok {
				if info.Size() == ex.Size && info.ModTime().Unix() == ex.ModTime.Unix() {
					result.Skipped++
					return nil
				}
				if err := s.db.ClearMD5(s.runID, relPath, info.Size(), info.ModTime(), getInode(info)); err != nil {
					return nil
				}
				result.Updated++
				result.ChangedPaths = append(result.ChangedPaths, relPath)
			} else {
				if err := s.db.InsertFile(s.runID, relPath, info.Size(), info.ModTime(), getInode(info)); err != nil {
					return nil
				}
				result.Inserted++
				result.NewPaths = append(result.NewPaths, relPath)
			}
			return nil
		})
		if err != nil {
			return result, fmt.Errorf("walk %s: %w", dir, err)
		}
	}

	// Delete orphaned entries
	for relPath := range existingMap {
		if !seen[relPath] {
			if err := s.db.DeleteByPath(s.runID, relPath); err == nil {
				result.Deleted++
			}
		}
	}

	return result, nil
}

// ComputeMD5s calculates MD5 hashes for the given relative paths concurrently.
func (s *Scanner) ComputeMD5s(ctx context.Context, relPaths []string, onDone func()) error {
	if len(relPaths) == 0 {
		return nil
	}

	ch := make(chan string, len(relPaths))
	for _, p := range relPaths {
		ch <- p
	}
	close(ch)

	workers := s.Workers
	if workers > len(relPaths) {
		workers = len(relPaths)
	}

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for relPath := range ch {
				if ctx.Err() != nil {
					return
				}
				absPath := s.toAbs(relPath)
				hash, err := computeFileMD5(absPath)
				if err == nil {
					info, _ := os.Stat(absPath)
					if info != nil {
						_ = s.db.UpdateMD5(s.runID, relPath, hash, info.Size(), info.ModTime(), getInode(info))
					}
				}
				if onDone != nil {
					onDone()
				}
			}
		}()
	}
	wg.Wait()
	return ctx.Err()
}

func computeFileMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func getInode(info os.FileInfo) uint64 {
	if st, ok := info.Sys().(*syscall.Stat_t); ok {
		return st.Ino
	}
	return 0
}
