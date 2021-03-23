package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func TranspileLiteral(lit *ast.Literal) (Stmt, []*types.Type, error) {
	var val string
	switch lit.Kind.Kind {
	case types.KindBool:
		if lit.Value == "true" {
			val = "true"
		} else {
			val = "false"
		}

	case types.KindNumber:
		val = "Big(\"" + lit.Value + "\")"

	case types.KindString:
		val = "\"" + strings.ReplaceAll(lit.Value, "\"", "\\\"") + "\""

	default:
		val = fmt.Sprintf("{ t: '%s', v: %s }", lit.Kind, "\"" + lit.Value + "\"")
	}

	return Expr(val), []*types.Type{lit.Kind}, nil
}
