package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileKey(e *ast.Key, scope *Scope) (Stmt, []*types.Type, error) {
	expr, exprType, err := TranspileStmt(e.Expr, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	key, _, err := TranspileStmt(e.Key, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	switch exprType[0].Kind {
	case types.KindArray:
		return Expr(fmt.Sprintf("%s[0][(%s).toNumber()]", expr, key)), nil, nil

	default:
		return Expr(fmt.Sprintf("%s[0][%s]", expr, key)), nil, nil
	}
}
