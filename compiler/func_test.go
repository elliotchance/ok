package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseFunc(s string) *ast.Func {
	p := parser.NewParser(0)
	p.ParseString(s, "a.ok")
	return p.Funcs()["1"]
}

func TestFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"no-statements": {
			fn: &ast.Func{},
		},
		"one-statement-no-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
			},
		},
		"two-statements-with-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
					},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							asttest.NewLiteralString("hello"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Print{
					Arguments: []vm.Register{"1"},
				},
			},
		},
		"oos-exists-arg": {
			fn: parseFunc("func foo(baz number) { func bar() number { return ^baz } }"),
			expected: []vm.Instruction{
				&vm.AssignFunc{
					Result:     "2",
					Type:       "2",
					UniqueName: "2",
				},
				&vm.ParentScope{
					X: "2",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "2",
				},
			},
		},
		"oos-exists-var": {
			fn: parseFunc("func foo() { baz = 0\n func bar() number { return ^baz } }"),
			expected: []vm.Instruction{
				&vm.AssignFunc{
					Result:     "1",
					Type:       "2",
					UniqueName: "2",
				},
				&vm.ParentScope{
					X: "1",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "baz",
					Register: "2",
				},
			},
		},
		"oos-illegal": {
			fn:  parseFunc("func foo() { return ^bar }"),
			err: errors.New("bar does not exist in the parent scope"),
		},
		"oos-nonexistent": {
			fn:  parseFunc("func foo() { func bar() number { return ^baz } }"),
			err: errors.New("baz does not exist in the parent scope"),
		},
		"func-with-if-closure": {
			fn: &ast.Func{
				Name: "foo",
				Statements: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "bar"},
						},
						Rights: []ast.Node{
							&ast.Func{
								Name:       "bar",
								UniqueName: "2",
								Statements: []ast.Node{
									&ast.Call{
										Expr: &ast.Identifier{Name: "print"},
										Arguments: []ast.Node{
											asttest.NewLiteralString("c"),
										},
									},
								},
							},
						},
					},
					&ast.If{
						Condition: &ast.Binary{
							Op:    "==",
							Left:  asttest.NewLiteralBool(true),
							Right: asttest.NewLiteralBool(false),
						},
						True: []ast.Node{
							&ast.Call{
								Expr: &ast.Identifier{Name: "print"},
								Arguments: []ast.Node{
									asttest.NewLiteralString("a"),
								},
							},
						},
						False: []ast.Node{
							&ast.Call{
								Expr: &ast.Identifier{Name: "print"},
								Arguments: []ast.Node{
									asttest.NewLiteralString("b"),
								},
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignFunc{
					Result:     "1",
					Type:       "0",
					UniqueName: "2",
				},
				&vm.ParentScope{
					X: "1",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "0",
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "1",
				},
				&vm.Equal{
					Left:   "2",
					Right:  "3",
					Result: "4",
				},
				&vm.JumpUnless{ // 6
					Condition: "4",
					To:        9,
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "2",
				},
				&vm.Print{
					Arguments: vm.Registers{"5"},
				},
				&vm.Jump{ // 9
					To: 11,
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "3",
				},
				&vm.Print{ // 11
					Arguments: vm.Registers{"6"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.fn,
				&vm.File{
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
					Types:   types.Registry{},
				}, nil, nil, nil, map[string]*types.Type{})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: statements,
	}
}
