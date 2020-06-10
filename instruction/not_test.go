package instruction_test

import (
	"ok/ast"
	"ok/instruction"
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
			left := ast.NewLiteralBool(test.left)
			ins := &instruction.Not{Left: left}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
