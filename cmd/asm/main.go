package asm

import (
	"fmt"
	"log"
	"sort"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/util"
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
	pkg, errs := compiler.Compile(okPath, packageName, false)
	util.CheckErrorsWithExit(errs)

	// Create a map as a function may match more than one glob.
	funcsToPrint := map[string]struct{}{}
	for _, glob := range args[1:] {
		for funcName := range pkg.FuncDefs {
			if util.MatchesGlob(funcName, glob) {
				funcsToPrint[funcName] = struct{}{}
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
