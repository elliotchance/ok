package compiler

import (
	"ok/ast"
	"ok/instruction"
	"os"
)

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func) ([]instruction.Instruction, error) {
	var instructions []instruction.Instruction
	for _, statement := range fn.Statements {
		ins := &instruction.Print{
			Stdout: os.Stdout,
		}
		args := statement.(*ast.Call).Arguments
		if len(args) > 0 {
			literal, err := compileNode(args[0])
			if err != nil {
				return nil, err
			}

			ins.Values = []*ast.Literal{literal}
		}

		instructions = append(instructions, ins)
	}

	return instructions, nil
}
