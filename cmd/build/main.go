package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/util"
	"github.com/elliotchance/ok/vm"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Command struct{}

// Description is shown in "ok -help".
func (*Command) Description() string {
	return "compile a program"
}

// Run is the entry point for the "ok run" command.
func (*Command) Run(args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		runArg(arg)
	}
}

func runArg(arg string) {
	pkg, errs := compiler.CompilePackage(arg, false)
	util.CheckErrorsWithExit(errs)

	goFile := path.Join(arg, "main.go")
	f, err := os.Create(goFile)
	check(err)

	defer os.Remove(goFile)

	fmt.Fprintf(f, "package main\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/vm\"\n\n")
	fmt.Fprintf(f, "func main() {\n")
	fmt.Fprintf(f, "\tfuncs := ")
	vm.Render(f, pkg.Funcs, "\t", false)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\tm := vm.NewVM(funcs, nil, \"\")\n")
	fmt.Fprintf(f, "\tif err := m.Run(); err != nil {\n")
	fmt.Fprintf(f, "\t\tpanic(err)\n")
	fmt.Fprintf(f, "\t}\n")
	fmt.Fprintf(f, "}\n")
	fmt.Fprintf(f, "\n")

	goExecutable, err := exec.LookPath("go")
	check(err)

	goBuild := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "build", "./" + arg},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	check(goBuild.Run())
}
