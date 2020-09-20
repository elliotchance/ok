package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Call
	}{
		"no-args": {
			str: "foo()",
			expected: &ast.Call{
				Expr: &ast.Identifier{Name: "foo"},
			},
		},
		"one-arg": {
			str: `bar("baz")`,
			expected: &ast.Call{
				Expr: &ast.Identifier{Name: "bar"},
				Arguments: []ast.Node{
					asttest.NewLiteralString("baz"),
				},
			},
		},
		"cast-string": {
			str: `string 'a'`,
			expected: &ast.Call{
				Expr: &ast.Identifier{Name: "string"},
				Arguments: []ast.Node{
					asttest.NewLiteralChar('a'),
				},
			},
		},
		"math-abs": {
			str: `math.abs(123)`,
			expected: &ast.Call{
				Expr: &ast.Key{
					Expr: &ast.Identifier{Name: "math"},
					Key:  &ast.Identifier{Name: "abs"},
				},
				Arguments: []ast.Node{
					asttest.NewLiteralNumber("123"),
				},
			},
		},
		"cast-string-index": {
			str: `string foo[10]`,
			expected: &ast.Call{
				Expr: &ast.Identifier{Name: "string"},
				Arguments: []ast.Node{
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  asttest.NewLiteralNumber("10"),
					},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser()
			p.ParseString(str, "a.ok")

			assert.Empty(t, p.Errors().String())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.Funcs())
			assert.Nil(t, p.Comments())
		})
	}
}
