package nexe

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
)

func TranspilePackage(rootPath, pkgPath string, anonFunctionName *int) (Stmt, []error) {
	*anonFunctionName += 10000
	p := parser.NewParser(*anonFunctionName)

	// TODO(elliot): This will prevent import from legitimately importing any
	//  packages that exist as a root. We need to keep track of the real stdlib
	//  imports.
	dir := path.Join(rootPath, pkgPath)
	if !strings.Contains(pkgPath, "/") {
		dir = "/" + pkgPath
	}

	p.ParseDirectory(dir, false)
	if errs := p.Errors(); len(errs) > 0 {
		return nil, errs
	}

	scope := NewScope()
	scope.Add("print", types.NewFunc(nil, nil))
	scope.Add("__pow", types.NewFunc(nil, []*types.Type{types.Number}))
	scope.Add("__log", types.NewFunc(nil, []*types.Type{types.Number}))
	scope.Add("__rand", types.NewFunc(nil, []*types.Type{types.Number}))

	var r Stmts

	for importVariable, importPath := range p.Imports() {
		pkgStmt, err := TranspilePackage(rootPath, importPath, anonFunctionName)
		if err != nil {
			return nil, err
		}

		scope.Add(importVariable, types.Any)
		r = append(r, pkgStmt)
	}

	scope = scope.PackageScope(PackageFunctionName(pkgPath))
	for _, fn := range p.Funcs() {
		scope.Add(fn.Name, fn.Type())
	}
	for name, ty := range p.Constants {
		scope.Add(name, ty.Kind)
	}

	compiled := NewObject(types.AnyMap)
	for _, fn := range p.Funcs() {
		transpiledFn, _, err := TranspileFunc(fn, scope)
		if err != nil {
			return nil, []error{err}
		}

		compiled.Elements[Expr(transpiledFn.Name)] = transpiledFn
	}

	tys := NewObject(types.AnyMap)
	cwd, _ := os.Getwd()
	for i, fn := range scope.funcs {
		tys.Elements[Expr(fmt.Sprintf("_fn%d", i))] = Expr(
			fmt.Sprintf("[\"%s\", \"%s\", \"%s\"]", fn.Name, fn.Type.String(),
				strings.TrimPrefix(fn.Pos, cwd)))
	}
	compiled.Elements[Expr("$fns")] = tys

	r = append(r, &Assign{
		Name:       PackageFunctionName(pkgPath),
		Expr:       compiled,
		IncludeLet: true,
	})

	return r, nil
}

func PackageFunctionName(path string) string {
	path = strings.ReplaceAll(path, "/", "__")
	path = strings.ReplaceAll(path, "-", "_")

	return path
}
