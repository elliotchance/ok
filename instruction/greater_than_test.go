package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreaterThanString_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"foo-foo": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("foo"),
			"false"},
		"foo-bar": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("bar"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &instruction.GreaterThanString{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestGreaterThanNumber_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"1-1.0": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.0"),
			"false"},
		"1-1.1": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.1"),
			"false",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &instruction.GreaterThanNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
