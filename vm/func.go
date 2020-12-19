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
	variables    map[string]*types.Type // name: type
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

func NewCompiledFunc(fn *ast.Func, parentFunc *CompiledFunc) *CompiledFunc {
	return &CompiledFunc{
		variables: map[string]*types.Type{},
		Type:      fn.Type(),

		// Name and Pos are used by the VM for stack traces.
		Name:       fn.Name,
		UniqueName: fn.UniqueName,
		Pos:        fn.Position(),

		Parent: parentFunc,
	}
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
	c.variables[variableName] = kind
}

func (c *CompiledFunc) GetTypeForVariable(
	variableName string,
	scopeOverrides map[string]*types.Type,
) (*types.Type, bool) {
	if ty, ok := scopeOverrides[variableName]; ok {
		return ty, true
	}

	ty, ok := c.variables[variableName]

	return ty, ok
}
