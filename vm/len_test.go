package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestLen_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		x        *ast.Literal
		expected *ast.Literal
	}{
		// TODO(elliot): Tests for array and map.
		"string": {
			asttest.NewLiteralString("foo bar"),
			asttest.NewLiteralNumber("7"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{
				"0": test.x,
			}
			ins := &vm.Len{Argument: "0", Result: "1"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result])
		})
	}
}
