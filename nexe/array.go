package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Array struct {
	Elements []Stmt
	Type     *types.Type
}

func NewArray(ty *types.Type) *Array {
	return &Array{
		Type: ty,
	}
}

func (a *Array) JS(indent int) string {
	var lines []string
	for _, v := range a.Elements {
		lines = append(lines, fmt.Sprintf("%s%s,", indentString(indent), v.JS(indent)))
	}

	return fmt.Sprintf("[ [\n%s\n%s], '%s' ]",
		strings.Join(lines, "\n"), indentString(indent-1), a.Type.String())
}

func TranspileArray(expr *ast.Array, scope *Scope) (*Array, []*types.Type, error) {
	a := NewArray(types.AnyArray)
	for _, el := range expr.Elements {
		e, _, err := TranspileStmt(el, scope)
		if err != nil {
			return nil, nil, err
		}

		a.Elements = append(a.Elements, e)
	}

	return a, []*types.Type{a.Type}, nil
}
