package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestNotEqual_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"false-false": {
			asttest.NewLiteralBool(false), asttest.NewLiteralBool(false),
			"false"},
		"false-true": {
			asttest.NewLiteralBool(false), asttest.NewLiteralBool(true),
			"true",
		},
		"a-a": {
			asttest.NewLiteralChar('a'), asttest.NewLiteralChar('a'),
			"false"},
		"b-B": {
			asttest.NewLiteralChar('b'), asttest.NewLiteralChar('B'),
			"true",
		},
		"d1-d1": {
			asttest.NewLiteralData([]byte("d1")), asttest.NewLiteralData([]byte("d1")),
			"false"},
		"d1-d2": {
			asttest.NewLiteralData([]byte("d1")), asttest.NewLiteralData([]byte("d2")),
			"true",
		},
		"foo-foo": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("foo"),
			"false"},
		"foo-bar": {
			asttest.NewLiteralString("foo"), asttest.NewLiteralString("bar"),
			"true",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": test.left,
				"1": test.right,
			}
			ins := &vm.NotEqual{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestNotEqualNumber_Execute(t *testing.T) {
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
			ins := &vm.NotEqualNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestNotEqual_String(t *testing.T) {
	ins := &vm.NotEqual{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = $1 != $2", ins.String())
}

func TestNotEqualNumber_String(t *testing.T) {
	ins := &vm.NotEqualNumber{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = $1 != $2", ins.String())
}
