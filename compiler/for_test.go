package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFor(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"for-in-array-1": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewArrayNumbers([]string{"1.5", "2.3"}),
					},
				},
				&ast.For{
					Condition: &ast.In{
						Key:  "k",
						Expr: &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc array
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number2",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 2 elements
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number0",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1.5",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number1",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number2.3",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					Result:   "foo",
					Register: "2",
				},

				// for in
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "number0",
				},
				&vm.NextArray{
					Array:       "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "",
					Result:      "8",
				},
				&vm.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&vm.Jump{
					To: 9,
				},
			},
		},
		"for-in-array-2": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewArrayNumbers([]string{"1.5", "2.3"}),
					},
				},
				&ast.For{
					Condition: &ast.In{
						Key:   "k",
						Value: "v",
						Expr:  &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc array
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number2",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 2 elements
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number0",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1.5",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number1",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number2.3",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					Result:   "foo",
					Register: "2",
				},

				// for in
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "number0",
				},
				&vm.NextArray{
					Array:       "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "v",
					Result:      "8",
				},
				&vm.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&vm.Jump{
					To: 9,
				},
			},
		},
		"for-in-map-1": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewMapNumbers(map[string]string{"foo": "1.5", "bar": "2.3"}),
					},
				},
				&ast.For{
					Condition: &ast.In{
						Key:  "k",
						Expr: &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc array
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number2",
				},
				&vm.MapAlloc{
					Kind:   "{}number",
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "stringbar",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2.3",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "stringfoo",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number1.5",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					Result:   "foo",
					Register: "2",
				},

				// for in
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "number0",
				},
				&vm.NextMap{
					Map:         "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "",
					Result:      "8",
				},
				&vm.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&vm.Jump{
					To: 9,
				},
			},
		},
		"for-in-map-2": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewMapNumbers(map[string]string{"foo": "1.5", "bar": "2.3"}),
					},
				},
				&ast.For{
					Condition: &ast.In{
						Key:   "k",
						Value: "v",
						Expr:  &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc array
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number2",
				},
				&vm.MapAlloc{
					Kind:   "{}number",
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "stringbar",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2.3",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "stringfoo",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number1.5",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					Result:   "foo",
					Register: "2",
				},

				// for in
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "number0",
				},
				&vm.NextMap{
					Map:         "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "v",
					Result:      "8",
				},
				&vm.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&vm.Jump{
					To: 9,
				},
			},
		},
		"for-3": {
			nodes: []ast.Node{
				&ast.For{
					Init: &ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("0"),
						},
					},
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenLessThan,
						Right: asttest.NewLiteralNumber("10"),
					},
					Next: &ast.Unary{
						Op:   lexer.TokenIncrement,
						Expr: &ast.Identifier{Name: "a"},
					},
					Statements: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
							Arguments: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				// a = 0
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},

				// a < 10
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "number10",
				},
				&vm.LessThanNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        8,
				},

				// print(a)
				&vm.Print{
					Arguments: []vm.Register{"a"},
				},

				// ++a
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1",
				},
				&vm.Add{
					Left:   "a",
					Right:  "4",
					Result: "a",
				},
				&vm.Jump{
					To: 2,
				},
			},
		},
		"for-continue-1": {
			nodes: []ast.Node{
				&ast.For{
					Init: &ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("0"),
						},
					},
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenLessThan,
						Right: asttest.NewLiteralNumber("10"),
					},
					Next: &ast.Unary{
						Op:   lexer.TokenIncrement,
						Expr: &ast.Identifier{Name: "a"},
					},
					Statements: []ast.Node{
						&ast.Continue{},
					},
				},
			},
			expected: []vm.Instruction{
				// a = 0
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},

				// a < 10
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "number10",
				},
				&vm.LessThanNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        8,
				},

				// continue
				&vm.Jump{
					To: 5,
				},

				// ++a
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1",
				},
				&vm.Add{
					Left:   "a",
					Right:  "4",
					Result: "a",
				},
				&vm.Jump{
					To: 2,
				},
			},
		},
		"empty-for-without-condition": {
			nodes: []ast.Node{
				&ast.For{},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "booltrue",
				},
				&vm.JumpUnless{
					Condition: "1",
					To:        2,
				},
				&vm.Jump{
					To: -1,
				},
			},
		},
		"for-true-with-statements": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1"),
					},
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "bar"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("2"),
							},
						},
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "baz"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("3"),
							},
						},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "qux"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("4"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number1",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "booltrue",
				},
				&vm.JumpUnless{
					Condition: "2",
					To:        8,
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number2",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "3",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number3",
				},
				&vm.Assign{
					Result:   "baz",
					Register: "4",
				},
				&vm.Jump{
					To: 1,
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number4",
				},
				&vm.Assign{
					Result:   "qux",
					Register: "5",
				},
			},
		},
		"for-condition": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.For{
					Condition: asttest.NewBinary(
						&ast.Identifier{Name: "i"},
						lexer.TokenLessThan,
						asttest.NewLiteralNumber("10"),
					),
					Statements: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "i"},
							},
							Rights: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "i"},
									lexer.TokenPlus,
									asttest.NewLiteralNumber("1"),
								),
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.Assign{
					Result:   "i",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "number10",
				},
				&vm.LessThanNumber{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        8,
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1",
				},
				&vm.Add{
					Left:   "i",
					Right:  "4",
					Result: "5",
				},
				&vm.Assign{
					Result:   "i",
					Register: "5",
				},
				&vm.Jump{
					To: 2,
				},
			},
		},
		"for-break": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "i"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("1"),
							},
						},
						&ast.Break{},
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "i"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("2"),
							},
						},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("3"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.Assign{
					Result:   "i",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "booltrue",
				},
				&vm.JumpUnless{
					Condition: "2",
					To:        9,
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number1",
				},
				&vm.Assign{
					Result:   "i",
					Register: "3",
				},
				&vm.Jump{
					To: 9,
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2",
				},
				&vm.Assign{
					Result:   "i",
					Register: "4",
				},
				&vm.Jump{
					To: 1,
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number3",
				},
				&vm.Assign{
					Result:   "i",
					Register: "5",
				},
			},
		},
		"for-continue": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "i"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("1"),
							},
						},
						&ast.Continue{},
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "i"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("2"),
							},
						},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("3"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.Assign{
					Result:   "i",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "booltrue",
				},
				&vm.JumpUnless{
					Condition: "2",
					To:        9,
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number1",
				},
				&vm.Assign{
					Result:   "i",
					Register: "3",
				},
				&vm.Jump{
					To: 1,
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2",
				},
				&vm.Assign{
					Result:   "i",
					Register: "4",
				},
				&vm.Jump{
					To: 1,
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number3",
				},
				&vm.Assign{
					Result:   "i",
					Register: "5",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{
					Types:   map[vm.TypeRegister]*types.Type{},
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
				}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
