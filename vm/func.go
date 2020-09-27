package vm

import (
	"fmt"

	"github.com/elliotchance/ok/types"
)

type CompiledFunc struct {
	Arguments    []string
	Instructions []Instruction
	Registers    int
	Variables    map[string]*types.Type // name: type
	Finally      [][]Instruction

	// These are copied from the function definition.
	Type             *types.Type
	Name, UniqueName string
}

type FinallyBlock struct {
	Run          bool
	Instructions []Instruction
}

func (c *CompiledFunc) NextRegister() Register {
	c.Registers++

	return Register(fmt.Sprintf("%d", c.Registers))
}

func (c *CompiledFunc) Append(instruction Instruction) {
	c.Instructions = append(c.Instructions, instruction)
}

func (c *CompiledFunc) NewVariable(variableName string, kind *types.Type) {
	// TODO(elliot): Check already registered variables.
	c.Variables[variableName] = kind
}
