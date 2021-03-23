package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Assign struct {
	Name       string
	Expr       Stmt
	IncludeLet bool
}

func (a *Assign) JS(indent int) string {
	var let string
	if a.IncludeLet {
		let = "let "
	}

	return fmt.Sprintf("%s%s = %s", let, a.Name, a.Expr.JS(indent+1))
}

func TranspileAssign(a *ast.Assign, scope *Scope) (Stmt, []*types.Type, error) {
	var lefts []string
	var rights []string

	for _, left := range a.Lefts {
		if leftIsKey, ok := left.(*ast.Key); ok {
			leftExpr, leftExprTypes, err := TranspileStmt(leftIsKey.Expr, scope)
			if err != nil {
				return nil, nil, err
			}

			leftKey, _, err := TranspileStmt(leftIsKey.Key, scope)
			if err != nil {
				return nil, nil, err
			}

			switch leftExprTypes[0].Kind {
			case types.KindArray:
				lefts = append(lefts, fmt.Sprintf("%s[0][(%s).toNumber()]", leftExpr, leftKey))

			default:
				lefts = append(lefts, fmt.Sprintf("%s[0][%s]", leftExpr, leftKey))
			}
		} else {
			localName := left.(*ast.Identifier).Name
			realName, _, err := scope.Get(localName)
			if err != nil {
				return Expr(""), nil, err
			}

			lefts = append(lefts, realName)
		}
	}

	var tys []*types.Type
	for _, right := range a.Rights {
		expr, ty, err := TranspileStmt(right, scope)
		if err != nil {
			return Expr(""), nil, err
		}

		rights = append(rights, expr.JS(0))
		tys = append(tys, ty[0])
	}

	if _, ok := a.Rights[0].(*ast.Call); ok {
		return &Assign{
			Name: "[" + strings.Join(lefts, ", ") + "]",
			Expr: Expr(rights[0]),
		}, tys, nil
	}

	return &Assign{
		Name: "[" + strings.Join(lefts, ", ") + "]",
		Expr: Expr("[" + strings.Join(rights, ", ") + "]"),
	}, tys, nil
}
