package parser

import (
	"path"

	"github.com/elliotchance/ok/fs"
)

// getAllOKFilesInPath non-recursively returns a list of OK files.
func getAllOKFilesInPath(dir string, includeTests bool) ([]string, error) {
	files, err := fs.Filesystem.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, f := range files {
		fileName := path.Join(dir, f.Name())
		if path.Ext(fileName) == ".ok" {
			fileNames = append(fileNames, fileName)
		}

		if includeTests && path.Ext(fileName) == ".okt" {
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames, nil
}
