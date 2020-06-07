package instruction_test

import (
	"errors"
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
		err         error
	}{
		"maintain-precision": {"1.2200", "4.7", "0.2596", nil},
		"divide-zero":        {"1.2200", "0", "0", errors.New("division by zero")},
	} {
		t.Run(testName, func(t *testing.T) {
			left := ast.NewLiteralNumber(test.left)
			right := ast.NewLiteralNumber(test.right)
			ins := &instruction.Divide{Left: left, Right: right}
			assert.Equal(t, test.err, ins.Execute())
			if test.err == nil {
				assert.Equal(t, test.expected, ins.Result.Value)
			}
		})
	}
}
