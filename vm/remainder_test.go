package vm_test

import (
	"errors"
	"ok/ast"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemainder_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
		err         error
	}{
		"maintain-precision": {"1.2200", "4.7", "1.2200", nil},
		"divide-zero":        {"1.2200", "0", "0", errors.New("division by zero")},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralNumber(test.left),
				"1": ast.NewLiteralNumber(test.right),
			}
			ins := &vm.Remainder{Left: "0", Right: "1", Result: "2"}
			assert.Equal(t, test.err, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}