package nexe

import (
	"fmt"
	"os"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileCall(call *ast.Call, scope *Scope) (Stmt, []*types.Type, error) {
	expr, exprType, err := TranspileStmt(call.Expr, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	var args []string
	for _, arg := range call.Arguments {
		a, _, err := TranspileStmt(arg, scope)
		if err != nil {
			return Expr(""), nil, err
		}

		args = append(args, a.JS(0))
	}

	var returnType []*types.Type
	if len(exprType) > 0 {
		returnType = exprType[0].Returns
	}

	cwd, _ := os.Getwd()
	return Expr(fmt.Sprintf("$call('%s', %s)(%s)",
		strings.TrimPrefix(call.Position(), cwd),
		expr, strings.Join(args, ", "))), returnType, nil
}
