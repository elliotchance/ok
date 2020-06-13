package compiler

import (
	"ok/ast"
)

func compileStatement(compiledFunc *CompiledFunc, statement ast.Node) error {
	switch n := statement.(type) {
	case *ast.Assign:
		return compileAssign(compiledFunc, n)
	}

	_, _, err := compileExpr(compiledFunc, statement)

	return err
}

func typeOf(node ast.Node) string {
	switch n := node.(type) {
	case *ast.Literal:
		return n.Kind

	case *ast.Binary:
		return typeOf(n.Left)

	case *ast.Group:
		return typeOf(n.Expr)
	}

	// TODO(elliot): Do something more sensible here.
	panic(node)
}
