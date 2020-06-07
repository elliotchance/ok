package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
	}{
		"maintain-precision": {"1.2200", "4.7", "5.9200"},
	} {
		t.Run(testName, func(t *testing.T) {
			left := ast.NewLiteralNumber(test.left)
			right := ast.NewLiteralNumber(test.right)
			ins := &instruction.Add{Left: left, Right: right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
