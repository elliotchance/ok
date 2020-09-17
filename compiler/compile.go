package compiler

import (
	"io/ioutil"
	"path"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

// Compile will return the compiled file. If there are any dependent packages
// they will also be compiled.
func Compile(rootPath, pkgPath string, includeTests bool) (*vm.File, []error) {
	packageName := util.PackageNameFromPath(rootPath, pkgPath)

	fileNames, err := util.GetAllOKFilesInPath(path.Join(rootPath, pkgPath), includeTests)
	if err != nil {
		return nil, []error{err}
	}

	funcs := map[string]*ast.Func{}
	imports := map[string]struct{}{}
	importedFuncs := map[string]*ast.Func{}
	interfaces := map[string]map[string]*types.Type{}
	var tests []*ast.Test
	constants := map[string]*ast.Literal{}

	for _, fileName := range fileNames {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, []error{err}
		}

		p := parser.ParseString(string(data), fileName)
		if errs := p.Errors(); len(errs) > 0 {
			return nil, errs
		}

		for pkg := range p.File.Imports {
			imports[pkg] = struct{}{}

			// TODO(elliot): Check import location exists.

			if p, ok := vm.Packages[pkg]; ok {
				for fnName, fn := range p.FuncDefs {
					importedFuncs[pkg+"."+fnName] = fn
				}
			} else {
				subFile, errs := Compile(rootPath, pkg, false)
				if len(errs) > 0 {
					return nil, errs
				}

				pkgVariable := path.Base(pkg)
				for fnName, fn := range subFile.FuncDefs {
					importedFuncs[pkgVariable+"."+fnName] = fn
				}
			}
		}

		for fnName, fn := range p.File.Funcs {
			funcs[fnName] = fn
		}

		for key, i := range p.Interfaces {
			interfaces[key] = i
		}

		for key, i := range p.Constants {
			constants[key] = i
		}

		tests = append(tests, p.File.Tests...)
	}

	okcFile, err := compile(funcs, importedFuncs, tests, interfaces, constants)
	if err != nil {
		return nil, []error{err}
	}
	okcFile.Imports = imports

	for _, funcDef := range okcFile.FuncDefs {
		// We don't need to serialize this.
		funcDef.Statements = nil
	}

	err = vm.Store(okcFile, packageName)
	if err != nil {
		return nil, []error{err}
	}

	return okcFile, nil
}
