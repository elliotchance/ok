package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type CompiledFunc struct {
	Arguments    []string
	Instructions []Instruction
	Registers    int
	Variables    map[string]*types.Type // name: type
	Finally      [][]Instruction

	// These are copied from the function definition.
	// Name and Pos are used by the VM for stack traces.
	Type                  *types.Type
	Name, UniqueName, Pos string

	// These are only transient for the compiler, they will be nil when
	// serialized.
	Parent                 *CompiledFunc
	DeferredFuncsToCompile []DeferredFunc
}

type DeferredFunc struct {
	Register Register
	Func     *ast.Func
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
