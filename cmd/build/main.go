package build

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/nexe"
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
	var flagCompile, flagNexe, flagVerbose bool
	var flagOutput string
	flag.BoolVar(&flagCompile, "c", false, "compile only")
	flag.BoolVar(&flagNexe, "nexe", false, "compile with nexe")
	flag.BoolVar(&flagVerbose, "v", false, "verbose output")
	flag.StringVar(&flagOutput, "o", "", "output path")
	check(flag.CommandLine.Parse(args))
	args = flag.Args()

	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		runArg(arg, flagCompile, flagNexe, flagVerbose, flagOutput)
	}
}

func runArg(arg string, flagCompile, flagNexe, flagVerbose bool, flagOutput string) {
	okPath, err := util.OKPath()
	check(err)

	packageName := util.PackageNameFromPath(okPath, arg)
	if arg == "." {
		packageName = "."
	}

	if flagNexe {
		anonFunctionName := 0
		js, errs := nexe.TranspilePackage(okPath, packageName, &anonFunctionName)
		util.CheckErrorsWithExit(errs)

		jsFile := path.Join(arg, "main.js")
		f, err := os.Create(jsFile)
		check(err)

		_, err = f.WriteString(nexe.Lib)
		check(err)

		_, err = f.WriteString(js.JS(0))
		check(err)

		_, err = f.WriteString(fmt.Sprintf("\n\ntry {\n  %s[0].main();\n} catch (e) {\n  console.log($stack(e));\n  process.exit(1);\n}\n",
			nexe.PackageFunctionName(packageName)))
		check(err)

		if flagCompile {
			return
		}
		//defer os.Remove(jsFile)

		nexeExecutable, err := exec.LookPath("nexe")
		check(err)

		if flagOutput == "" {
			flagOutput = path.Base(arg)
		}

		nexeBuild := &exec.Cmd{
			Path:   nexeExecutable,
			Args:   []string{nexeExecutable, "--build", "-o", flagOutput, jsFile},
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}

		check(nexeBuild.Run())

		return
	}

	anonFunctionName := 0
	file, _, errs := compiler.Compile(okPath, packageName, false,
		&anonFunctionName, flagVerbose)
	util.CheckErrorsWithExit(errs)

	if flagCompile {
		return
	}

	goFile := path.Join(arg, "main.go")
	f, err := os.Create(goFile)
	check(err)

	defer os.Remove(goFile)

	fmt.Fprintf(f, "package main\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/vm\"\n\n")
	fmt.Fprintf(f, "func main() {\n")
	fmt.Fprintf(f, "\tfile := ")
	vm.Render(f, file, "\t", false)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "\tm := vm.NewVM(nil, nil, nil, \"\")\n")
	fmt.Fprintf(f, "\tif err := m.LoadFile(\"\", file); err != nil {\n")
	fmt.Fprintf(f, "\t\tpanic(err)\n")
	fmt.Fprintf(f, "\t}\n")
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
