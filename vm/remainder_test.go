package vm_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

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
				"0": asttest.NewLiteralNumber(test.left),
				"1": asttest.NewLiteralNumber(test.right),
			}
			ins := &vm.Remainder{Left: "0", Right: "1", Result: "2"}
			assert.Equal(t, test.err, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
