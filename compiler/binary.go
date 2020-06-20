package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
	"ok/lexer"
	"strings"
)

func getBinaryInstruction(op string, left, right, result string) (instruction.Instruction, string) {
	switch op {
	case "data + data":
		return &instruction.Combine{Left: left, Right: right, Result: result}, "data"

	case "data += data":
		return &instruction.Combine{Left: left, Right: right, Result: left}, "data"

	case "number + number":
		return &instruction.Add{Left: left, Right: right, Result: result}, "number"

	case "number += number":
		return &instruction.Add{Left: left, Right: right, Result: left}, "number"

	case "number - number":
		return &instruction.Subtract{Left: left, Right: right, Result: result}, "number"

	case "number -= number":
		return &instruction.Subtract{Left: left, Right: right, Result: left}, "number"

	case "number * number":
		return &instruction.Multiply{Left: left, Right: right, Result: result}, "number"

	case "number *= number":
		return &instruction.Multiply{Left: left, Right: right, Result: left}, "number"

	case "number / number":
		return &instruction.Divide{Left: left, Right: right, Result: result}, "number"

	case "number /= number":
		return &instruction.Divide{Left: left, Right: right, Result: left}, "number"

	case "number % number":
		return &instruction.Remainder{Left: left, Right: right, Result: result}, "number"

	case "number %= number":
		return &instruction.Remainder{Left: left, Right: right, Result: left}, "number"

	case "string + string":
		return &instruction.Concat{Left: left, Right: right, Result: result}, "string"

	case "string += string":
		return &instruction.Concat{Left: left, Right: right, Result: left}, "string"

	case "bool and bool":
		return &instruction.And{Left: left, Right: right, Result: result}, "bool"

	case "bool or bool":
		return &instruction.Or{Left: left, Right: right, Result: result}, "bool"

	case "bool == bool", "char == char", "data == data", "string == string":
		return &instruction.Equal{Left: left, Right: right, Result: result}, "bool"

	case "number == number":
		return &instruction.EqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "bool != bool", "char != char", "data != data", "string != string":
		return &instruction.NotEqual{Left: left, Right: right, Result: result}, "bool"

	case "number != number":
		return &instruction.NotEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "number > number":
		return &instruction.GreaterThanNumber{Left: left, Right: right, Result: result}, "bool"

	case "string > string":
		return &instruction.GreaterThanString{Left: left, Right: right, Result: result}, "bool"

	case "number < number":
		return &instruction.LessThanNumber{Left: left, Right: right, Result: result}, "bool"

	case "string < string":
		return &instruction.LessThanString{Left: left, Right: right, Result: result}, "bool"

	case "number >= number":
		return &instruction.GreaterThanEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "string >= string":
		return &instruction.GreaterThanEqualString{Left: left, Right: right, Result: result}, "bool"

	case "number <= number":
		return &instruction.LessThanEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "string <= string":
		return &instruction.LessThanEqualString{Left: left, Right: right, Result: result}, "bool"
	}

	return nil, ""
}

func compileBinary(compiledFunc *CompiledFunc, node *ast.Binary) (string, string, error) {
	if node.Op == lexer.TokenAssign ||
		node.Op == lexer.TokenPlusAssign ||
		node.Op == lexer.TokenMinusAssign ||
		node.Op == lexer.TokenTimesAssign ||
		node.Op == lexer.TokenDivideAssign ||
		node.Op == lexer.TokenRemainderAssign {

		right, rightKind, err := compileExpr(compiledFunc, node.Right)
		if err != nil {
			return "", "", err
		}

		// TODO(elliot): Check +=, etc.
		if key, ok := node.Left.(*ast.Key); ok {
			arrayOrMapResult, arrayOrMapKind, err := compileExpr(compiledFunc, key.Expr)
			if err != nil {
				return "", "", err
			}

			// TODO(elliot): Check this is a sane operation.
			keyResult, _, err := compileExpr(compiledFunc, key.Key)
			if err != nil {
				return "", "", err
			}

			if strings.HasPrefix(arrayOrMapKind, "[]") {
				ins := &instruction.ArraySet{
					Array: arrayOrMapResult,
					Index: keyResult,
					Value: right,
				}
				compiledFunc.append(ins)
			} else {
				ins := &instruction.MapSet{
					Map:   arrayOrMapResult,
					Key:   keyResult,
					Value: right,
				}
				compiledFunc.append(ins)
			}

			// TODO(elliot): Return something more reasonable here.
			return "", "", nil
		}

		variable, ok := node.Left.(*ast.Identifier)
		if !ok {
			return "", "", fmt.Errorf("cannot assign to non-variable")
		}

		// Make sure we do not assign the wrong type to an existing variable.
		if v, ok := compiledFunc.variables[variable.Name]; ok && rightKind != v {
			return "", "", fmt.Errorf(
				"cannot assign %s to variable %s (expecting %s)",
				rightKind, variable.Name, v)
		}

		returns := right
		if node.Op != lexer.TokenAssign {
			returns = compiledFunc.nextRegister()
		}

		switch node.Op {
		case lexer.TokenPlusAssign:
			switch rightKind {
			case "data":
				compiledFunc.append(&instruction.Combine{
					Left:   variable.Name,
					Right:  right,
					Result: returns,
				})

			case "number":
				compiledFunc.append(&instruction.Add{
					Left:   variable.Name,
					Right:  right,
					Result: returns,
				})

			case "string":
				compiledFunc.append(&instruction.Concat{
					Left:   variable.Name,
					Right:  right,
					Result: returns,
				})
			}

		case lexer.TokenMinusAssign:
			compiledFunc.append(&instruction.Subtract{
				Left:   variable.Name,
				Right:  right,
				Result: returns,
			})

		case lexer.TokenTimesAssign:
			compiledFunc.append(&instruction.Multiply{
				Left:   variable.Name,
				Right:  right,
				Result: returns,
			})

		case lexer.TokenDivideAssign:
			compiledFunc.append(&instruction.Divide{
				Left:   variable.Name,
				Right:  right,
				Result: returns,
			})

		case lexer.TokenRemainderAssign:
			compiledFunc.append(&instruction.Remainder{
				Left:   variable.Name,
				Right:  right,
				Result: returns,
			})
		}

		ins := &instruction.Assign{
			VariableName: variable.Name,
			Register:     returns,
		}
		compiledFunc.append(ins)
		compiledFunc.newVariable(variable.Name, rightKind)

		return variable.Name, rightKind, nil
	}

	left, leftKind, err := compileExpr(compiledFunc, node.Left)
	if err != nil {
		return "", "", err
	}

	right, rightKind, err := compileExpr(compiledFunc, node.Right)
	if err != nil {
		return "", "", err
	}

	returns := compiledFunc.nextRegister()

	op := fmt.Sprintf("%s %s %s", leftKind, node.Op, rightKind)
	if bop, kind := getBinaryInstruction(op, left, right, returns); bop != nil {
		// TODO(elliot): It would be nice to be able to evaluate expressions
		//  involving literals here. So, 1 + 1 just becomes 2.

		compiledFunc.append(bop)

		return returns, kind, nil
	}

	return "", "", fmt.Errorf("cannot perform %s", op)
}
