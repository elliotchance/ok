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

func (c *CompiledFunc) newVariable(variableName string, kind string) {
	// TODO(elliot): Check already registered variables.
	c.variables[variableName] = kind
}

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func) (*CompiledFunc, error) {
	compiled := &CompiledFunc{
		variables: map[string]string{},
	}
	err := compileBlock(compiled, fn.Statements, nil, nil)
	if err != nil {
		return nil, err
	}

	return compiled, nil
}
