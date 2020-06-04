package compiler

import (
	"ok/ast"
	"ok/instruction"
)

// CompileFunction translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunction(fn *ast.Func) []instruction.Instruction {
	var instructions []instruction.Instruction
	for _, statement := range fn.Statements {
		ins := &instruction.Print{}
		args := statement.(*ast.Call).Arguments
		if len(args) > 0 {
			ins.Value = args[0]
		}

		instructions = append(instructions, ins)
	}

	return instructions
}
