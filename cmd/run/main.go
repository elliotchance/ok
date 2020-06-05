package run

import (
	"io/ioutil"
	"log"
	"ok/compiler"
	"ok/parser"
	"ok/vm"
	"path"
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

	// TODO(elliot): Fix static file name.
	data, err := ioutil.ReadFile(path.Join(args[0], "main.ok"))
	check(err)

	fn, err := parser.ParseString(string(data))
	check(err)

	instructions := compiler.CompileFunction(fn)

	p := vm.NewVM(instructions)
	p.Run()
}
