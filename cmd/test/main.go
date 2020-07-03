package test

import (
	"fmt"
	"log"
	"os"
	"time"

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
	return "run tests"
}

// Run is the entry point for the "ok test" command.
func (*Command) Run(args []string) {
	// TODO: Multiple packages provided.
	if len(args) == 0 {
		args = []string{"."}
	}

	pkg, errs := compiler.CompilePackage(args[0], true)
	util.CheckErrorsWithExit(errs)

	m := vm.NewVM(pkg.Funcs, pkg.Tests, args[0])
	startTime := time.Now()
	err := m.RunTests()
	elapsed := time.Since(startTime)
	check(err)

	assertWord := pluralise("assert", m.TotalAssertions)
	if m.TestsFailed > 0 {
		fmt.Printf("%s: %d failed %d passed %d %s (%s)\n",
			args[0], m.TestsFailed, m.TestsPass,
			m.TotalAssertions, assertWord, elapsed)
	} else {
		fmt.Printf("%s: %d passed %d %s (%s)\n",
			args[0], m.TestsPass,
			m.TotalAssertions, assertWord, elapsed)
	}

	if m.TestsFailed > 0 {
		os.Exit(1)
	}
}

func pluralise(word string, n int) string {
	if n == 1 {
		return word
	}

	return word + "s"
}
