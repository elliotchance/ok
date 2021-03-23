package nexe

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileGroup(g *ast.Group, scope *Scope) (Stmt, []*types.Type, error) {
	expr, exprTypes, err := TranspileStmt(g.Expr, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	return Expr("(" + expr.JS(0) + ")"), exprTypes, nil
}
