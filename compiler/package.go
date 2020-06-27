package compiler

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
)

func CompilePackage(dir string, includeTests bool) (*Compiled, error) {
	// Step 1: Find all the files that need to be compiled.
	fileNames, err := getAllOKFilesInPath(dir, includeTests)
	if err != nil {
		return nil, err
	}

	dir = path.Clean(dir)

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
			if fileNameDir := path.Clean(path.Dir(fileName)); fileNameDir != dir {
				pkg := filepath.Base(fileNameDir)
				funcs[pkg+"."+name] = fn
			} else {
				funcs[name] = fn
			}
		}

		tests = append(tests, p.File.Tests...)

		for pkg := range p.File.Imports {
			// TODO(elliot): Check import location exists.

			// Ignore builtin libraries as they are already provided to the VM
			// through vm/lib.go. See Makefile.
			//
			// TODO(elliot): This needs to be smarter.
			if pkg == "math" {
				continue
			}

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