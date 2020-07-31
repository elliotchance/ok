package asm

import (
	"fmt"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/util"
)

type Command struct{}

// Description is shown in "ok -help".
func (*Command) Description() string {
	return "show compiled instructions"
}

// Run is the entry point for the "ok asm" command.
func (*Command) Run(args []string) {
	if len(args) < 1 {
		args = []string{"."}
	}
	if len(args) < 2 {
		args = append(args, "main")
	}

	pkg, errs := compiler.CompilePackage(args[0], false)
	util.CheckErrorsWithExit(errs)

	for i, fnName := range args[1:] {
		// Just for vanity, put an empty line between functions.
		if i > 0 {
			fmt.Println()
		}

		fn, exists := pkg.FuncDefs[fnName]
		if !exists {
			// TODO(elliot): This could be handled more gracefully.
			panic("no such function: " + fnName)
		}

		fmt.Println(fn.String() + ":")

		for i, ins := range pkg.Funcs[fn.Name].Instructions {
			ty := fmt.Sprintf("%T", ins)[4:]

			// "-22" is chosen here because is is the longer instruction name.
			fmt.Printf("  %3d %-22s # %s\n", i+1, ty, ins)
		}
	}
}
