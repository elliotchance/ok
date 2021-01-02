package compiler

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

// Compile will return the compiled file. If there are any dependent packages
// they will also be compiled.
func Compile(
	rootPath,
	pkgPath string,
	includeTests bool,
	anonFunctionName *int,
	verbose bool,
) (*vm.File, *types.Type, []error) {
	*anonFunctionName += 10000
	p := parser.NewParser(*anonFunctionName)

	// TODO(elliot): This will prevent import from legitimately importing any
	//  packages that exist as a root. We need to keep track of the real stdlib
	//  imports.
	dir := path.Join(rootPath, pkgPath)
	if !strings.Contains(pkgPath, "/") {
		dir = "/" + pkgPath
	}

	p.ParseDirectory(dir, includeTests)
	if errs := p.Errors(); len(errs) > 0 {
		return nil, nil, errs
	}

	packageName := util.PackageNameFromPath(rootPath, pkgPath)
	imports := map[string]*types.Type{}
	var dependencies []*vm.File

	// This is not strictly necessary, but it makes comparing outputs and
	// debugging easier.
	var importNames []string
	for _, pkgName := range p.Imports() {
		importNames = append(importNames, pkgName)
	}
	sort.Strings(importNames)

	for _, pkgName := range importNames {
		// TODO(elliot): Check import location exists.

		subFile, packageFunc, errs := Compile(rootPath, pkgName,
			false, anonFunctionName, verbose)
		if len(errs) > 0 {
			return nil, nil, errs
		}

		imports[pkgName] = packageFunc
		dependencies = append(dependencies, subFile)
	}

	file := &vm.File{
		Types:   types.Registry{},
		Symbols: map[vm.SymbolRegister]*vm.Symbol{},
		Globals: map[string]string{},
	}

	err := p.ResolveTypes(file.Types, imports)
	if err != nil {
		return nil, nil, []error{err}
	}

	packageAlias := strings.ReplaceAll(packageName, "/", "__")

	// Functions at the package level are treated as constants so that a
	// package-level function is callable from any depth. It is also prevents
	// variables or reassignment of these as well.
	for _, fn := range p.Funcs() {
		if fn.Name != "" {
			p.Constants[fn.Name] = &ast.Literal{
				Kind:  fn.Type(),
				Value: fn.UniqueName,
				Pos:   fn.Position(),
			}
		}
	}

	// For any functions that are constructors we need to register the interface
	// types they produce. However, there is a catch. Function will be iterated,
	// but can also reference each other in any order. So we have to keep
	// iterating this step until all constructors have been resolved.
	//
	// This isn't the best way to handle this but it will have to do for now.
	constructorNames := map[string]struct{}{}
	for {
		resolved := 0
		for _, fnName := range p.SortedFuncNames() {
			fn := p.Funcs()[fnName]
			if fn.IsConstructor() {
				constructorNames[fn.Name] = struct{}{}

				interfaceType, err := fn.Interface()
				if err != nil {
					return nil, nil, []error{err}
				}

				_, err = file.Types.Add(types.NewInterface(fn.Name, interfaceType))
				if err == nil {
					delete(constructorNames, fn.Name)
					resolved++
				}
			}
		}

		if len(constructorNames) == 0 {
			break
		}

		if resolved == 0 {
			panic("infinite loop trying to resolve constructor interfaces")
		}
	}

	// Imports are treated as constants for the same reason as described above.
	// A import cannot be reassigned, but it also only needs to be initialized
	// once since it's not possible to call a package-level function and have
	// side effects.
	for pkgPath, pkgType := range imports {
		pkgVariable := path.Base(pkgPath)
		p.Constants[pkgVariable] = &ast.Literal{
			Kind:     pkgType,
			Value:    strings.ReplaceAll(pkgPath, "/", "__"),
			IsGlobal: true,
		}
		p.Constants["_"+pkgVariable] = &ast.Literal{
			Kind: types.NewFunc(nil, []*types.Type{pkgType}),
		}
	}

	compiledPackageFn, err := CompileFunc(p.Package(packageAlias), file,
		nil, p.Constants, imports, map[string]*types.Type{})
	if err != nil {
		return nil, nil, []error{err}
	}
	file.AddSymbolFunc(compiledPackageFn)

	file.Globals["$"+packageAlias] = compiledPackageFn.UniqueName

	// If there we imports we need to merge the objects together.
	//
	// Note: Since the type reference of compiledPackageFn may move we have to
	// resolve the type now for the return later.
	compiledPackageFnType := file.Types.Get(string(compiledPackageFn.Type)).Returns[0]
	if len(imports) > 0 {
		file = vm.Merge(append(append([]*vm.File(nil), file), dependencies...)...)
	}

	err = vm.Store(file, packageName)
	if err != nil {
		return nil, nil, []error{err}
	}

	if verbose {
		// TODO(elliot): Fix plurals.
		fmt.Printf("compiled %s: %d symbols, %d types\n",
			packageName, len(file.Symbols), len(file.Types))
	}

	return file, compiledPackageFnType, nil
}
