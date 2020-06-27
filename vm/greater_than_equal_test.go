package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestGreaterThanEqualString_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"foo-foo": {
			ast.NewLiteralString("foo"), ast.NewLiteralString("foo"),
			"true"},
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
			ins := &vm.GreaterThanEqualString{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestGreaterThanEqualNumber_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"1-1.0": {
			ast.NewLiteralNumber("1"), ast.NewLiteralNumber("1.0"),
			"true"},
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
			ins := &vm.GreaterThanEqualNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
