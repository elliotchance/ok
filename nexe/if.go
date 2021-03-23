package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type If struct {
	Condition   Stmt
	True, False []Stmt
}

func (a *If) JS(indent int) string {
	var trueLines, falseLines []string
	for _, v := range a.True {
		trueLines = append(trueLines, fmt.Sprintf("%s%s;", indentString(indent+1), v.JS(indent+1)))
	}
	for _, v := range a.False {
		falseLines = append(falseLines, fmt.Sprintf("%s%s;", indentString(indent+1), v.JS(indent+1)))
	}

	if len(falseLines) > 0 {
		return fmt.Sprintf("if (%s) {\n%s\n%s} else {\n%s\n%s}", a.Condition.JS(0),
			strings.Join(trueLines, "\n"), indentString(indent),
			strings.Join(falseLines, "\n"), indentString(indent))
	}

	return fmt.Sprintf("if (%s) {\n%s\n%s}", a.Condition.JS(0),
		strings.Join(trueLines, "\n"), indentString(indent))
}

func TranspileIf(n *ast.If, scope *Scope) (*If, []*types.Type, error) {
	a := &If{}
	var err error

	a.Condition, _, err = TranspileStmt(n.Condition, scope)
	if err != nil {
		return nil, nil, err
	}

	for _, stmt := range n.True {
		e, _, err := TranspileStmt(stmt, scope)
		if err != nil {
			return nil, nil, err
		}

		a.True = append(a.True, e)
	}

	for _, stmt := range n.False {
		e, _, err := TranspileStmt(stmt, scope)
		if err != nil {
			return nil, nil, err
		}

		a.False = append(a.False, e)
	}

	return a, nil, nil
}
