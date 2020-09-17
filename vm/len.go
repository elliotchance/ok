package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
)

// Len is used to determine the size of an array or map.
type Len struct {
	Argument, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Len) Execute(_ *int, vm *VM) error {
	r := vm.Get(ins.Argument)
	var result int
	switch r.Kind.Kind {
	case types.KindArray:
		result = len(r.Array)

	case types.KindMap:
		result = len(r.Map)

	case types.KindString:
		result = len([]rune(r.Value))

	default:
		result = len(r.Value)
	}

	vm.Set(ins.Result, asttest.NewLiteralNumber(fmt.Sprintf("%d", result)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Len) String() string {
	return fmt.Sprintf("%s = len(%s)", ins.Result, ins.Argument)
}
