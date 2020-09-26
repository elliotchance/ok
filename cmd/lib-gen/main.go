package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

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

	pkgs := map[string]*vm.File{}
	for _, pathInfo := range packages {
		pkgName := pathInfo.Name()

		// lang is not importable - it only contains tests.
		if !pathInfo.IsDir() || pkgName == "lang" {
			continue
		}

		f, errs := compiler.Compile("lib", pkgName, false, 0)
		util.CheckErrorsWithExit(errs)

		pkgs[pkgName] = f
	}

	f, err := os.Create("vm/lib.go")
	check(err)

	fmt.Fprintf(f, "package vm\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/types\"\n\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tPackages = ")
	vm.Render(f, pkgs, "\t", true)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "}\n")
	fmt.Fprintf(f, "\n")
}
