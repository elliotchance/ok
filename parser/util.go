package parser

import (
	"io/ioutil"
	"path"
)

// getAllOKFilesInPath non-recursively returns a list of OK files.
func getAllOKFilesInPath(dir string, includeTests bool) ([]string, error) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, f := range fs {
		fileName := path.Join(dir, f.Name())
		if path.Ext(fileName) == ".ok" {
			files = append(files, fileName)
		}

		if includeTests && path.Ext(fileName) == ".okt" {
			files = append(files, fileName)
		}
	}

	return files, nil
}
