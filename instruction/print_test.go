package instruction_test

import (
	"ok/instruction"
)

func ExamplePrint_Execute() {
	ins := instruction.Print{Value: "hello world"}
	ins.Execute()
	// Output: hello world
}
