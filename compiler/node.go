package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

var binaryOperations = map[string]func(*ast.Literal, *ast.Literal) (*ast.Literal, error){
	"data + data": func(left, right *ast.Literal) (*ast.Literal, error) {
		return ast.NewLiteralData(
			append([]byte(left.Value), []byte(right.Value)...)), nil
	},
	"number + number": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Add{
			Left:  left,
			Right: right,
		}
		err := result.Execute()

		return ast.NewLiteralNumber(result.Result.Value), err
	},
	"number - number": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Subtract{
			Left:  left,
			Right: right,
		}
		err := result.Execute()

		return ast.NewLiteralNumber(result.Result.Value), err
	},
	"number * number": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Multiply{
			Left:  left,
			Right: right,
		}
		err := result.Execute()

		return ast.NewLiteralNumber(result.Result.Value), err
	},
	"number / number": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Divide{
			Left:  left,
			Right: right,
		}
		err := result.Execute()
		if err != nil {
			return nil, err
		}

		return ast.NewLiteralNumber(result.Result.Value), nil
	},
	"number % number": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Remainder{
			Left:  left,
			Right: right,
		}
		err := result.Execute()

		return ast.NewLiteralNumber(result.Result.Value), err
	},
	"string + string": func(left, right *ast.Literal) (*ast.Literal, error) {
		result := instruction.Concat{
			Left:  left,
			Right: right,
		}
		err := result.Execute()

		return ast.NewLiteralString(result.Result.Value), err
	},
}

func compileNode(node ast.Node) (*ast.Literal, error) {
	switch n := node.(type) {
	case *ast.Literal:
		return n, nil

	case *ast.Binary:
		leftType := typeOf(n.Left)
		rightType := typeOf(n.Right)
		op := fmt.Sprintf("%s %s %s", leftType, n.Op, rightType)

		if bop, ok := binaryOperations[op]; ok {
			left, err := compileNode(n.Left)
			if err != nil {
				return nil, err
			}

			right, err := compileNode(n.Right)
			if err != nil {
				return nil, err
			}

			return bop(left, right)
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
	}

	// TODO(elliot): Do something more sensible here.
	panic(node)
}
