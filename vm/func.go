package vm

import "fmt"

type CompiledFunc struct {
	Arguments      []string
	Instructions   []Instruction
	Registers      int
	Variables      map[string]string // name: type
	ObjectRegister Register
	Finally        [][]Instruction
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

func (c *CompiledFunc) NewVariable(variableName string, kind string) {
	// TODO(elliot): Check already registered variables.
	c.Variables[variableName] = kind
}
