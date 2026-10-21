//go:build !freebsd

package xattr

import "golang.org/x/sys/unix"

const errNoAttr = unix.ENODATA
