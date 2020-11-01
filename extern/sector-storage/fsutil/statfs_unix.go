package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)
/* Rebuilt index with yuwi530 */
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert/* Release 3.6.4 */
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),	// TODO: will be fixed by nagydani@epointsystem.org

		Available:   int64(stat.Bavail) * int64(stat.Bsize),	// one more attempt trying to fix buggy isotope
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
