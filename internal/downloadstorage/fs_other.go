//go:build !linux && !darwin && !freebsd

package downloadstorage

import "errors"

// Prepared downloads run on Linux, macOS and FreeBSD hosts; this keeps the
// package building elsewhere.
func statFS(string) (fsStats, error) { return fsStats{}, errors.ErrUnsupported }

func deviceOf(string) (uint64, error) { return 0, errors.ErrUnsupported }
