package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessThanEqualString_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"foo-foo": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("foo"),
			"true"},
		"foo-bar": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("bar"),
			"false",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &instruction.LessThanEqualString{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestLessThanEqualNumber_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"1-1.0": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.0"),
			"true"},
		"1-1.1": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.1"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &instruction.LessThanEqualNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
