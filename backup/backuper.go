package backup

import (
	"archive/tar"

	"github.com/Noah-Huppert/golog"
)

// Backuper performs the action of backing up a file.
type Backuper interface {
	// Backup files to w. Returns the number of files which were backed up.
	Backup(logger golog.Logger, w *tar.Writer) (int, error)
}
