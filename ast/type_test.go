package ast_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestTypeOf(t *testing.T) {
	tests := map[string]struct {
		n        ast.Node
		expected *types.Type
	}{
		"literal-bool": {
			&ast.Literal{
				Kind:  types.Bool,
				Value: "true",
			},
			types.Bool,
		},
		"func": {
			&ast.Func{
				Name: "foo",
				Arguments: []*ast.Argument{
					{Name: "bar", Type: types.String},
				},
				Returns: []*types.Type{types.Number},
			},
			types.TypeFromString("func(string) number"),
		},
		"closure": {
			&ast.Func{},
			types.TypeFromString("func()"),
		},
		"interpolate": {
			&ast.Interpolate{},
			types.String,
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
