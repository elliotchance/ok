package compiler_test

import (
	"errors"
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

func TestBinary(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"plus-assign-data": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralData([]byte("foo")),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: asttest.NewLiteralData([]byte("bar")),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralData([]byte("foo")),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralData([]byte("bar")),
				},
				&vm.Combine{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"plus-assign-string": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralString("foo"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: asttest.NewLiteralString("bar"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.Concat{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"plus-assign-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Add{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"minus-assign-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenMinusAssign,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Subtract{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"times-assign-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenTimesAssign,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Multiply{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"divide-assign-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenDivideAssign,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Divide{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"remainder-assign-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenRemainderAssign,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Remainder{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"string-divide-number": {
			nodes: []ast.Node{
				&ast.Call{
					Expr: &ast.Identifier{Name: "print"},
					Arguments: []ast.Node{
						&ast.Binary{
							Left: &ast.Literal{
								Kind:  types.String,
								Value: "foo",
							},
							Op: lexer.TokenDivide,
							Right: &ast.Literal{
								Kind:  types.Number,
								Value: "123",
							},
						},
					},
				},
			},
			err: errors.New(" cannot perform string / number"),
		},
		"data-plus-data": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralData([]byte("foo")),
					lexer.TokenPlus,
					asttest.NewLiteralData([]byte("bar")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralData([]byte("foo")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralData([]byte("bar")),
				},
				&vm.Combine{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-plus-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.20"),
					lexer.TokenPlus,
					asttest.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-plus-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("foo"),
					lexer.TokenPlus,
					asttest.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.Concat{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-minus-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.20"),
					lexer.TokenMinus,
					asttest.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Subtract{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-times-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.20"),
					lexer.TokenTimes,
					asttest.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Multiply{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-divide-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.20"),
					lexer.TokenDivide,
					asttest.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Divide{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-remainder-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("5"),
					lexer.TokenRemainder,
					asttest.NewLiteralNumber("1.20"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1.20"),
				},
				&vm.Remainder{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"variable-bad-reassign": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1.5"),
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralString("1.5"),
					},
				},
			},
			err: errors.New(" cannot assign string to variable foo (expecting number)"),
		},
		"bool-greater-than-bool": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenGreaterThan,
					asttest.NewLiteralBool(true),
				),
			},
			err: errors.New(" cannot perform bool > bool"),
		},
		"bool-equals-bool": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenEqual,
					asttest.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"char-equals-char": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralChar('a'),
					lexer.TokenEqual,
					asttest.NewLiteralChar('B'),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralChar('a'),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralChar('B'),
				},
				&vm.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"data-equals-data": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralData([]byte("a")),
					lexer.TokenEqual,
					asttest.NewLiteralData([]byte("B")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralData([]byte("a")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralData([]byte("B")),
				},
				&vm.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-equals-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenEqual,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.EqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-equals-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("foo"),
					lexer.TokenEqual,
					asttest.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"bool-not-equals-bool": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenNotEqual,
					asttest.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"char-not-equals-char": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralChar('a'),
					lexer.TokenNotEqual,
					asttest.NewLiteralChar('B'),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralChar('a'),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralChar('B'),
				},
				&vm.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"data-not-equals-data": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralData([]byte("a")),
					lexer.TokenNotEqual,
					asttest.NewLiteralData([]byte("B")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralData([]byte("a")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralData([]byte("B")),
				},
				&vm.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-not-equals-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenNotEqual,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.NotEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-not-equals-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("foo"),
					lexer.TokenNotEqual,
					asttest.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-greater-than-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenGreaterThan,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.GreaterThanNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-greater-than-equal-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenGreaterThanEqual,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.GreaterThanEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-less-than-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenLessThan,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.LessThanNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"number-less-than-equal-number": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.0"),
					lexer.TokenLessThanEqual,
					asttest.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.LessThanEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-greater-than-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("1.0"),
					lexer.TokenGreaterThan,
					asttest.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("1"),
				},
				&vm.GreaterThanString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-greater-than-equal-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("1.0"),
					lexer.TokenGreaterThanEqual,
					asttest.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("1"),
				},
				&vm.GreaterThanEqualString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-less-than-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("1.0"),
					lexer.TokenLessThan,
					asttest.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("1"),
				},
				&vm.LessThanString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"string-less-than-equal-string": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralString("1.0"),
					lexer.TokenLessThanEqual,
					asttest.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralString("1"),
				},
				&vm.LessThanEqualString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"print-2": {
			nodes: []ast.Node{
				asttest.NewLiteralString("total is"),
				asttest.NewBinary(
					asttest.NewLiteralNumber("1.5"),
					lexer.TokenPlus,
					asttest.NewLiteralNumber("0.8"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("total is"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1.5"),
				},
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("0.8"),
				},
				&vm.Add{
					Left:   "2",
					Right:  "3",
					Result: "4",
				},
			},
		},
		"true-and-true": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenAnd,
					asttest.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.And{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"true-and-false": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenAnd,
					asttest.NewLiteralBool(false),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(false),
				},
				&vm.And{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"true-or-true": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenOr,
					asttest.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Or{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"true-or-false": {
			nodes: []ast.Node{
				asttest.NewBinary(
					asttest.NewLiteralBool(true),
					lexer.TokenOr,
					asttest.NewLiteralBool(false),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralBool(false),
				},
				&vm.Or{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"bool-plus-bool": {
			nodes: []ast.Node{
				&ast.Call{
					Expr: &ast.Identifier{Name: "print"},
					Arguments: []ast.Node{
						&ast.Binary{
							Left: &ast.Literal{
								Kind:  types.Bool,
								Value: "true",
							},
							Op: lexer.TokenPlus,
							Right: &ast.Literal{
								Kind:  types.Bool,
								Value: "false",
							},
						},
					},
				},
			},
			err: errors.New(" cannot perform bool + bool"),
		},
		"plus-assign-array": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewArrayNumbers(nil),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: asttest.NewArrayNumbers(nil),
				},
			},
			expected: []vm.Instruction{
				// first array
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   types.NumberArray,
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "2",
				},

				// second array
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.ArrayAlloc{
					Size:   "3",
					Result: "4",
					Kind:   types.NumberArray,
				},

				// +=
				&vm.Append{
					A:      "i",
					B:      "4",
					Result: "i",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
