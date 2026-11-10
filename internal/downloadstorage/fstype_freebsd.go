//go:build freebsd

package downloadstorage

import "golang.org/x/sys/unix"

// fsTypeName names no filesystem on FreeBSD: there Statfs_t.Type is a runtime
// vfs index assigned per boot, not the stable magic number Linux publishes, and
// x/sys wraps neither getvfsbyname nor getvfsstat to turn it into a name.
// Answering "" keeps a box's index from being read as a filesystem name; the
// cost is that ephemeralFS cannot recognise tmpfs on FreeBSD.
func fsTypeName(*unix.Statfs_t) string { return "" }
