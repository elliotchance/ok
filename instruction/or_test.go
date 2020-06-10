package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

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
			left := ast.NewLiteralBool(test.left)
			right := ast.NewLiteralBool(test.right)
			ins := &instruction.Or{Left: left, Right: right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
