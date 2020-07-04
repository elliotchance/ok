package vm

import "fmt"

type CompiledFunc struct {
	Arguments      []string
	Instructions   []Instruction
	Registers      int
	Variables      map[string]string
	ObjectRegister string
}

func (c *CompiledFunc) NextRegister() string {
	c.Registers++
	return fmt.Sprintf("%d", c.Registers)
}

func (c *CompiledFunc) Append(instruction Instruction) {
	c.Instructions = append(c.Instructions, instruction)
}

func (c *CompiledFunc) NewVariable(variableName string, kind string) {
	// TODO(elliot): Check already registered variables.
	c.Variables[variableName] = kind
}
