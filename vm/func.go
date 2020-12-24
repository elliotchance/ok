package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type CompiledFunc struct {
	Arguments    []string      `json:",omitempty"`
	Instructions *Instructions `json:",omitempty"`
	Registers    int
	variables    map[string]*types.Type // name: type
	Finally      []*Instructions        `json:",omitempty"`

	// These are copied from the function definition.
	// Name and Pos are used by the VM for stack traces.
	Type                  TypeRegister `json:",omitempty"`
	Name, UniqueName, Pos string       `json:",omitempty"`

	// These are only transient for the compiler.
	Parent                 *CompiledFunc           `json:"-"`
	DeferredFuncsToCompile []DeferredFunc          `json:"-"`
	Constants              map[string]*ast.Literal `json:"-"`
}

func NewCompiledFunc(
	fn *ast.Func,
	parentFunc *CompiledFunc,
	constants map[string]*ast.Literal,
	file *File,
) *CompiledFunc {
	return &CompiledFunc{
		variables:    map[string]*types.Type{},
		Constants:    constants,
		Type:         file.AddType(fn.Type()),
		Instructions: new(Instructions),

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
	Instructions *Instructions
}

func (c *CompiledFunc) NextRegister() Register {
	c.Registers++

	return Register(fmt.Sprintf("%d", c.Registers))
}

func (c *CompiledFunc) Append(instruction Instruction) {
	c.Instructions.Instructions = append(c.Instructions.Instructions, instruction)
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
