package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageNameFromPath(t *testing.T) {
	for testName, test := range map[string]struct {
		baseDir     string
		packagePath string
		expected    string
	}{
		"1": {"/", "/foo", "foo"},
		"2": {"/foo", "/foo/bar", "bar"},
		"3": {"/", "/foo/bar", "foo/bar"},
		"4": {"/foo/bar", "/foo/bar", "bar"},
		"5": {"/", "/", "/"},
		"6": {"/", "/foo/", "foo"},
		"7": {"/foo", "/foo/bar/", "bar"},
		"8": {"/foo/", "/foo/bar", "bar"},
		"9": {"/foo/", "/foo/bar/", "bar"},
	} {
		t.Run(testName, func(t *testing.T) {
			actual := PackageNameFromPath(test.baseDir, test.packagePath)
			assert.Equal(t, test.expected, actual)
		})
	}
}
