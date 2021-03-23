package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileRaise(e *ast.Raise, scope *Scope) (Stmt, []*types.Type, error) {
	errExpr, _, err := TranspileStmt(e.Err, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	// Wrap the error value in a new instance of Error so that the stack trace
	// is attached in case we need to decode it later.
	return Expr(fmt.Sprintf("throw new $Error(%s)", errExpr.JS(0))),
		nil, nil
}
