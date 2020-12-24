package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// A Symbol contains a literal value that can be referenced by instructions.
type Symbol struct {
	Type TypeRegister

	// Value can only be used when Func is nil.
	Value string `json:",omitempty"`

	Func *CompiledFunc `json:",omitempty"`

	// Interface will only be set when Func is provided.
	Interface string `json:",omitempty"`
}

func (s *Symbol) Literal(registry types.Registry) *ast.Literal {
	return &ast.Literal{
		Kind:  registry.Get(string(s.Type)),
		Value: s.Value,
	}
}
