package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

type CompiledFunc struct {
	Instructions []instruction.Instruction
	Registers    int
	variables    map[string]string
}

func (c *CompiledFunc) nextRegister() string {
	c.Registers++
	return fmt.Sprintf("%d", c.Registers)
}

func (c *CompiledFunc) append(instruction instruction.Instruction) {
	c.Instructions = append(c.Instructions, instruction)
}

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func) (*CompiledFunc, error) {
	compiled := &CompiledFunc{
		variables: map[string]string{},
	}
	for _, statement := range fn.Statements {
		err := compileStatement(compiled, statement)
		if err != nil {
			return nil, err
		}
	}

	return compiled, nil
}
