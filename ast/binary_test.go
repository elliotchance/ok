package ast_test

import (
	"ok/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinary(t *testing.T) {
	node := ast.NewBinary(
		ast.NewLiteralData([]byte("foo")),
		"/",
		ast.NewLiteralData([]byte("bar")),
	)
	assert.Equal(t, ast.NewLiteralData([]byte("foo")), node.Left)
	assert.Equal(t, "/", node.Op)
	assert.Equal(t, ast.NewLiteralData([]byte("bar")), node.Right)
}
