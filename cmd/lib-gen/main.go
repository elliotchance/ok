package main

import (
	"fmt"
	"io"
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
	pkg, errs := compiler.CompilePackage("lib/math", false)
	util.CheckErrorsWithExit(errs)

	f, err := os.Create("vm/lib.go")
	check(err)

	funcs := map[string]*vm.InternalDefinition{}
	for name, fn := range pkg.Funcs {
		funcDef := pkg.FuncDefs[name]

		// We don't need to serialize this.
		funcDef.Statements = nil

		funcs["math."+name] = &vm.InternalDefinition{
			CompiledFunc: fn,
			FuncDef:      funcDef,
		}
	}

	fmt.Fprintf(f, "package vm\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tLib = ")
	util.Render(f, funcs, "\t", true)
	fmt.Fprintf(f, "\n}\n")
	fmt.Fprintf(f, "\n")
}
