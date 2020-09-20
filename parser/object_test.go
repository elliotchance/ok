package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
)

func TestObject(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"get-property": {
			str: `foo.bar`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  &ast.Identifier{Name: "bar"},
			},
		},
		"set-property": {
			str: `foo.bar = 2`,
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  &ast.Identifier{Name: "bar"},
					},
				},
				Rights: []ast.Node{
					asttest.NewLiteralNumber("2"),
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser()
			p.ParseString(str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.Funcs())
		})
	}
}
