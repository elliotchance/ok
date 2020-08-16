package vm

import (
	"fmt"
	"sort"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/util"
)

// Props returns the property names of an object.
type Props struct {
	Object, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Props) Execute(_ *int, vm *VM) error {
	object := vm.Get(ins.Object)

	if kind.IsObject(object.Kind) {
		var props []*ast.Literal
		for prop := range object.Map {
			if util.IsPublic(prop) {
				props = append(props, asttest.NewLiteralString(prop))
			}
		}

		sort.Slice(props, func(i, j int) bool {
			return props[i].Value < props[j].Value
		})

		vm.Set(ins.Result, &ast.Literal{
			Kind:  "[]string",
			Array: props,
		})

		return nil
	}

	vm.Raise("cannot use props on non-object, received " + object.Kind)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Props) String() string {
	return fmt.Sprintf("%s = reflect.Props(%s)", ins.Result, ins.Object)
}
