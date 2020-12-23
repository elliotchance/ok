package asm

import (
	"fmt"
	"log"
	"sort"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

type Command struct{}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

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

	okPath, err := util.OKPath()
	check(err)

	packageName := util.PackageNameFromPath(okPath, args[0])
	if args[0] == "." {
		packageName = "."
	}
	pkg, errs := compiler.Compile(okPath, packageName, false, 0)
	util.CheckErrorsWithExit(errs)

	// Create a map as a function may match more than one glob.
	funcsToPrint := map[string]struct{}{}
	for _, glob := range args[1:] {
		for _, fn := range pkg.Symbols {
			if fn.Func != nil &&
				(util.MatchesGlob(fn.Func.UniqueName, glob) ||
					util.MatchesGlob(fn.Func.Name, glob)) {
				funcsToPrint[fn.Func.UniqueName] = struct{}{}
			}
		}
	}

	// Display all functions in order. It makes the output easier to read but
	// also deterministic if we need to diff the output.
	var funcs []string
	for funcName := range funcsToPrint {
		funcs = append(funcs, funcName)
	}
	sort.Strings(funcs)

	for i, fnName := range funcs {
		// Just for vanity, put an empty line between functions.
		if i > 0 {
			fmt.Println()
		}

		fn := pkg.Symbols[vm.SymbolRegister(fnName)]
		if fn == nil {
			// TODO(elliot): This could be handled more gracefully.
			panic("no such function: " + fnName)
		}

		if fn.Func.Name != "" {
			fmt.Printf("%s: ", fn.Func.Name)
		}
		fmt.Println(fn.Func.UniqueName+":", fn.Type+":")

		for i, ins := range fn.Func.Instructions {
			ty := fmt.Sprintf("%T", ins)[4:]

			// "-22" is chosen here because is is the longer instruction name.
			fmt.Printf("  %3d %-22s # %s\n", i+1, ty, ins)
		}
	}
}
