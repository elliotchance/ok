package main

import (
	"fmt"
	"hash/crc32"
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
	crc32q := crc32.MakeTable(0xD5828281)
	for _, pathInfo := range packages {
		pkgName := pathInfo.Name()

		// lang is not importable - it only contains tests.
		if !pathInfo.IsDir() || pkgName == "lang" {
			continue
		}

		// TODO(elliot): This is a pretty crude way to make sure functions
		//  compiled from different inbuilt packages end up using predictable
		//  but unique function numbers.
		anonFunctionName := int(crc32.Checksum([]byte(pkgName), crc32q))

		f, errs := compiler.Compile("lib", pkgName, false, anonFunctionName)
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
