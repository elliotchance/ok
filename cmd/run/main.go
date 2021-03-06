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

	okPath, err := util.OKPath()
	check(err)

	for _, arg := range args {
		packageName := util.PackageNameFromPath(okPath, arg)
		if arg == "." {
			packageName = "."
		}

		m := vm.NewVM("no-package")
		anonFunctionName := 0
		_, packageType, errs := compiler.Compile(okPath, packageName, false,
			&anonFunctionName, false)
		util.CheckErrorsWithExit(errs)

		check(m.LoadPackage(packageName))
		check(m.Run("$" + packageType.Name))
	}
}
