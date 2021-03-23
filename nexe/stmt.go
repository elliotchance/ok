package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Stmt interface {
	JS(indent int) string
}

type Expr string

func (s Expr) JS(_ int) string {
	return string(s)
}

type Stmts []Stmt

func (s Stmts) JS(indent int) string {
	var ss []string
	for _, stmt := range s {
		ss = append(ss, fmt.Sprintf("%s%s", indentString(indent), stmt.JS(indent)))
	}

	return strings.Join(ss, "\n")
}

func TranspileStmt(node ast.Node, scope *Scope) (Stmt, []*types.Type, error) {
	switch n := node.(type) {
	case *ast.Array:
		return TranspileArray(n, scope)

	case *ast.Assign:
		return TranspileAssign(n, scope)

	case *ast.Binary:
		return TranspileBinary(n, scope)

	case *ast.Break:
		return TranspileBreak()

	case *ast.Call:
		return TranspileCall(n, scope)

	case *ast.Continue:
		return TranspileContinue()

	case *ast.Key:
		return TranspileKey(n, scope)

	case *ast.Group:
		return TranspileGroup(n, scope)

	case *ast.Identifier:
		return TranspileIdentifier(n, scope)

	case *ast.If:
		return TranspileIf(n, scope)

	case *ast.Interpolate:
		return TranspileInterpolate(n, scope)

	case *ast.For:
		return TranspileFor(n, scope)

	case *ast.Func:
		f, tys, err := TranspileFunc(n, scope)
		if err != nil {
			return Expr(""), nil, err
		}

		return f, tys, nil

	case *ast.Literal:
		return TranspileLiteral(n)

	case *ast.Map:
		return TranspileMap(n, scope)

	case *ast.Raise:
		return TranspileRaise(n, scope)

	case *ast.Return:
		return TranspileReturn(n, scope)

	case *ast.Unary:
		return TranspileUnary(n, scope)

	default:
		panic(fmt.Sprintf("%T", n))
	}
}
