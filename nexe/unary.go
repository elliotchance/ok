package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

func TranspileUnary(e *ast.Unary, scope *Scope) (Stmt, []*types.Type, error) {
	expr, _, err := TranspileStmt(e.Expr, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	switch e.Op {
	case lexer.TokenIncrement:
		return Expr(fmt.Sprintf("%s = %s.add(1)", expr, expr)), nil, nil

	case lexer.TokenMinus:
		return Expr(fmt.Sprintf("Big(\"0\").minus(%s)", expr)), nil, nil
	}

	panic(e.Op)
}
