package vm_test

import (
	"ok/ast"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"false-false": {
			ast.NewLiteralBool(false), ast.NewLiteralBool(false),
			"true"},
		"false-true": {
			ast.NewLiteralBool(false), ast.NewLiteralBool(true),
			"false",
		},
		"a-a": {
			ast.NewLiteralChar('a'), ast.NewLiteralChar('a'),
			"true"},
		"b-B": {
			ast.NewLiteralChar('b'), ast.NewLiteralChar('B'),
			"false",
		},
		"d1-d1": {
			ast.NewLiteralData([]byte("d1")), ast.NewLiteralData([]byte("d1")),
			"true"},
		"d1-d2": {
			ast.NewLiteralData([]byte("d1")), ast.NewLiteralData([]byte("d2")),
			"false",
		},
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
			ins := &vm.Equal{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestEqualNumber_Execute(t *testing.T) {
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
			ins := &vm.EqualNumber{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
