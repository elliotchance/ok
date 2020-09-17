package asttest_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/stretchr/testify/assert"
)

func TestNewArrayNumbers(t *testing.T) {
	t.Run("zero-elements", func(t *testing.T) {
		node := asttest.NewArrayNumbers(nil)

		assert.Equal(t, "[]number", node.Kind.String())
		assert.Nil(t, node.Elements)
	})

	t.Run("two-elements", func(t *testing.T) {
		node := asttest.NewArrayNumbers([]string{"123", "4.56"})

		assert.Equal(t, "[]number", node.Kind.String())
		assert.Equal(t, []ast.Node{
			asttest.NewLiteralNumber("123"),
			asttest.NewLiteralNumber("4.56"),
		}, node.Elements)
	})
}
