package vm_test

import (
	"ok/ast"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left     bool
		expected string
	}{
		"false": {false, "true"},
		"true":  {true, "false"},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralBool(test.left),
			}
			ins := &vm.Not{Left: "0", Result: "1"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}