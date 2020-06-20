package vm_test

import (
	"ok/ast"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
	}{
		"maintain-precision": {"1.2200", "4.7", "-3.4800"},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralNumber(test.left),
				"1": ast.NewLiteralNumber(test.right),
			}
			ins := &vm.Subtract{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
