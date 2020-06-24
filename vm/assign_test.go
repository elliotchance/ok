package vm_test

import (
	"ok/ast"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssign_Execute(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		registers := map[string]*ast.Literal{}
		ins := &vm.Assign{VariableName: "1", Value: ast.NewLiteralNumber("1.5")}
		assert.NoError(t, ins.Execute(registers, nil, nil))
		assert.Equal(t, "1.5", registers["1"].Value)
	})

	t.Run("register", func(t *testing.T) {
		registers := map[string]*ast.Literal{
			"0": ast.NewLiteralNumber("1.5"),
		}
		ins := &vm.Assign{VariableName: "1", Register: "0"}
		assert.NoError(t, ins.Execute(registers, nil, nil))
		assert.Equal(t, "1.5", registers["1"].Value)
	})
}