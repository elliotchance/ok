package ast_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/stretchr/testify/assert"
)

func TestTypeOf(t *testing.T) {
	tests := map[string]struct {
		n        ast.Node
		expected string
	}{
		"literal-bool": {
			&ast.Literal{
				Kind:  "bool",
				Value: "true",
			},
			"bool",
		},
		"func": {
			&ast.Func{
				Name: "foo",
				Arguments: []*ast.Argument{
					{Name: "bar", Type: "string"},
				},
				Returns: []string{"number"},
			},
			"func(bar string) number",
		},
		"closure": {
			&ast.Func{},
			"func()",
		},
		"interpolate": {
			&ast.Interpolate{},
			"string",
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			ty, err := ast.TypeOf(tt.n)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, ty)
		})
	}
}
