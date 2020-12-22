package compile

import (
	"encoding/json"
	"fmt"
	"log"

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
	return "compile"
}

// Run is the entry point for the "ok asm" command.
func (*Command) Run(args []string) {
	if len(args) < 1 {
		args = []string{"."}
	}

	okPath, err := util.OKPath()
	check(err)

	for _, arg := range args {
		packageName := util.PackageNameFromPath(okPath, arg)
		if args[0] == "." {
			packageName = "."
		}
		pkg, errs := compiler.Compile(okPath, packageName, false, 0)
		util.CheckErrorsWithExit(errs)

		data, err := json.MarshalIndent(pkg, "", "  ")
		check(err)
		fmt.Println(string(data))
	}
}
