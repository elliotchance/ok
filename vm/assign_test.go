package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestAssign_Execute(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		registers := map[string]*ast.Literal{}
		ins := &vm.Assign{VariableName: "1", Value: asttest.NewLiteralNumber("1.5")}
		assert.NoError(t, ins.Execute(registers, nil, nil))
		assert.Equal(t, "1.5", registers["1"].Value)
	})

	t.Run("register", func(t *testing.T) {
		registers := map[string]*ast.Literal{
			"0": asttest.NewLiteralNumber("1.5"),
		}
		ins := &vm.Assign{VariableName: "1", Register: "0"}
		assert.NoError(t, ins.Execute(registers, nil, nil))
		assert.Equal(t, "1.5", registers["1"].Value)
	})
}
