package ast_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
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

func TestNewLiteralBool(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		literal := ast.NewLiteralBool(true)
		assert.Equal(t, "bool", literal.Kind)
		assert.Equal(t, "true", literal.Value)
	})

	t.Run("false", func(t *testing.T) {
		literal := ast.NewLiteralBool(false)
		assert.Equal(t, "bool", literal.Kind)
		assert.Equal(t, "false", literal.Value)
	})
}

func TestNewLiteralChar(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		literal := ast.NewLiteralChar('a')
		assert.Equal(t, "char", literal.Kind)
		assert.Equal(t, "a", literal.Value)
	})

	t.Run("emoji", func(t *testing.T) {
		literal := ast.NewLiteralChar('😃')
		assert.Equal(t, "char", literal.Kind)
		assert.Equal(t, "😃", literal.Value)
	})
}
