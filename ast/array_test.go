package ast_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/stretchr/testify/assert"
)

func TestNewArrayNumbers(t *testing.T) {
	t.Run("zero-elements", func(t *testing.T) {
		node := ast.NewArrayNumbers(nil)

		assert.Equal(t, "[]number", node.Kind)
		assert.Nil(t, node.Elements)
	})

	t.Run("two-elements", func(t *testing.T) {
		node := ast.NewArrayNumbers([]string{"123", "4.56"})

		assert.Equal(t, "[]number", node.Kind)
		assert.Equal(t, []ast.Node{
			ast.NewLiteralNumber("123"),
			ast.NewLiteralNumber("4.56"),
		}, node.Elements)
	})
}
