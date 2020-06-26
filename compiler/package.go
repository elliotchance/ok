package compiler

import (
	"io/ioutil"
	"ok/ast"
	"ok/parser"
	"path"
	"path/filepath"
)

func CompilePackage(dir string, includeTests bool) (*Compiled, error) {
	// Step 1: Find all the files that need to be compiled.
	fileNames, err := getAllOKFilesInPath(dir, includeTests)
	if err != nil {
		return nil, err
	}

	// Step 2: Parse all files and combine.
	var errs []error
	funcs := map[string]*ast.Func{}
	var tests []*ast.Test

	for len(fileNames) > 0 {
		fileName := fileNames[0]
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}

		p := parser.ParseString(string(data))
		errs = append(errs, p.Errors...)

		for name, fn := range p.File.Funcs {
			// TODO(elliot): Check for already defined function.
			if path.Dir(fileName) != dir {
				pkg := filepath.Base(path.Dir(fileName))
				funcs[pkg+"."+name] = fn
			} else {
				funcs[name] = fn
			}
		}

		tests = append(tests, p.File.Tests...)

		for pkg := range p.File.Imports {
			// TODO(elliot): Check import location exists.
			newFileNames, err := getAllOKFilesInPath(path.Join(dir, pkg), false)
			if err != nil {
				return nil, err
			}

			fileNames = append(fileNames, newFileNames...)
		}

		fileNames = fileNames[1:]
	}

	// Step 3: Compile everything all at once.
	return compile(funcs, tests)
}

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
