package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"
)

type resultKindPair struct {
	result, kind string
}

func compileAssign(compiledFunc *vm.CompiledFunc, node *ast.Assign, fns map[string]*ast.Func) error {
	// First evaluate all the right expressions.
	var rightResults []resultKindPair
	for _, r := range node.Rights {
		right, rightKind, err := compileExpr(compiledFunc, r, fns)
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
			if v, ok := compiledFunc.Variables[variableName]; ok && rightResults[0].kind != v {
				return fmt.Errorf(
					"cannot assign %s to variable %s (expecting %s)",
					rightResults[0].kind, variableName, v)
			}

			compiledFunc.NewVariable(variableName, rr.kind)

			// When in an object we need to change an assign into a MapSet for
			// this instance variable.
			if compiledFunc.ObjectRegister != "" {
				keyRegister := compiledFunc.NextRegister()
				compiledFunc.Append(&vm.Assign{
					VariableName: keyRegister,
					Value:        asttest.NewLiteralString(variableName),
				})

				compiledFunc.Append(&vm.MapSet{
					Map:   compiledFunc.ObjectRegister,
					Key:   keyRegister,
					Value: rr.result,
				})
			} else {
				ins := &vm.Assign{
					VariableName: variableName,
					Register:     rr.result,
				}
				compiledFunc.Append(ins)
			}

		case *ast.Key:
			arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc,
				l.Expr, fns)
			if err != nil {
				return err
			}

			// TODO(elliot): Check this is a sane operation.
			keyResults, _, err := compileExpr(compiledFunc, l.Key, fns)
			if err != nil {
				return err
			}

			if strings.HasPrefix(arrayOrMapKind[0], "[]") {
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
