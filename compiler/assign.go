package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

type resultKindPair struct {
	result vm.Register
	kind   *types.Type
}

func compileAssign(
	compiledFunc *vm.CompiledFunc,
	node *ast.Assign,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	// First evaluate all the right expressions.
	var rightResults []resultKindPair
	for _, r := range node.Rights {
		right, rightKind, err := compileExpr(compiledFunc, r, file, scopeOverrides)
		if err != nil {
			return err
		}

		// len(right) > 1 when function returns multiple values.
		for i, r2 := range right {
			rightResults = append(rightResults, resultKindPair{r2, rightKind[i]})
		}
	}

	if len(rightResults) != len(node.Lefts) {
		return fmt.Errorf("cannot assign %d values to %d variables",
			len(rightResults), len(node.Lefts))
	}

	// Perform final assigns.
	for i, rr := range rightResults {
		switch l := node.Lefts[i].(type) {
		case *ast.Identifier:
			variableName := l.Name

			// Make sure we do not assign the wrong type to an existing variable.
			if v, ok := compiledFunc.GetTypeForVariable(variableName, scopeOverrides); ok && v.String() != "any" && rightResults[0].kind.String() != v.String() {
				return fmt.Errorf(
					"%s cannot assign %s to variable %s (expecting %s)",
					node.Position(), rightResults[0].kind, variableName, v)
			}

			compiledFunc.NewVariable(variableName, rr.kind)

			compiledFunc.Append(&vm.Assign{
				VariableName: vm.Register(variableName),
				Register:     rr.result,
			})

		case *ast.Key:
			arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc,
				l.Expr, file, scopeOverrides)
			if err != nil {
				return err
			}

			// TODO(elliot): Check this is a sane operation.
			key := l.Key
			if e, ok := key.(*ast.Identifier); ok {
				key = asttest.NewLiteralString(e.Name)
			}
			keyResults, _, err := compileExpr(compiledFunc, key, file, scopeOverrides)
			if err != nil {
				return err
			}

			if arrayOrMapKind[0].Kind == types.KindArray {
				ins := &vm.ArraySet{
					Array: arrayOrMapResults[0],
					Index: keyResults[0],
					Value: rightResults[0].result,
				}
				compiledFunc.Append(ins)
			} else {
				ins := &vm.MapSet{
					Map:   arrayOrMapResults[0],
					Key:   keyResults[0],
					Value: rightResults[0].result,
				}
				compiledFunc.Append(ins)
			}
		}
	}

	return nil
}
