package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

var f io.Writer

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	packages, err := ioutil.ReadDir("lib")
	check(err)

	packageNames := map[string]bool{}

	funcs := map[string]*vm.InternalDefinition{}
	constants := map[string]*ast.Literal{}
	interfaces := map[string]map[string]string{}
	for _, pathInfo := range packages {
		if !pathInfo.IsDir() {
			continue
		}

		pkgName := pathInfo.Name()

		// lang is not importable. It's basic language features and tests.
		if pkgName != "lang" {
			packageNames[pkgName] = true
		}

		pkg, errs := compiler.CompilePackage("lib/"+pkgName, false)
		util.CheckErrorsWithExit(errs)

		for name, fn := range pkg.Funcs {
			if !util.IsPublic(name) {
				continue
			}

			funcDef := pkg.FuncDefs[name]

			// We don't need to serialize this.
			funcDef.Statements = nil

			key := pkgName + "." + name
			if pkgName == "lang" {
				key = name
			}

			funcs[key] = &vm.InternalDefinition{
				CompiledFunc: fn,
				FuncDef:      funcDef,
			}
		}

		for name, c := range pkg.Constants {
			if !util.IsPublic(name) {
				continue
			}

			key := pkgName + "." + name
			if pkgName == "lang" {
				key = name
			}

			constants[key] = c
		}

		for name, c := range pkg.Interfaces {
			if !util.IsPublic(name) {
				continue
			}

			key := pkgName + "." + name
			if pkgName == "lang" {
				key = name
			}

			interfaces[key] = c
		}
	}

	f, err := os.Create("vm/lib.go")
	check(err)

	fmt.Fprintf(f, "package vm\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tPackages = ")
	vm.Render(f, packageNames, "\t", true)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\tLib = ")
	vm.Render(f, funcs, "\t", true)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\tConstants = ")
	vm.Render(f, constants, "\t", true)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\tInterfaces = ")
	vm.Render(f, interfaces, "\t", true)
	fmt.Fprintf(f, "\n}\n")
	fmt.Fprintf(f, "\n")
}
