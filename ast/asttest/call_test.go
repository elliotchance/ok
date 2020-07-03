package asttest_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/stretchr/testify/assert"
)

func TestNewCall(t *testing.T) {
	call := asttest.NewCall(
		"foo",
		asttest.NewLiteralString("bar"),
		asttest.NewLiteralString("baz"),
	)
	assert.Equal(t, "foo", call.FunctionName)
	assert.Equal(t, []ast.Node{
		asttest.NewLiteralString("bar"),
		asttest.NewLiteralString("baz"),
	}, call.Arguments)
}
