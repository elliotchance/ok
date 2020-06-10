package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func getBinaryInstruction(op string, left, right, result string) instruction.Instruction {
	switch op {
	case "data + data":
		return &instruction.Combine{Left: left, Right: right, Result: result}

	case "number + number":
		return &instruction.Add{Left: left, Right: right, Result: result}

	case "number - number":
		return &instruction.Subtract{Left: left, Right: right, Result: result}

	case "number * number":
		return &instruction.Multiply{Left: left, Right: right, Result: result}

	case "number / number":
		return &instruction.Divide{Left: left, Right: right, Result: result}

	case "number % number":
		return &instruction.Remainder{Left: left, Right: right, Result: result}

	case "string + string":
		return &instruction.Concat{Left: left, Right: right, Result: result}

	case "bool and bool":
		return &instruction.And{Left: left, Right: right, Result: result}

	case "bool or bool":
		return &instruction.Or{Left: left, Right: right, Result: result}

	case "bool == bool", "char == char", "data == data", "string == string":
		return &instruction.Equal{Left: left, Right: right, Result: result}

	case "number == number":
		return &instruction.EqualNumber{Left: left, Right: right, Result: result}

	case "bool != bool", "char != char", "data != data", "string != string":
		return &instruction.NotEqual{Left: left, Right: right, Result: result}

	case "number != number":
		return &instruction.NotEqualNumber{Left: left, Right: right, Result: result}

	case "number > number":
		return &instruction.GreaterThanNumber{Left: left, Right: right, Result: result}

	case "string > string":
		return &instruction.GreaterThanString{Left: left, Right: right, Result: result}

	case "number < number":
		return &instruction.LessThanNumber{Left: left, Right: right, Result: result}

	case "string < string":
		return &instruction.LessThanString{Left: left, Right: right, Result: result}

	case "number >= number":
		return &instruction.GreaterThanEqualNumber{Left: left, Right: right, Result: result}

	case "string >= string":
		return &instruction.GreaterThanEqualString{Left: left, Right: right, Result: result}

	case "number <= number":
		return &instruction.LessThanEqualNumber{Left: left, Right: right, Result: result}

	case "string <= string":
		return &instruction.LessThanEqualString{Left: left, Right: right, Result: result}
	}

	return nil
}

func compileBinary(compiledFunc *CompiledFunc, node *ast.Binary) (string, string, error) {
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
	if bop := getBinaryInstruction(op, left, right, returns); bop != nil {
		// TODO(elliot): It would be nice to be able to evaluate expressions
		//  involving literals here. So, 1 + 1 just becomes 2.

		compiledFunc.append(bop)

		return returns, leftKind, nil
	}

	return "", "", fmt.Errorf("cannot perform %s", op)
}
