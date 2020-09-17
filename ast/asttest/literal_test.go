package asttest_test

import (
	"testing"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/stretchr/testify/assert"
)

func TestNewLiteralData(t *testing.T) {
	literal := asttest.NewLiteralData([]byte("foo"))
	assert.Equal(t, "data", literal.Kind.String())
	assert.Equal(t, "foo", literal.Value)
}

func TestNewLiteralNumber(t *testing.T) {
	literal := asttest.NewLiteralNumber("1.23")
	assert.Equal(t, "number", literal.Kind.String())
	assert.Equal(t, "1.23", literal.Value)
}

func TestNewLiteralString(t *testing.T) {
	literal := asttest.NewLiteralString("foo bar")
	assert.Equal(t, "string", literal.Kind.String())
	assert.Equal(t, "foo bar", literal.Value)
}

func TestNewLiteralBool(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		literal := asttest.NewLiteralBool(true)
		assert.Equal(t, "bool", literal.Kind.String())
		assert.Equal(t, "true", literal.Value)
	})

	t.Run("false", func(t *testing.T) {
		literal := asttest.NewLiteralBool(false)
		assert.Equal(t, "bool", literal.Kind.String())
		assert.Equal(t, "false", literal.Value)
	})
}

func TestNewLiteralChar(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		literal := asttest.NewLiteralChar('a')
		assert.Equal(t, "char", literal.Kind.String())
		assert.Equal(t, "a", literal.Value)
	})

	t.Run("emoji", func(t *testing.T) {
		literal := asttest.NewLiteralChar('😃')
		assert.Equal(t, "char", literal.Kind.String())
		assert.Equal(t, "😃", literal.Value)
	})
}
