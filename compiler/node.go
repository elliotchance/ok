package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func getBinaryInstruction(op string, left, right *ast.Literal) instruction.Instruction {
	switch op {
	case "data + data":
		return &instruction.Combine{Left: left, Right: right}

	case "number + number":
		return &instruction.Add{Left: left, Right: right}

	case "number - number":
		return &instruction.Subtract{Left: left, Right: right}

	case "number * number":
		return &instruction.Multiply{Left: left, Right: right}

	case "number / number":
		return &instruction.Divide{Left: left, Right: right}

	case "number % number":
		return &instruction.Remainder{Left: left, Right: right}

	case "string + string":
		return &instruction.Concat{Left: left, Right: right}

	case "bool and bool":
		return &instruction.And{Left: left, Right: right}

	case "bool or bool":
		return &instruction.Or{Left: left, Right: right}

	case "bool == bool", "char == char", "data == data", "string == string":
		return &instruction.Equal{Left: left, Right: right}

	case "number == number":
		return &instruction.EqualNumber{Left: left, Right: right}

	case "bool != bool", "char != char", "data != data", "string != string":
		return &instruction.NotEqual{Left: left, Right: right}

	case "number != number":
		return &instruction.NotEqualNumber{Left: left, Right: right}

	case "number > number":
		return &instruction.GreaterThanNumber{Left: left, Right: right}

	case "string > string":
		return &instruction.GreaterThanString{Left: left, Right: right}

	case "number < number":
		return &instruction.LessThanNumber{Left: left, Right: right}

	case "string < string":
		return &instruction.LessThanString{Left: left, Right: right}

	case "number >= number":
		return &instruction.GreaterThanEqualNumber{Left: left, Right: right}

	case "string >= string":
		return &instruction.GreaterThanEqualString{Left: left, Right: right}

	case "number <= number":
		return &instruction.LessThanEqualNumber{Left: left, Right: right}

	case "string <= string":
		return &instruction.LessThanEqualString{Left: left, Right: right}
	}

	return nil
}

func compileNode(node ast.Node) (*ast.Literal, error) {
	switch n := node.(type) {
	case *ast.Literal:
		return n, nil

	case *ast.Group:
		return compileNode(n.Expr)

	case *ast.Unary:
		// TODO(elliot): Need "-"
		val, err := compileNode(n.Expr)
		if err != nil {
			return nil, err
		}

		return ast.NewLiteralBool(val.Value == "false"), nil

	case *ast.Binary:
		leftType := typeOf(n.Left)
		rightType := typeOf(n.Right)
		op := fmt.Sprintf("%s %s %s", leftType, n.Op, rightType)

		left, err := compileNode(n.Left)
		if err != nil {
			return nil, err
		}

		right, err := compileNode(n.Right)
		if err != nil {
			return nil, err
		}

		if bop := getBinaryInstruction(op, left, right); bop != nil {
			err := bop.Execute()
			if err != nil {
				return nil, err
			}

			return bop.Answer(), nil
		}

		return nil, fmt.Errorf("cannot perform %s", op)
	}

	// TODO(elliot): Do something more sensible here.
	panic(node)
}

func typeOf(node ast.Node) string {
	switch n := node.(type) {
	case *ast.Literal:
		return n.Kind

	case *ast.Binary:
		return typeOf(n.Left)

	case *ast.Group:
		return typeOf(n.Expr)
	}

	// TODO(elliot): Do something more sensible here.
	panic(node)
}
