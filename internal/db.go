package internal

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
	mu   sync.Mutex
}

func OpenDB(path string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	conn, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	db := &DB{conn: conn}
	if err := db.init(); err != nil {
		conn.Close()
		return nil, err
	}
	return db, nil
}

func (d *DB) init() error {
	_, err := d.conn.Exec(`
		CREATE TABLE IF NOT EXISTS files (
			path    TEXT NOT NULL,
			size    INTEGER NOT NULL,
			md5     TEXT NOT NULL DEFAULT '',
			modtime INTEGER NOT NULL,
			inode   INTEGER NOT NULL DEFAULT 0,
			run_id  INTEGER NOT NULL DEFAULT 0,
			status  INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY (path, run_id)
		);
		CREATE INDEX IF NOT EXISTS idx_files_md5   ON files(md5);
		CREATE INDEX IF NOT EXISTS idx_files_size  ON files(size);
		CREATE INDEX IF NOT EXISTS idx_files_runid ON files(run_id);

		CREATE TABLE IF NOT EXISTS duplicates (
			run_id  INTEGER NOT NULL,
			md5     TEXT NOT NULL,
			path    TEXT NOT NULL,
			size    INTEGER NOT NULL,
			modtime INTEGER NOT NULL,
			inode   INTEGER NOT NULL,
			keep    INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY (run_id, md5, path)
		);
		CREATE INDEX IF NOT EXISTS idx_dup_md5 ON duplicates(md5);
		CREATE INDEX IF NOT EXISTS idx_dup_runid ON duplicates(run_id);

		CREATE TABLE IF NOT EXISTS meta (
			key   TEXT PRIMARY KEY,
			value TEXT NOT NULL
		);
	`)
	return err
}

func (d *DB) Close() error { return d.conn.Close() }

// ── run management (via meta table) ──────────────────────────────────────────

func (d *DB) NewRun(command string) (int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	now := time.Now().Unix()
	d.conn.Exec("INSERT OR REPLACE INTO meta (key,value) VALUES ('run:command',?)", command)
	d.conn.Exec("INSERT OR REPLACE INTO meta (key,value) VALUES ('run:started',?)", now)
	d.conn.Exec("DELETE FROM meta WHERE key LIKE 'run:root:%'")
	d.conn.Exec("DELETE FROM meta WHERE key='run:finished'")

	var id int64
	d.conn.QueryRow("SELECT COALESCE(MAX(run_id),0) FROM files").Scan(&id)
	id++
	d.conn.Exec("INSERT OR REPLACE INTO meta (key,value) VALUES ('run:id',?)", id)
	return id, nil
}

func (d *DB) FinishRun() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.conn.Exec("INSERT OR REPLACE INTO meta (key,value) VALUES ('run:finished',?)", time.Now().Unix())
}

func (d *DB) GetCurrentRunID() int64 {
	var id int64
	d.conn.QueryRow("SELECT CAST(value AS INTEGER) FROM meta WHERE key='run:id'").Scan(&id)
	return id
}

func (d *DB) SetRunRoots(roots []string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.conn.Exec("DELETE FROM meta WHERE key LIKE 'run:root:%'")
	for i, r := range roots {
		d.conn.Exec("INSERT INTO meta (key,value) VALUES (?,?)", fmt.Sprintf("run:root:%d", i), r)
	}
}

func (d *DB) GetRunRoots() []string {
	rows, err := d.conn.Query("SELECT value FROM meta WHERE key LIKE 'run:root:%' ORDER BY key")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var roots []string
	for rows.Next() {
		var v string
		rows.Scan(&v)
		roots = append(roots, v)
	}
	return roots
}

func (d *DB) GetRunStarted() time.Time {
	var ts int64
	d.conn.QueryRow("SELECT CAST(value AS INTEGER) FROM meta WHERE key='run:started'").Scan(&ts)
	if ts == 0 {
		return time.Time{}
	}
	return time.Unix(ts, 0)
}

func (d *DB) GetRunFinished() time.Time {
	var ts int64
	d.conn.QueryRow("SELECT CAST(value AS INTEGER) FROM meta WHERE key='run:finished'").Scan(&ts)
	if ts == 0 {
		return time.Time{}
	}
	return time.Unix(ts, 0)
}

