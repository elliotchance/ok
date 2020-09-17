package vm

import (
	"fmt"
	"sort"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
)

// Interpolate combines strings and expressions into one string result.
type Interpolate struct {
	Result Register
	Args   Registers
}

// Execute implements the Instruction interface for the VM.
func (ins *Interpolate) Execute(_ *int, vm *VM) error {
	s := ""

	for _, arg := range ins.Args {
		s += renderLiteral(vm.Get(arg), false)
	}

	vm.Set(ins.Result, asttest.NewLiteralString(s))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Interpolate) String() string {
	return fmt.Sprintf("%s = interpolate %s", ins.Result, ins.Args)
}

func renderLiteral(v *ast.Literal, asJSON bool) string {
	if v.Kind.Kind == types.KindFunc {
		return v.String()
	}

	if v.Kind.Kind == types.KindArray {
		s := "["
		for j, element := range v.Array {
			if j > 0 {
				s += ", "
			}

			s += renderLiteral(element, true)
		}

		return s + "]"
	}

	// Literals.
	switch v.Kind.Kind {
	case types.KindChar, types.KindString, types.KindData:
		if asJSON {
			// TODO(elliot): This is not escaped correctly.
			return fmt.Sprintf(`"%s"`, v.Value)
		}

		return v.Value

	case types.KindNumber:
		return number.Format(number.NewNumber(v.Value), -1)

	case types.KindBool:
		return v.Value
	}

	// Maps or objects are handled the same way. We can recognise maps with:
	//
	//   strings.HasPrefix(v.Kind, "{}")
	//
	// However, it's not trivial to identify objects. Doesn't matter, if none of
	// the above matched then it must be a map or object.

	// Keys are always sorted.
	var keys []string
	for key := range v.Map {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	s := "{"
	j := 0
	for _, key := range keys {
		// If it's an object we do not expose non-public entities.
		if v.Kind.Kind != types.KindMap && !util.IsPublic(key) {
			continue
		}

		// We do not render function literals. These would almost never be
		// useful in a JSON output.
		element := v.Map[key]
		if element.Kind.Kind == types.KindFunc {
			continue
		}

		if j > 0 {
			s += ", "
		}

		s += fmt.Sprintf(`"%s": `, key)
		s += renderLiteral(element, true)
		j++
	}

	return s + "}"
}
