package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type For struct {
	Init, Condition, Next Stmt
	Block                 []Stmt
}

func (a *For) JS(indent int) string {
	var blockLines []string
	for _, v := range a.Block {
		blockLines = append(blockLines, fmt.Sprintf("%s%s;", indentString(indent+1), v.JS(indent+1)))
	}

	var init, condition, next string
	if a.Init != nil {
		init = "let " + a.Init.JS(indent)
	}
	if a.Condition != nil {
		condition = a.Condition.JS(indent)
	}
	if a.Next != nil {
		next = a.Next.JS(indent)
	}

	return fmt.Sprintf("for (%s; %s; %s) {\n%s\n%s}",
		init, condition, next,
		strings.Join(blockLines, "\n"), indentString(indent))
}

func TranspileFor(n *ast.For, scope *Scope) (*For, []*types.Type, error) {
	a := &For{}
	var err error

	if n.Init != nil {
		if assign, ok := n.Init.(*ast.Assign); ok {
			scope.variables[assign.Lefts[0].(*ast.Identifier).Name] = types.Number
		}

		a.Init, _, err = TranspileStmt(n.Init, scope)
		if err != nil {
			return nil, nil, err
		}
	}

	if n.Next != nil {
		a.Next, _, err = TranspileStmt(n.Next, scope)
		if err != nil {
			return nil, nil, err
		}
	}

	if n.Condition != nil {
		// In has to be handled at the loop level.
		if in, ok := n.Condition.(*ast.In); ok {
			var keyVar, valueVar string
			valueVar = scope.Add(in.Value, types.Any)
			if in.Key != "" {
				keyVar = scope.Add(in.Key, types.Any)
			}

			expr, exprType, err := TranspileStmt(in.Expr, scope)
			if err != nil {
				return nil, nil, err
			}

			variable := "i"

			if n.Init == nil {
				a.Init = Expr(fmt.Sprintf("%s = 0", variable))
			} else {
				a.Init = Expr(fmt.Sprintf("%s, %s = 0", a.Init.JS(0), variable))
			}

			if n.Next == nil {
				a.Next = Expr(fmt.Sprintf("++%s", variable))
			} else {
				a.Next = Expr(fmt.Sprintf("%s, ++%s", a.Next.JS(0), variable))
			}

			if exprType[0].Kind == types.KindMap {
				a.Condition = Expr(fmt.Sprintf("%s < Object.keys(%s[0]).length", variable, expr.JS(0)))
				a.Block = append(a.Block, Expr(fmt.Sprintf("let %s = Object.values(%s[0])[%s]", valueVar, expr.JS(0), variable)))
				if in.Key != "" {
					a.Block = append(a.Block, Expr(fmt.Sprintf("let %s = Object.keys(%s[0])[%s]", keyVar, expr.JS(0), variable)))
				}
			} else {
				a.Condition = Expr(fmt.Sprintf("%s < %s[0].length", variable, expr.JS(0)))
				a.Block = append(a.Block, Expr(fmt.Sprintf("let %s = %s[0][%s]", valueVar, expr.JS(0), variable)))
				if in.Key != "" {
					a.Block = append(a.Block, Expr(fmt.Sprintf("let %s = Big(%s)", keyVar, variable)))
				}
			}
		} else {
			a.Condition, _, err = TranspileStmt(n.Condition, scope)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	for _, stmt := range n.Statements {
		e, _, err := TranspileStmt(stmt, scope)
		if err != nil {
			return nil, nil, err
		}

		a.Block = append(a.Block, e)
	}

	return a, nil, nil
}
