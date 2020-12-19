package vm

import (
	"fmt"
	"sort"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
)

// Props returns the property names of an object.
type Props struct {
	Value  Register // In: object or map
	Result Register // Out: []string
}

// Execute implements the Instruction interface for the VM.
func (ins *Props) Execute(_ *int, vm *VM) error {
	value := vm.Get(ins.Value)

	if value.Kind.Kind == types.KindMap {
		// Copy out the keys because we will be sorting them.
		keys := make([]*ast.Literal, len(value.Array))
		for i := 0; i < len(keys); i++ {
			keys[i] = value.Array[i]
		}

		sort.Slice(keys, func(i, j int) bool {
			return keys[i].Value < keys[j].Value
		})

		vm.Set(ins.Result, &ast.Literal{
			Kind:  types.StringArray,
			Array: keys,
		})

		return nil
	}

	if value.Kind.Kind == types.KindResolvedInterface ||
		value.Kind.Kind == types.KindUnresolvedInterface {
		// TODO(elliot): This should not allow KindUnresolvedInterface.

		var props []*ast.Literal
		for prop := range value.Map {
			if util.IsPublic(prop) {
				props = append(props, asttest.NewLiteralString(prop))
			}
		}

		sort.Slice(props, func(i, j int) bool {
			return props[i].Value < props[j].Value
		})

		vm.Set(ins.Result, &ast.Literal{
			Kind:  types.StringArray,
			Array: props,
		})

		return nil
	}

	vm.Raise("cannot use props on non-object, received " + value.Kind.String())

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Props) String() string {
	return fmt.Sprintf("%s = reflect.Props(%s)", ins.Result, ins.Value)
}
