package nexe

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileIdentifier(ident *ast.Identifier, scope *Scope) (Stmt, []*types.Type, error) {
	name, ty, err := scope.Get(ident.Name)
	if err != nil {
		return Expr(""), nil, err
	}

	return Expr(name), []*types.Type{ty}, nil
}
