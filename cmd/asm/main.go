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
	anonFunctionName := 0
	pkg, _, errs := compiler.Compile(okPath, packageName, false,
		&anonFunctionName, false)
	util.CheckErrorsWithExit(errs)

	// Create a map as a function may match more than one glob.
	funcsToPrint := map[vm.SymbolRegister]struct{}{}
	for _, glob := range args[1:] {
		for symbolRegister, symbol := range pkg.Symbols {
			if symbol.Func != nil &&
				(util.MatchesGlob(symbol.Func.UniqueName, glob) ||
					util.MatchesGlob(symbol.Func.Name, glob)) {
				funcsToPrint[symbolRegister] = struct{}{}
			} else if util.MatchesGlob(string(symbolRegister), glob) {
				funcsToPrint[symbolRegister] = struct{}{}
			}
		}
	}

	// Display all functions in order. It makes the output easier to read but
	// also deterministic if we need to diff the output.
	var funcs []string
	for funcName := range funcsToPrint {
		funcs = append(funcs, string(funcName))
	}
	sort.Strings(funcs)

	for i, symbolRegister := range funcs {
		// Just for vanity, put an empty line between functions.
		if i > 0 {
			fmt.Println()
		}

		symbol := pkg.Symbols[vm.SymbolRegister(symbolRegister)]
		if symbol == nil {
			// TODO(elliot): This could be handled more gracefully.
			panic("no such symbol: " + symbolRegister)
		}

		if symbol.Func == nil {
			fmt.Printf("%s (symbol=%s):\n  Value = \"%s\"\n",
				pkg.Types.Get(string(symbol.Type)).String(), symbolRegister,
				symbol.Value)
		} else {
			fmt.Printf("%s (symbol=%s, name=%s, unique=%s):\n",
				pkg.Types.Get(string(symbol.Type)).String(), symbolRegister,
				symbol.Func.Name, symbol.Func.UniqueName)

			for i, ins := range symbol.Func.Instructions.Instructions {
				ty := fmt.Sprintf("%T", ins)[4:]

				// "-22" is chosen here because is is the longer instruction name.
				fmt.Printf("  %3d %-22s # %s\n", i+1, ty, ins)
			}
		}
	}
}
