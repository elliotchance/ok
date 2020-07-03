package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestLessThanEqualString_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"foo-foo": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("foo"),
			"true"},
		"foo-bar": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("bar"),
			"false",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &vm.LessThanEqualString{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
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
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.0"),
			"true"},
		"1-1.1": {
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.1"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &vm.LessThanEqualNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
