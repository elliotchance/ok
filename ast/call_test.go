package ast_test

import (
	"ok/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCall(t *testing.T) {
	call := ast.NewCall(
		"foo",
		ast.NewLiteralString("bar"),
		ast.NewLiteralString("baz"),
	)
	assert.Equal(t, "foo", call.FunctionName)
	assert.Equal(t, []ast.Node{
		ast.NewLiteralString("bar"),
		ast.NewLiteralString("baz"),
	}, call.Arguments)
}
