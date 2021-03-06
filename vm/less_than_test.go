package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestLessThanString_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"foo-foo": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("foo"),
			"false"},
		"foo-bar": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("bar"),
			"false",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &vm.LessThanString{Left: "0", Right: "1", Result: "2"}
			vm := &vm.VM{
				Stack: []map[vm.Register]*ast.Literal{registers},
			}
			assert.NoError(t, ins.Execute(nil, vm))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestLessThanNumber_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"1-1.0": {
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.0"),
			"false"},
		"1-1.1": {
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.1"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &vm.LessThanNumber{Left: "0", Right: "1", Result: "2"}
			vm := &vm.VM{
				Stack: []map[vm.Register]*ast.Literal{registers},
			}
			assert.NoError(t, ins.Execute(nil, vm))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestLessThanString_String(t *testing.T) {
	ins := &vm.LessThanString{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = $1 < $2", ins.String())
}

func TestLessThanNumber_String(t *testing.T) {
	ins := &vm.LessThanNumber{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = $1 < $2", ins.String())
}
