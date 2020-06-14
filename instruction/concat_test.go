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
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralString(test.left),
				"1": ast.NewLiteralString(test.right),
			}
			ins := &instruction.Concat{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
