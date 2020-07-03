package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestEqual_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right *ast.Literal
		expected    string
	}{
		"false-false": {
			asttest.NewLiteralBool(false), asttest.NewLiteralBool(false),
			"true"},
		"false-true": {
			asttest.NewLiteralBool(false), asttest.NewLiteralBool(true),
			"false",
		},
		"a-a": {
			asttest.NewLiteralChar('a'), asttest.NewLiteralChar('a'),
			"true"},
		"b-B": {
			asttest.NewLiteralChar('b'), asttest.NewLiteralChar('B'),
			"false",
		},
		"d1-d1": {
			asttest.NewLiteralData([]byte("d1")), asttest.NewLiteralData([]byte("d1")),
			"true"},
		"d1-d2": {
			asttest.NewLiteralData([]byte("d1")), asttest.NewLiteralData([]byte("d2")),
			"false",
		},
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
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.0"),
			"true"},
		"1-1.1": {
			asttest.NewLiteralNumber("1"), asttest.NewLiteralNumber("1.1"),
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
