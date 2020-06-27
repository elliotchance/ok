package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
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
						ast.NewLiteralData([]byte("foo")),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralData([]byte("bar")),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("foo")),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("bar")),
				},
				&vm.Combine{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralString("foo"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralString("bar"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
				},
				&vm.Concat{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Add{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenMinusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Subtract{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenTimesAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Multiply{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenDivideAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Divide{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
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
						ast.NewLiteralNumber("0"),
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenRemainderAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Remainder{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"string-divide-number": {
			nodes: []ast.Node{
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Binary{
							Left: &ast.Literal{
								Kind:  "string",
								Value: "foo",
							},
							Op: lexer.TokenDivide,
							Right: &ast.Literal{
								Kind:  "number",
								Value: "123",
							},
						},
					},
				},
			},
			err: errors.New("cannot perform string / number"),
		},
		"data-plus-data": {
			nodes: []ast.Node{
				ast.NewBinary(
					ast.NewLiteralData([]byte("foo")),
					lexer.TokenPlus,
					ast.NewLiteralData([]byte("bar")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("foo")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("bar")),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.20"),
					lexer.TokenPlus,
					ast.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
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
				ast.NewBinary(
					ast.NewLiteralString("foo"),
					lexer.TokenPlus,
					ast.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.20"),
					lexer.TokenMinus,
					ast.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.20"),
					lexer.TokenTimes,
					ast.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.20"),
					lexer.TokenDivide,
					ast.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("5"),
					lexer.TokenRemainder,
					ast.NewLiteralNumber("1.20"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1.20"),
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
						ast.NewLiteralNumber("1.5"),
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						ast.NewLiteralString("1.5"),
					},
				},
			},
			err: errors.New("cannot assign string to variable foo (expecting number)"),
		},
		"bool-greater-than-bool": {
			nodes: []ast.Node{
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenGreaterThan,
					ast.NewLiteralBool(true),
				),
			},
			err: errors.New("cannot perform bool > bool"),
		},
		"bool-equals-bool": {
			nodes: []ast.Node{
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenEqual,
					ast.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
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
				ast.NewBinary(
					ast.NewLiteralChar('a'),
					lexer.TokenEqual,
					ast.NewLiteralChar('B'),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralChar('a'),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralChar('B'),
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
				ast.NewBinary(
					ast.NewLiteralData([]byte("a")),
					lexer.TokenEqual,
					ast.NewLiteralData([]byte("B")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("a")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("B")),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenEqual,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("foo"),
					lexer.TokenEqual,
					ast.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
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
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenNotEqual,
					ast.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
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
				ast.NewBinary(
					ast.NewLiteralChar('a'),
					lexer.TokenNotEqual,
					ast.NewLiteralChar('B'),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralChar('a'),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralChar('B'),
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
				ast.NewBinary(
					ast.NewLiteralData([]byte("a")),
					lexer.TokenNotEqual,
					ast.NewLiteralData([]byte("B")),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("a")),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("B")),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenNotEqual,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("foo"),
					lexer.TokenNotEqual,
					ast.NewLiteralString("bar"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenGreaterThan,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenGreaterThanEqual,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenLessThan,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralNumber("1.0"),
					lexer.TokenLessThanEqual,
					ast.NewLiteralNumber("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("1.0"),
					lexer.TokenGreaterThan,
					ast.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("1.0"),
					lexer.TokenGreaterThanEqual,
					ast.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("1.0"),
					lexer.TokenLessThan,
					ast.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
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
				ast.NewBinary(
					ast.NewLiteralString("1.0"),
					lexer.TokenLessThanEqual,
					ast.NewLiteralString("1"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
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
				ast.NewLiteralString("total is"),
				ast.NewBinary(
					ast.NewLiteralNumber("1.5"),
					lexer.TokenPlus,
					ast.NewLiteralNumber("0.8"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("total is"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0.8"),
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
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenAnd,
					ast.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
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
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenAnd,
					ast.NewLiteralBool(false),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(false),
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
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenOr,
					ast.NewLiteralBool(true),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
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
				ast.NewBinary(
					ast.NewLiteralBool(true),
					lexer.TokenOr,
					ast.NewLiteralBool(false),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(false),
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
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Binary{
							Left: &ast.Literal{
								Kind:  "bool",
								Value: "true",
							},
							Op: lexer.TokenPlus,
							Right: &ast.Literal{
								Kind:  "bool",
								Value: "false",
							},
						},
					},
				},
			},
			err: errors.New("cannot perform bool + bool"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...), nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
