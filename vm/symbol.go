package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// A Symbol contains a literal value that can be referenced by instructions.
type Symbol struct {
	Type  string
	Value string
}

func (s *Symbol) Literal() *ast.Literal {
	ty := types.TypeFromString(s.Type)
	//if ty.Kind == types.KindMap ||
	//	ty.Kind == types.KindResolvedInterface ||
	//	ty.Kind == types.KindUnresolvedInterface {
	//	return &ast.Literal{
	//		Kind: ty,
	//		Map:  map[string]*ast.Literal{},
	//	}
	//}

	return &ast.Literal{
		Kind:  ty,
		Value: s.Value,
	}
}
