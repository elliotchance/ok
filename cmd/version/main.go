package version

import (
	"fmt"
)

type Command struct{}

// Version will be replaced when building the binaries. See "make version".
const Version = "ok version unknown"

// Description is shown in "ok -help".
func (*Command) Description() string {
	return "print ok version"
}

// Run is the entry point for the "ok version" command.
func (*Command) Run(args []string) {
	fmt.Println(Version)
}
