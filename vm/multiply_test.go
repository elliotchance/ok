package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestMultiply_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
	}{
		"success": {"1.2200", "4.7", "5.734"},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralNumber(test.left),
				"1": asttest.NewLiteralNumber(test.right),
			}
			ins := &vm.Multiply{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			actual := number.NewNumber(registers[ins.Result].Value)
			assert.Equal(t, test.expected, number.Format(actual, -1))
		})
	}
}
