package imagefile

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

// Mover is an implementation of the Organizer interface
type Mover struct {
	Files  []FileInfo
	Source string
	Dest   string
}

// NewMover returns a pointer to an intialized Mover
func NewMover(files []FileInfo, src, dst string) *Mover {
	return &Mover{files, src, dst}
}

// Execute moves a file from the source directory to the destionation directory
func (m *Mover) Execute(dryRun bool) error {
	for _, f := range m.Files {
		destDir := m.getDestDir(f.DateTime)

		// create the directory if it doesn't already exist
		if err := m.makeDir(destDir, dryRun); err != nil {
			return errors.Wrapf(err, "Failed to create directory for file '%s'", f.Path)
		}

		to := filepath.Join(
			m.getDestDir(f.DateTime),
			f.Name,
		)

		if dryRun {
			fmt.Printf("[DRY RUN]  Would have run os.Rename(%s, %s)\n", f.Path, to)
			continue
		}

		if err := os.Rename(f.Path, to); err != nil {
			return errors.Wrapf(err, "Failed to move %s to %s", f.Name, m.Dest)
		}
	}

	return nil
}

func (m *Mover) makeDir(dir string, dryRun bool) error {
	// if it already exists, bail early
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return nil
	}

	info, err := os.Lstat(m.Dest)
	if err != nil {
		return err
	}
	mode := info.Mode()

	if dryRun {
		// TODO: make a DryPrinter for consistent dry run output
		fmt.Printf("[DRY RUN]  Would have run MkdirAll(%s, %d)\n", dir, mode)
		return nil
	}

	if err := os.MkdirAll(dir, mode); err != nil {
		return err
	}

	return nil
}

func (m *Mover) getDestDir(dateTime time.Time) string {
	d := filepath.Join(
		m.Dest,
		fmt.Sprintf("%d", dateTime.Year()),
		fmt.Sprintf("%02d", dateTime.Month()),
	)

	return d
}
