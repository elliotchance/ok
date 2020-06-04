package parser_test

import (
	"errors"
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Func
		comments []*ast.Comment
		err      error
	}{
		"empty": {
			str: "",
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
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenString,
								Value: "hello world",
							},
						},
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
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenString,
								Value: "hello",
							},
						},
					},
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenString,
								Value: "world",
							},
						},
					},
				},
			},
		},
		"extra-token": {
			str: "func main() {\n} (",
			err: errors.New("found extra '(' at the end of the file"),
		},
		"only-comment": {
			str: "// nothing to see here",
			comments: []*ast.Comment{
				{Comment: " nothing to see here"},
			},
		},
		"comments-everywhere": {
			str: "// foo\n //bar\nfunc main() {\n// baz\nprint(\"hello\") // qux\n// quux\n}//corge\n//grault",
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenString,
								Value: "hello",
							},
						},
					},
				},
			},
			comments: []*ast.Comment{
				{Comment: " foo"},
				{Comment: "bar"},
				{Comment: " baz"},
				{Comment: " qux"},
				{Comment: " quux"},
				{Comment: "corge"},
				{Comment: "grault"},
			},
		},
		"literal-true": {
			str: `func main() { print(true) }`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenBool,
								Value: "true",
							},
						},
					},
				},
			},
		},
		"literal-false": {
			str: `func main() { print(false) }`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenBool,
								Value: "false",
							},
						},
					},
				},
			},
		},
		"literal-char": {
			str: `func main() { print('a') }`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenCharacter,
								Value: "a",
							},
						},
					},
				},
			},
		},
		"literal-zero-length-char": {
			str: `func main() { print('') }`,
			err: errors.New("character literal cannot be empty"),
		},
		"literal-number-zero": {
			str: `func main() { print(0) }`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []*ast.Literal{
							{
								Kind:  lexer.TokenNumber,
								Value: "0",
							},
						},
					},
				},
			},
		},
		"call-identifier-close": {
			str: `func main() { print) }`,
			err: errors.New("expecting ( after identifier, but found )"),
		},
		"call-identifier-without-literal": {
			str: `func main() { print( }`,
			err: errors.New("expecting literal after (, but found }"),
		},
		"call-identifier-missing-close": {
			str: `func main() { print("hello" }`,
			err: errors.New("expecting ) after string, but found }"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			f, err := parser.ParseString(test.str)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}

			if f != nil {
				assert.Equal(t, test.expected, f.Root)
				assert.Equal(t, test.comments, f.Comments)
			}
		})
	}
}
