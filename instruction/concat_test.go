package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
	}{
		"both-empty":     {"", "", ""},
		"both-non-empty": {"foo", "bar", "foobar"},
	} {
		t.Run(testName, func(t *testing.T) {
			left := ast.NewLiteralString(test.left)
			right := ast.NewLiteralString(test.right)
			ins := &instruction.Concat{Left: left, Right: right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
