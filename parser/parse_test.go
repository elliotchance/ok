package parser_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Func
		comments []*ast.Comment
		errs     []error
		imports  map[string]string
	}{
		"empty": {
			str: "",
		},
		"func-paren-close": {
			str: "func)",
			errs: []error{
				errors.New("a.ok:1:1 expecting identifier after func, but found )"),
			},
		},
		"func-curly-open": {
			str: "func {",
			errs: []error{
				errors.New("a.ok:1:1 expecting identifier after func, but found {"),
			},
		},
		"func-name-paren-close": {
			str: "func main)",
			errs: []error{
				errors.New(`a.ok:1:1 expecting ( after identifier, but found )`),
			},
		},
		"func-name-paren-open": {
			str: "func main (",
			errs: []error{
				errors.New("a.ok:1:1 expecting identifier after (, but found end of file"),
			},
		},
		"func-name-paren-open-close": {
			str: "func main ()",
			errs: []error{
				errors.New("a.ok:1:1 expecting identifier after ), but found end of file"),
			},
		},
		"func-name-paren-open-close-open": {
			str: "func main () {",
			errs: []error{
				errors.New("a.ok:1:15 expecting statement"),
				errors.New("a.ok:1:1 expecting statement"),
			},
		},
		"func-empty": {
			str:      "func main() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"func-empty-2": {
			str:      "func\nmain() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"unterminated-string": {
			str: `func "`,
			errs: []error{
				errors.New("unterminated string found after 'func'"),
			},
		},
		"unterminated-string-first-token": {
			str: `"`,
			errs: []error{
				errors.New("unterminated string found at the start"),
			},
		},
		"hello-world": {
			str:      `func main() {print("hello world")}`,
			expected: newFuncPrint(asttest.NewLiteralString("hello world")),
		},
		"hello-world-2": {
			str: "func main() {print(\"hello\")\nprint(\"world\")}",
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							asttest.NewLiteralString("hello"),
						},
					},
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							asttest.NewLiteralString("world"),
						},
					},
				},
			},
		},
		"extra-token": {
			str:      "func main() {\n} (",
			expected: newFunc(),
			errs: []error{
				errors.New("found extra '(' at the end of the file"),
			},
		},
		"only-comment": {
			str: "// nothing to see here",
			comments: []*ast.Comment{
				{Comment: " nothing to see here", Pos: "a.ok:1:1"},
			},
		},
		"comments-everywhere": {
			str:      "// foo\n //bar\nfunc main() {\n// baz\nprint(\"hello\") // qux\n// quux\n}//corge\n//grault",
			expected: newFuncPrint(asttest.NewLiteralString("hello")),
			comments: []*ast.Comment{
				{Comment: " foo\nbar", Func: "main", Pos: "a.ok:1:1"},
				{Comment: " baz", Pos: "a.ok:4:1"},
				{Comment: " qux\n quux", Pos: "a.ok:5:16"},
				{Comment: "corge\ngrault", Pos: "a.ok:7:2"},
			},
		},
		"call-identifier-close": {
			str: `func main() { print) }`,
			errs: []error{
				// TODO(elliot): Prevent duplicate errors from bubbling up?
				errors.New("a.ok:1:20 expecting statement"),
				errors.New("a.ok:1:1 expecting statement"),
			},
		},
		"call-identifier-without-literal": {
			str: `func main() { print( }`,
			errs: []error{
				errors.New("a.ok:1:20 expecting statement"),
				errors.New("a.ok:1:1 expecting statement"),
			},
		},
		"call-identifier-missing-close": {
			str: `func main() { print("hello" }`,
			errs: []error{
				errors.New("a.ok:1:20 expecting statement"),
				errors.New("a.ok:1:1 expecting statement"),
			},
		},
		"print-2": {
			str: `func main() { print(true, false) }`,
			expected: newFuncPrint(
				asttest.NewLiteralBool(true),
				asttest.NewLiteralBool(false),
			),
		},
		"end-of-line-1": {
			str: "func main() { a = 1\nprint(a) }",
			expected: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1"),
					},
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
				},
			),
		},
		"end-of-line-2": {
			str: "func main() { b = true\nprint(b)\nc = 1.23 }",
			expected: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "b"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralBool(true),
					},
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "b"},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "c"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1.23"),
					},
				},
			),
		},
		"import-and-func": {
			str:      "import \"math\"\nfunc main() {}",
			expected: newFunc(),
			imports: map[string]string{
				"math": "math",
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			if test.expected == nil {
				asttest.AssertEqual(t, map[string]*ast.Func{}, p.File.Funcs)
			} else {
				asttest.AssertEqual(t, map[string]*ast.Func{
					"main": test.expected,
				}, p.File.Funcs)
			}
			assert.Equal(t, test.comments, p.File.Comments)
		})
	}
}

func assertEqualErrors(t *testing.T, expected, actual []error) {
	var e []string
	for _, err := range expected {
		e = append(e, err.Error())
	}

	var a []string
	for _, err := range actual {
		a = append(a, err.Error())
	}

	assert.Equal(t, e, a)
}

func newFuncPrint(args ...ast.Node) *ast.Func {
	return &ast.Func{
		Name: "main",
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    args,
			},
		},
	}
}

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Name:       "main",
		Statements: statements,
		Pos:        "a.ok:1:1",
	}
}
