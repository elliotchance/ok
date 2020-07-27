package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestAnd_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right bool
		expected    string
	}{
		"false-false": {false, false, "false"},
		"false-true":  {false, true, "false"},
		"true-false":  {true, false, "false"},
		"true-true":   {true, true, "true"},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralBool(test.left),
				"1": asttest.NewLiteralBool(test.right),
			}
			ins := &vm.And{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestAnd_String(t *testing.T) {
	ins := &vm.And{Left: "0", Right: "1", Result: "2"}
	assert.Equal(t, "$2 = $0 and $1", ins.String())
}
