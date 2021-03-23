package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileInterpolate(n *ast.Interpolate, scope *Scope) (Stmt, []*types.Type, error) {
	var s string
	for i, p := range n.Parts {

		if i%2 == 0 {
			s += p.(*ast.Literal).Value
		} else {
			expr, _, err := TranspileStmt(p, scope)
			if err != nil {
				return nil, nil, err
			}
			s += fmt.Sprintf("${%s}", expr.JS(0))
		}
	}

	return Expr("`" + s + "`"), []*types.Type{types.String}, nil
}
