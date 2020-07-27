package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
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
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralString(test.left),
				"1": asttest.NewLiteralString(test.right),
			}
			ins := &vm.Concat{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestConcat_String(t *testing.T) {
	ins := &vm.Concat{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = concat($1, $2)", ins.String())
}
