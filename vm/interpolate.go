package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Interpolate combines strings and expressions into one string result.
type Interpolate struct {
	Result string
	Args   []string
}

// Execute implements the Instruction interface for the VM.
func (ins *Interpolate) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	s := ""

	for _, arg := range ins.Args {
		s += renderLiteral(registers[arg], false)
	}

	registers[ins.Result] = asttest.NewLiteralString(s)

	return nil
}

func renderLiteral(literal *ast.Literal, asJSON bool) string {
	switch literal.Kind {
	case "char", "string", "data":
		if asJSON {
			// TODO(elliot): This is not escaped correctly.
			return fmt.Sprintf(`"%s"`, literal.Value)
		}

	case "number":
		return number.Format(number.NewNumber(literal.Value), -1)
	}

	return literal.Value
}
