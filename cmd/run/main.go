package run

import (
	"log"

	"github.com/elliotchance/ok/compiler"
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
	// TODO: Multiple packages provided.
	if len(args) == 0 {
		args = []string{"."}
	}

	pkg, err := compiler.CompilePackage(args[0], false)
	check(err)

	m := vm.NewVM(pkg.Funcs, pkg.Tests, args[0])
	err = m.Run()
	check(err)
}
