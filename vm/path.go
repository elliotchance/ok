package vm

import (
	"path"
	"strings"
)

// TODO(elliot): This should be configurable through an environment variable.
const Directory = "/tmp/okc"

// PathForPackage returns the absolute path of where the okc file should exist
// for a package name. However, the path returned may not exist.
func PathForPackage(packageName string) string {
	return path.Join(Directory,
		strings.ReplaceAll(packageName, "/", "__")+".okc")
}
