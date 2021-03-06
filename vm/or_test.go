package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
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
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralBool(test.left),
				"1": asttest.NewLiteralBool(test.right),
			}
			ins := &vm.Or{Left: "0", Right: "1", Result: "2"}
			vm := &vm.VM{
				Stack: []map[vm.Register]*ast.Literal{registers},
			}
			assert.NoError(t, ins.Execute(nil, vm))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
