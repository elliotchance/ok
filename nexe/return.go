package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Return struct {
	Results []Stmt
}

func (r *Return) JS(indent int) string {
	var results []string
	for _, result := range r.Results {
		results = append(results, result.JS(indent+1))
	}

	return fmt.Sprintf("return [%s]", strings.Join(results, ", "))
}

func TranspileReturn(ret *ast.Return, scope *Scope) (Stmt, []*types.Type, error) {
	var exprs []Stmt
	for _, expr := range ret.Exprs {
		e, _, err := TranspileStmt(expr, scope)
		if err != nil {
			return Expr(""), nil, err
		}

		exprs = append(exprs, e)
	}

	return &Return{exprs}, nil, nil
}
