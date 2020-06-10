package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotEqual_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"false-false": {
			ast.NewLiteralBool(false), ast.NewLiteralBool(false),
			"false"},
		"false-true": {
			ast.NewLiteralBool(false), ast.NewLiteralBool(true),
			"true",
		},
		"a-a": {
			ast.NewLiteralChar('a'), ast.NewLiteralChar('a'),
			"false"},
		"b-B": {
			ast.NewLiteralChar('b'), ast.NewLiteralChar('B'),
			"true",
		},
		"d1-d1": {
			ast.NewLiteralData([]byte("d1")), ast.NewLiteralData([]byte("d1")),
			"false"},
		"d1-d2": {
			ast.NewLiteralData([]byte("d1")), ast.NewLiteralData([]byte("d2")),
			"true",
		},
		"foo-foo": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("foo"),
			"false"},
		"foo-bar": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("bar"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			ins := &instruction.NotEqual{Left: test.left, Right: test.right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}

func TestNotEqualNumber_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"1-1.0": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.0"),
			"false"},
		"1-1.1": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.1"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			ins := &instruction.NotEqualNumber{Left: test.left, Right: test.right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
