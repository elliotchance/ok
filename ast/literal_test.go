package ast_test

import (
	"ok/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLiteralData(t *testing.T) {
	literal := ast.NewLiteralData([]byte("foo"))
	assert.Equal(t, "data", literal.Kind)
	assert.Equal(t, "foo", literal.Value)
}

func TestNewLiteralNumber(t *testing.T) {
	literal := ast.NewLiteralNumber("1.23")
	assert.Equal(t, "number", literal.Kind)
	assert.Equal(t, "1.23", literal.Value)
}

func TestNewLiteralString(t *testing.T) {
	literal := ast.NewLiteralString("foo bar")
	assert.Equal(t, "string", literal.Kind)
	assert.Equal(t, "foo bar", literal.Value)
}
