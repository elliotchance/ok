package compiler

import (
	"path"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

// Compile will return the compiled file. If there are any dependent packages
// they will also be compiled.
func Compile(rootPath, pkgPath string, includeTests bool, anonFunctionName int) (*vm.File, []error) {
	p := parser.NewParser(anonFunctionName)
	p.ParseDirectory(path.Join(rootPath, pkgPath), includeTests)
	if errs := p.Errors(); len(errs) > 0 {
		return nil, errs
	}

	packageName := util.PackageNameFromPath(rootPath, pkgPath)

	funcs := map[string]*ast.Func{}
	imports := map[string]map[string]*types.Type{}

	for _, pkgName := range p.Imports() {
		imports[pkgName] = map[string]*types.Type{}

		// TODO(elliot): Check import location exists.

		if p, ok := vm.Packages[pkgName]; ok {
			for fnName, fn := range p.FuncDefs {
				imports[pkgName][fnName] = fn.Type()
			}
		} else {
			subFile, errs := Compile(rootPath, pkgName, false, anonFunctionName+10000)
			if len(errs) > 0 {
				return nil, errs
			}

			for fnName, fn := range subFile.FuncDefs {
				imports[pkgName][fnName] = fn.Type()
			}
		}
	}

	for fnName, fn := range p.Funcs() {
		funcs[fnName] = fn
	}

	okcFile, err := compile(funcs, p.Tests(), p.Constants, imports)
	if err != nil {
		return nil, []error{err}
	}

	packageAlias := strings.ReplaceAll(packageName, "/", "__")
	compiledPackageFn, err := CompileFunc(p.Package(packageAlias), &vm.File{
		Funcs:     map[string]*vm.CompiledFunc{},
		FuncDefs:  funcs,
		Constants: p.Constants,
		Imports:   imports,
	})
	if err != nil {
		return nil, []error{err}
	}
	okcFile.PackageFunc = compiledPackageFn

	for _, funcDef := range okcFile.FuncDefs {
		// We don't need to serialize this.
		funcDef.Statements = nil
	}

	for _, fn := range okcFile.Funcs {
		// We don't need to serialize this.
		fn.Type = nil
	}

	err = vm.Store(okcFile, packageName)
	if err != nil {
		return nil, []error{err}
	}

	return okcFile, nil
}
