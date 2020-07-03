package vm

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Len is used to determine the size of an array or map.
type Len struct {
	Argument, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Len) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	r := registers[ins.Argument]
	var result int
	if strings.HasPrefix(r.Kind, "[]") {
		result = len(r.Array)
	} else {
		result = len(r.Map)
	}

	registers[ins.Result] = asttest.NewLiteralNumber(fmt.Sprintf("%d", result))

	return nil
}
