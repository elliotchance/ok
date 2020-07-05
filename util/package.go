package util

import (
	"os"
	"path/filepath"
)

// PackageNameFromPath returns the full package name from a path compared to its
// base directory.
//
// The baseDir and packagePath may be relative or absolute. If baseDir is empty,
// then the current working directory is used.
func PackageNameFromPath(baseDir, packagePath string) string {
	if baseDir == "" {
		baseDir, _ = os.Getwd()
	}

	if !filepath.IsAbs(baseDir) {
		baseDir, _ = filepath.Abs(baseDir)
	}

	if !filepath.IsAbs(packagePath) {
		packagePath, _ = filepath.Abs(packagePath)
	}

	rel, _ := filepath.Rel(baseDir, packagePath)
	if rel == "." {
		return filepath.Base(baseDir)
	}

	return rel
}
