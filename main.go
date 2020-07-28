package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/elliotchance/ok/cmd/asm"
	"github.com/elliotchance/ok/cmd/build"
	"github.com/elliotchance/ok/cmd/doc"
	"github.com/elliotchance/ok/cmd/run"
	"github.com/elliotchance/ok/cmd/test"
	"github.com/elliotchance/ok/cmd/version"
)

type command interface {
	Run(args []string)
	Description() string
}

var commands = map[string]command{
	"asm":     &asm.Command{},
	"build":   &build.Command{},
	"doc":     &doc.Command{},
	"run":     &run.Command{},
	"test":    &test.Command{},
	"version": &version.Command{},
}

func main() {
	flag.Usage = func() {
		fmt.Println("ok is a tool for managing ok source code.")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println("\tok <command> [arguments]")
		fmt.Println("")
		fmt.Println("The commands are:")
		fmt.Println("")

		var keys []string
		for name := range commands {
			keys = append(keys, name)
		}

		sort.Strings(keys)

		for _, name := range keys {
			fmt.Printf("\t%s\t\t%s\n", name, commands[name].Description())
		}

		fmt.Println("")
		fmt.Println(`Use "ok <command> -help" for more information about a command.`)
		fmt.Println("")
	}

	// This is just for the -help (usage).
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalln("missing command")
	}

	cmdName := os.Args[1]
	if cmd, ok := commands[cmdName]; ok {
		cmd.Run(os.Args[2:])
		return
	}

	log.Fatalln("unknown command:", cmdName)
}
