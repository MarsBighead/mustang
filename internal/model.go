package internal

import "time"

// File processing status.
const (
	StatusAnalyze = 0 // scanned, MD5 pending
	StatusDedup   = 1 // dedup processed
)

// FileInfo represents a file's metadata.
type FileInfo struct {
	Path    string // relative to run root
	Size    int64
	MD5     string
	ModTime time.Time
	Inode   uint64 // filesystem inode; files sharing an inode are hard links
	RunID   int64
	Status  int
}

// DuplicateGroup represents a group of files sharing the same content.
type DuplicateGroup struct {
	MD5   string
	Files []FileInfo
}

// Size returns the size of each file in the group.
func (g DuplicateGroup) Size() int64 {
	if len(g.Files) > 0 {
		return g.Files[0].Size
	}
	return 0
}

// UniqueInodes returns the number of distinct inodes in the group.
func (g DuplicateGroup) UniqueInodes() int {
	seen := make(map[uint64]struct{})
	for _, f := range g.Files {
		seen[f.Inode] = struct{}{}
	}
	return len(seen)
}

// WastedBytes returns bytes wasted by true duplicates (excluding hard links).
func (g DuplicateGroup) WastedBytes() int64 {
	u := g.UniqueInodes()
	if u <= 1 {
		return 0
	}
	return g.Size() * int64(u-1)
}
