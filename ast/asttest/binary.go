package asttest

import "github.com/elliotchance/ok/ast"

// NewBinary creates a binary operation.
func NewBinary(left ast.Node, op string, right ast.Node) *ast.Binary {
	return &ast.Binary{
		Left:  left,
		Op:    op,
		Right: right,
	}
}
