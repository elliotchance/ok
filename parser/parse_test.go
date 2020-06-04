package parser_test

import (
	"errors"
	"ok/ast"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Func
		err      error
	}{
		"empty": {
			str: "",
			err: errors.New("cannot parse empty string"),
		},
		"func-paren-close": {
			str: "func)",
			err: errors.New("expecting identifier after func, but found )"),
		},
		"func-curly-open": {
			str: "func {",
			err: errors.New("expecting identifier after func, but found {"),
		},
		"func-name-paren-close": {
			str: "func main)",
			err: errors.New(`expecting ( after identifier, but found )`),
		},
		"func-name-paren-open": {
			str: "func main (",
			err: errors.New("expecting ) after (, but found end of file"),
		},
		"func-name-paren-open-close": {
			str: "func main ()",
			err: errors.New("expecting { after ), but found end of file"),
		},
		"func-name-paren-open-close-open": {
			str: "func main () {",
			err: errors.New("expecting } after {, but found end of file"),
		},
		"func-empty": {
			str:      "func main() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"unterminated-string": {
			str: `func "`,
			err: errors.New("unterminated string found after 'func'"),
		},
		"unterminated-string-first-token": {
			str: `"`,
			err: errors.New("unterminated string found at the start"),
		},
		"hello-world": {
			str: `func main() {print("hello world")}`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments:    []string{"hello world"},
					},
				},
			},
		},
		"hello-world-2": {
			str: `func main() {print("hello") print("world")}`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments:    []string{"hello"},
					},
					&ast.Call{
						FunctionName: "print",
						Arguments:    []string{"world"},
					},
				},
			},
		},
		"extra-token": {
			str: "func main() {\n} (",
			err: errors.New("found extra '(' at the end of the file"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			node, err := parser.ParseString(test.str)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, node)
		})
	}
}
