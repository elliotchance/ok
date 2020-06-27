package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"

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
			ins := &vm.Concat{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
