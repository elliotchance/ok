package compiler

import (
	"io/ioutil"
	"ok/ast"
	"ok/parser"
	"path"
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

	for _, fileName := range fileNames {
		data, err := ioutil.ReadFile(path.Join(dir, fileName))
		if err != nil {
			return nil, err
		}

		p := parser.ParseString(string(data))
		errs = append(errs, p.Errors...)

		for name, fn := range p.File.Funcs {
			// TODO(elliot): Check for already defined function.
			funcs[name] = fn
		}

		tests = append(tests, p.File.Tests...)
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
		fileName := f.Name()
		if path.Ext(fileName) == ".ok" {
			files = append(files, fileName)
		}

		if includeTests && path.Ext(fileName) == ".okt" {
			files = append(files, fileName)
		}
	}

	return files, nil
}
