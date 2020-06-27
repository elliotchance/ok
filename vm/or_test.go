package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestOr_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right bool
		expected    string
	}{
		"false-false": {false, false, "false"},
		"false-true":  {false, true, "true"},
		"true-false":  {true, false, "true"},
		"true-true":   {true, true, "true"},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralBool(test.left),
				"1": ast.NewLiteralBool(test.right),
			}
			ins := &vm.Or{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
