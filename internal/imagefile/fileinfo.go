package imagefile

import "time"

// FileInfo contains information and metadata about a file
type FileInfo struct {
	Name     string
	Path     string
	DateTime time.Time
}

// ByDate is a FileInfo sorter
type ByDate []FileInfo

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDate) Less(i, j int) bool { return d[i].DateTime.Before(d[j].DateTime) }
