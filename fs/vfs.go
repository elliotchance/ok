package fs

import (
	"github.com/blang/vfs/mountfs"
)

// Filesystem wraps the OS field system, but allow the stdlib to be mounted. See
// lib-gen.
var Filesystem *mountfs.MountFS
