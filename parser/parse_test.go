package parser_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Funcs(t *testing.T) {
	t.Run("main function", func(t *testing.T) {
		p := parser.NewParser(0)
		p.ParseString(`func main() { }`, "a.ok")
		require.Nil(t, p.Errors())

		asttest.AssertEqual(t, map[string]*ast.Func{
			"1": {
				Name: "main",
				Pos:  "a.ok:1:1",
			},
		}, p.Funcs())
	})

	t.Run("calling ParseString multiple times appends functions", func(t *testing.T) {
		p := parser.NewParser(0)
		p.ParseString(`func main() { }`, "a.ok")
		p.ParseString(`func foo() { }`, "a.ok")
		require.Nil(t, p.Errors())

		asttest.AssertEqual(t, map[string]*ast.Func{
			"1": {
				Name: "main",
				Pos:  "a.ok:1:1",
			},
			"2": {
				Name: "foo",
				Pos:  "a.ok:1:1",
			},
		}, p.Funcs())
	})

	t.Run("normal function type", func(t *testing.T) {
		p := parser.NewParser(0)
		p.ParseString(`func main(a, b number) string { }`, "a.ok")
		require.Nil(t, p.Errors())

		assert.Equal(t, &types.Type{
			Kind:      types.KindFunc,
			Arguments: []*types.Type{types.Number, types.Number},
			Returns:   []*types.Type{types.String},
		}, p.Funcs()["1"].Type())
	})

	t.Run("constructor function", func(t *testing.T) {
		p := parser.NewParser(0)
		p.ParseString(`
func Person(Foo, bar, Baz number) Person {
	func Qux() string { }
	func quux() { }
	Corge = func () number { }
}`, "a.ok")
		require.Nil(t, p.Errors())

		assert.Equal(t, &types.Type{
			Kind:      types.KindFunc,
			Arguments: []*types.Type{types.Number, types.Number, types.Number},
			Returns: []*types.Type{
				{
					Kind: types.KindResolvedInterface,
					Name: "Person",
					Properties: map[string]*types.Type{
						"Foo": {Kind: types.KindNumber},
						"Baz": {Kind: types.KindNumber},
						"Qux": {
							Kind:    types.KindFunc,
							Returns: []*types.Type{types.String},
						},
						"Corge": {
							Kind:    types.KindFunc,
							Returns: []*types.Type{types.Number},
						},
					},
				},
			},
		}, p.Funcs()["1"].Type())
	})
}

func TestParser_ParseString(t *testing.T) {
	for testName, test := range map[string]struct {
		str       string
		expected  *ast.Func
		comments  []*ast.Comment
		errs      []error
		imports   map[string]string
		constants map[string]*ast.Literal
	}{
		"empty": {
			str: "",
		},
		"func-paren-close": {
			str: "func)",
			errs: []error{
				errors.New("a.ok:1:1 expecting ( after func, but found )"),
			},
		},
		"func-curly-open": {
			str: "func {",
			errs: []error{
				errors.New("a.ok:1:1 expecting ( after func, but found {"),
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
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							asttest.NewLiteralString("hello"),
						},
					},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
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
				errors.New("a.ok: found extra '(' at the end of the file"),
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
				errors.New("a.ok:1:15 expecting statement"),
				errors.New("a.ok:1:1 expecting statement"),
			},
		},
		"call-identifier-missing-close": {
			str: `func main() { print("hello" }`,
			errs: []error{
				errors.New("a.ok:1:15 expecting statement"),
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
					Expr: &ast.Identifier{Name: "print"},
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
					Expr: &ast.Identifier{Name: "print"},
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
		"package-level-variable": {
			str:      "Pi = 3.14\nfunc main() {}",
			expected: newFunc(),
			constants: map[string]*ast.Literal{
				"Pi": asttest.NewLiteralNumber("3.14"),
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.NewParser(0)
			p.ParseString(test.str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			if test.expected == nil {
				asttest.AssertEqual(t, map[string]*ast.Func{}, p.Funcs())
			} else {
				asttest.AssertEqual(t, map[string]*ast.Func{
					"1": test.expected,
				}, p.Funcs())
			}
			assert.Equal(t, test.comments, p.Comments())

			// Constants is always initialized by the parser, we only need to
			// specify in tests when needed.
			if test.constants == nil {
				test.constants = map[string]*ast.Literal{}
			}
			asttest.AssertEqual(t, test.constants, p.Constants)
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
				Expr:      &ast.Identifier{Name: "print"},
				Arguments: args,
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
