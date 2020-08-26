package compiler

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

func CompilePackage(dir string, includeTests bool) (*Compiled, []error) {
	// Step 1: Find all the files that need to be compiled.
	fileNames, err := util.GetAllOKFilesInPath(dir, includeTests)
	if err != nil {
		return nil, []error{err}
	}

	dir = path.Clean(dir)

	// Step 2: Parse all files and combine.
	var errs []error
	funcs := map[string]*ast.Func{}
	var tests []*ast.Test
	interfaces := map[string]map[string]string{}
	constants := map[string]*ast.Literal{}

	for len(fileNames) > 0 {
		fileName := fileNames[0]
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, []error{err}
		}

		p := parser.ParseString(string(data), fileName)
		errs = append(errs, p.Errors()...)

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

		for key, i := range p.Interfaces {
			interfaces[key] = i
		}

		for key, i := range p.Constants {
			constants[key] = i
		}

		for pkg := range p.File.Imports {
			// TODO(elliot): Check import location exists.

			// Ignore builtin libraries as they are already provided to the VM
			// through vm/lib.go. See Makefile.
			if vm.Packages[pkg] {
				continue
			}

			newFileNames, err := util.GetAllOKFilesInPath(path.Join(dir, pkg), false)
			if err != nil {
				return nil, []error{err}
			}

			fileNames = append(fileNames, newFileNames...)
		}

		fileNames = fileNames[1:]
	}

	if len(errs) > 0 {
		return nil, errs
	}

	// Step 3: Compile everything all at once.
	compiled, err := compile(funcs, tests, interfaces, constants)
	if err != nil {
		return nil, []error{err}
	}

	return compiled, nil
}
