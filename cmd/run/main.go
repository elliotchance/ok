package run

import (
	"log"

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
	return "run ok program"
}

// Run is the entry point for the "ok run" command.
func (*Command) Run(args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		packageName := util.PackageNameFromPath("", arg)

		pkg, errs := compiler.CompilePackage(arg, false)
		util.CheckErrorsWithExit(errs)

		m := vm.NewVM(pkg.Funcs, pkg.Tests, packageName)
		err := m.Run()
		check(err)
	}
}