func (d *DB) GetRunCommand() string {
	var cmd string
	d.conn.QueryRow("SELECT value FROM meta WHERE key='run:command'").Scan(&cmd)
	return cmd
}

// ── file CRUD ────────────────────────────────────────────────────────────────

func (d *DB) AllEntries(runID int64) ([]FileInfo, error) {
	rows, err := d.conn.Query("SELECT path,size,md5,modtime,inode,run_id,status FROM files WHERE run_id=?", runID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []FileInfo
	for rows.Next() {
		var f FileInfo
		var mt int64
		if err := rows.Scan(&f.Path, &f.Size, &f.MD5, &mt, &f.Inode, &f.RunID, &f.Status); err != nil {
			return nil, err
		}
		f.ModTime = time.Unix(mt, 0)
		out = append(out, f)
	}
	return out, rows.Err()
}

func (d *DB) InsertFile(runID int64, relPath string, size int64, modTime time.Time, inode uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec(
		"INSERT OR IGNORE INTO files (path,size,md5,modtime,inode,run_id,status) VALUES (?,?,?,?,?,?,?)",
		relPath, size, "", modTime.Unix(), inode, runID, StatusAnalyze)
	return err
}

func (d *DB) ClearMD5(runID int64, relPath string, size int64, modTime time.Time, inode uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec(
		"UPDATE files SET size=?, md5='', modtime=?, inode=?, status=? WHERE path=? AND run_id=?",
		size, modTime.Unix(), inode, StatusAnalyze, relPath, runID)
	return err
}

func (d *DB) UpdateMD5(runID int64, relPath, md5val string, size int64, modTime time.Time, inode uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec(
		"UPDATE files SET md5=?, size=?, modtime=?, inode=? WHERE path=? AND run_id=?",
		md5val, size, modTime.Unix(), inode, relPath, runID)
	return err
}

func (d *DB) MarkUnique(runID int64, relPath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec("UPDATE files SET md5='unique' WHERE path=? AND run_id=?", relPath, runID)
	return err
}

func (d *DB) MarkDedupDone(runID int64, relPaths []string) error {
	if len(relPaths) == 0 {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare("UPDATE files SET status=? WHERE path=? AND run_id=?")
	for _, p := range relPaths {
		stmt.Exec(StatusDedup, p, runID)
	}
	return tx.Commit()
}

func (d *DB) DeleteByPath(runID int64, relPath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec("DELETE FROM files WHERE path=? AND run_id=?", relPath, runID)
	return err
}

func (d *DB) DeleteByPaths(runID int64, paths []string) error {
	if len(paths) == 0 {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare("DELETE FROM files WHERE path=? AND run_id=?")
	for _, p := range paths {
		stmt.Exec(p, runID)
	}
	return tx.Commit()
}

// DeleteOrphanPaths removes entries for the given run whose paths are NOT in keep.
func (d *DB) DeleteOrphanPaths(runID int64, keep map[string]bool) (int, error) {
	all, err := d.AllEntries(runID)
	if err != nil {
		return 0, err
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	deleted := 0
	for _, f := range all {
		if !keep[f.Path] {
			d.conn.Exec("DELETE FROM files WHERE path=? AND run_id=?", f.Path, runID)
			deleted++
		}
	}
	return deleted, nil
}

// ── grouping & queries ───────────────────────────────────────────────────────

func (d *DB) ComputeGrouped(runID int64, paths []string) (map[int64][]FileInfo, error) {
	if len(paths) == 0 {
		return nil, nil
	}

	// First, find all duplicate sizes for the entire run (no path filtering needed)
	dupRows, err := d.conn.Query(
		"SELECT size FROM files WHERE run_id=? GROUP BY size HAVING COUNT(*)>1",
		runID)
	if err != nil {
		return nil, err
	}
	dupSizes := make(map[int64]bool)
	for dupRows.Next() {
		var s int64
		dupRows.Scan(&s)
		dupSizes[s] = true
	}
	dupRows.Close()

	if len(dupSizes) == 0 {
		return nil, nil
	}

	// Now, query files in batches to find which new paths have duplicate sizes
	groups := make(map[int64][]FileInfo)
	const batchSize = 500

	for i := 0; i < len(paths); i += batchSize {
		end := i + batchSize
		if end > len(paths) {
			end = len(paths)
		}
		batch := paths[i:end]

		ph := "?" + strings.Repeat(",?", len(batch)-1)
		args := make([]any, len(batch))
		for j, p := range batch {
			args[j] = p
		}

		rows, err := d.conn.Query(
			"SELECT path,size,md5,modtime,inode,run_id,status FROM files WHERE path IN ("+ph+") AND run_id=?",
			append(args, runID)...)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var f FileInfo
			var mt int64
			if err := rows.Scan(&f.Path, &f.Size, &f.MD5, &mt, &f.Inode, &f.RunID, &f.Status); err != nil {
				rows.Close()
				return nil, err
			}
			f.ModTime = time.Unix(mt, 0)
			// Only include files whose size appears more than once
			if dupSizes[f.Size] {
				groups[f.Size] = append(groups[f.Size], f)
			}
		}
		rows.Close()
	}

	return groups, nil
}

func (d *DB) FindDuplicates(runID int64) ([]DuplicateGroup, error) {
	rows, err := d.conn.Query(`
		SELECT md5, path, size, modtime, inode, run_id, status
		FROM files
		WHERE run_id = ? AND md5 != '' AND md5 != 'unique'
		  AND md5 IN (
			SELECT md5 FROM files WHERE run_id=? AND md5 != '' AND md5 != 'unique'
			GROUP BY md5 HAVING COUNT(DISTINCT inode) > 1
		  )
		ORDER BY md5, modtime`, runID, runID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		groups []DuplicateGroup
		cur    *DuplicateGroup
	)
	for rows.Next() {
		var (
			md5v, path string
			size       int64
			mt         int64
			inode      uint64
			rid        int64
			status     int
		)
		if err := rows.Scan(&md5v, &path, &size, &mt, &inode, &rid, &status); err != nil {
			return nil, err
		}
		if cur == nil || cur.MD5 != md5v {
			groups = append(groups, DuplicateGroup{MD5: md5v})
			cur = &groups[len(groups)-1]
		}
		cur.Files = append(cur.Files, FileInfo{
			Path: path, Size: size, MD5: md5v,
			ModTime: time.Unix(mt, 0), Inode: inode, RunID: rid, Status: status,
		})
	}
	return groups, rows.Err()
}

type DBStats struct {
	TotalFiles  int64
	TotalSize   int64
	HashedFiles int64
	DupGroups   int64
	WastedFiles int64
	WastedBytes int64
}

func (d *DB) Stats(runID int64) (*DBStats, error) {
	s := &DBStats{}
	d.conn.QueryRow("SELECT COUNT(*), COALESCE(SUM(size),0) FROM files WHERE run_id=?", runID).
		Scan(&s.TotalFiles, &s.TotalSize)
	d.conn.QueryRow("SELECT COUNT(*) FROM files WHERE run_id=? AND md5 != '' AND md5 != 'unique'", runID).
		Scan(&s.HashedFiles)
	d.conn.QueryRow(`SELECT COUNT(*) FROM (
		SELECT md5 FROM files WHERE run_id=? AND md5 != '' AND md5 != 'unique'
		GROUP BY md5 HAVING COUNT(DISTINCT inode)>1)`, runID).Scan(&s.DupGroups)
	d.conn.QueryRow(`SELECT COALESCE(SUM(wasted_cnt),0), COALESCE(SUM(wasted_bytes),0) FROM (
		SELECT (COUNT(DISTINCT inode)-1) as wasted_cnt, size*(COUNT(DISTINCT inode)-1) as wasted_bytes
		FROM files WHERE run_id=? AND md5 != '' AND md5 != 'unique'
		GROUP BY md5 HAVING COUNT(DISTINCT inode) > 1)`, runID).
		Scan(&s.WastedFiles, &s.WastedBytes)
	return s, nil
}

func (d *DB) TopDuplicateMD5s(runID int64, limit int) ([]string, error) {
	rows, err := d.conn.Query(`
		SELECT md5 FROM files
		WHERE run_id=? AND md5 != '' AND md5 != 'unique'
		GROUP BY md5 HAVING COUNT(DISTINCT inode) > 1
		ORDER BY size * (COUNT(DISTINCT inode) - 1) DESC
		LIMIT ?`, runID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var md5s []string
	for rows.Next() {
		var m string
		rows.Scan(&m)
		md5s = append(md5s, m)
	}
	return md5s, rows.Err()
}

func (d *DB) FilesByMD5(runID int64, md5val string) ([]FileInfo, error) {
	rows, err := d.conn.Query(
		"SELECT path,size,md5,modtime,inode,run_id,status FROM files WHERE run_id=? AND md5=? ORDER BY modtime",
		runID, md5val)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var files []FileInfo
	for rows.Next() {
		var f FileInfo
		var mt int64
		if err := rows.Scan(&f.Path, &f.Size, &f.MD5, &mt, &f.Inode, &f.RunID, &f.Status); err != nil {
			return nil, err
		}
		f.ModTime = time.Unix(mt, 0)
		files = append(files, f)
	}
	return files, rows.Err()
}

func (d *DB) ClearFiles(runID int64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec("DELETE FROM files WHERE run_id=?", runID)
	return err
}

// ── duplicates table ─────────────────────────────────────────────────────────

func (d *DB) ClearDuplicates(runID int64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.conn.Exec("DELETE FROM duplicates WHERE run_id=?", runID)
	return err
}

func (d *DB) InsertDuplicates(runID int64, groups []DuplicateGroup, strategy string) error {
	if len(groups) == 0 {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare("INSERT INTO duplicates (run_id,md5,path,size,modtime,inode,keep) VALUES (?,?,?,?,?,?,?)")

	for _, g := range groups {
		keep, remove := pickFilesForDup(g, strategy)
		// Insert the kept file
		stmt.Exec(runID, g.MD5, keep.Path, keep.Size, keep.ModTime.Unix(), keep.Inode, 1)
		// Insert files to remove
		for _, f := range remove {
			stmt.Exec(runID, g.MD5, f.Path, f.Size, f.ModTime.Unix(), f.Inode, 0)
		}
	}
	return tx.Commit()
}

func (d *DB) GetDuplicates(runID int64) ([]DuplicateGroup, error) {
	rows, err := d.conn.Query(`
		SELECT md5, path, size, modtime, inode, keep
		FROM duplicates
		WHERE run_id = ?
		ORDER BY md5, modtime`, runID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		groups []DuplicateGroup
		cur    *DuplicateGroup
	)
	for rows.Next() {
		var (
			md5v, path string
			size       int64
			mt         int64
			inode      uint64
			keep       int
		)
		if err := rows.Scan(&md5v, &path, &size, &mt, &inode, &keep); err != nil {
			return nil, err
		}
		if cur == nil || cur.MD5 != md5v {
			groups = append(groups, DuplicateGroup{MD5: md5v})
			cur = &groups[len(groups)-1]
		}
		cur.Files = append(cur.Files, FileInfo{
			Path:    path,
			Size:    size,
			MD5:     md5v,
			ModTime: time.Unix(mt, 0),
			Inode:   inode,
			RunID:   runID,
			Status:  keep, // reuse Status field to indicate keep(1) or remove(0)
		})
	}
	return groups, rows.Err()
}

func (d *DB) DeleteDuplicatesByPaths(runID int64, paths []string) error {
	if len(paths) == 0 {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare("DELETE FROM duplicates WHERE run_id=? AND path=?")
	for _, p := range paths {
		stmt.Exec(runID, p)
	}
	return tx.Commit()
}

// pickFilesForDup separates kept file from files to remove based on strategy.
func pickFilesForDup(g DuplicateGroup, strategy string) (keep FileInfo, remove []FileInfo) {
	sorted := make([]FileInfo, len(g.Files))
	copy(sorted, g.Files)
	switch strategy {
	case "newest":
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].ModTime.After(sorted[j].ModTime)
		})
	default:
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].ModTime.Before(sorted[j].ModTime)
		})
	}
	return sorted[0], sorted[1:]
}
