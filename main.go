package main

import (
	"io/ioutil"
	"log"
	"ok/compiler"
	"ok/parser"
	"ok/vm"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("missing command")
	}

	runCmd(os.Args[1], os.Args[2:])
}

func runCmd(cmd string, args []string) {
	switch cmd {
	case "run":
		cmdRun(args)

	default:
		log.Fatalln("unknown command:", cmd)
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func cmdRun(args []string) {
	// TODO: Multiple packages provided.
	if len(args) == 0 {
		args = []string{"."}
	}

	data, err := ioutil.ReadFile(path.Join(args[0], "main.ok"))
	check(err)

	fn, err := parser.ParseString(string(data))
	check(err)

	instructions := compiler.CompileFunction(fn)

	p := vm.NewVM(instructions)
	p.Run()
}
