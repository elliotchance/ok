package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/vm"
)

func getBinaryInstruction(op string, left, right, result vm.Register) (vm.Instruction, string) {
	switch op {
	// TODO(elliot): These below is not documented in the language spec.
	case "[]bool + []bool",
		"[]char + []char",
		"[]data + []data",
		"[]number + []number",
		"[]string + []string":
		return &vm.Append{A: left, B: right, Result: result},
			strings.Split(op, " ")[0]

	case "data + data":
		return &vm.Combine{Left: left, Right: right, Result: result}, "data"

	case "data += data":
		return &vm.Combine{Left: left, Right: right, Result: left}, "data"

	case "number + number":
		return &vm.Add{Left: left, Right: right, Result: result}, "number"

	case "number += number":
		return &vm.Add{Left: left, Right: right, Result: left}, "number"

	case "number - number":
		return &vm.Subtract{Left: left, Right: right, Result: result}, "number"

	case "number -= number":
		return &vm.Subtract{Left: left, Right: right, Result: left}, "number"

	case "number * number":
		return &vm.Multiply{Left: left, Right: right, Result: result}, "number"

	case "number *= number":
		return &vm.Multiply{Left: left, Right: right, Result: left}, "number"

	case "number / number":
		return &vm.Divide{Left: left, Right: right, Result: result}, "number"

	case "number /= number":
		return &vm.Divide{Left: left, Right: right, Result: left}, "number"

	case "number % number":
		return &vm.Remainder{Left: left, Right: right, Result: result}, "number"

	case "number %= number":
		return &vm.Remainder{Left: left, Right: right, Result: left}, "number"

	case "string + string":
		return &vm.Concat{Left: left, Right: right, Result: result}, "string"

	case "string += string":
		return &vm.Concat{Left: left, Right: right, Result: left}, "string"

	case "bool and bool":
		return &vm.And{Left: left, Right: right, Result: result}, "bool"

	case "bool or bool":
		return &vm.Or{Left: left, Right: right, Result: result}, "bool"

	case "bool == bool", "char == char", "data == data", "string == string",
		// TODO(elliot): These below is not documented in the language spec.
		"[]bool == []bool",
		"[]char == []char",
		"[]data == []data",
		"[]number == []number",
		"[]string == []string":
		return &vm.Equal{Left: left, Right: right, Result: result}, "bool"

	case "number == number":
		return &vm.EqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "bool != bool", "char != char", "data != data", "string != string",
		// TODO(elliot): These below is not documented in the language spec.
		"[]bool != []bool",
		"[]char != []char",
		"[]data != []data",
		"[]number != []number",
		"[]string != []string":
		return &vm.NotEqual{Left: left, Right: right, Result: result}, "bool"

	case "number != number":
		return &vm.NotEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "number > number":
		return &vm.GreaterThanNumber{Left: left, Right: right, Result: result}, "bool"

	case "string > string":
		return &vm.GreaterThanString{Left: left, Right: right, Result: result}, "bool"

	case "number < number":
		return &vm.LessThanNumber{Left: left, Right: right, Result: result}, "bool"

	case "string < string":
		return &vm.LessThanString{Left: left, Right: right, Result: result}, "bool"

	case "number >= number":
		return &vm.GreaterThanEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "string >= string":
		return &vm.GreaterThanEqualString{Left: left, Right: right, Result: result}, "bool"

	case "number <= number":
		return &vm.LessThanEqualNumber{Left: left, Right: right, Result: result}, "bool"

	case "string <= string":
		return &vm.LessThanEqualString{Left: left, Right: right, Result: result}, "bool"

	}

	return nil, ""
}

func compileBinary(compiledFunc *vm.CompiledFunc, node *ast.Binary, file *Compiled) (vm.Register, string, error) {
	// TokenAssign is not possible here because that is handled by an Assign
	// operation.
	if node.Op == lexer.TokenPlusAssign ||
		node.Op == lexer.TokenMinusAssign ||
		node.Op == lexer.TokenTimesAssign ||
		node.Op == lexer.TokenDivideAssign ||
		node.Op == lexer.TokenRemainderAssign {

		right, rightKind, err := compileExpr(compiledFunc, node.Right, file)
		if err != nil {
			return "", "", err
		}

		// TODO(elliot): Check +=, etc.
		if key, ok := node.Left.(*ast.Key); ok {
			arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc,
				key.Expr, file)
			if err != nil {
				return "", "", err
			}

			// TODO(elliot): Check this is a sane operation.
			keyResults, _, err := compileExpr(compiledFunc, key.Key, file)
			if err != nil {
				return "", "", err
			}

			if kind.IsArray(arrayOrMapKind[0]) {
				ins := &vm.ArraySet{
					Array: arrayOrMapResults[0],
					Index: keyResults[0],
					Value: right[0],
				}
				compiledFunc.Append(ins)
			} else {
				// TODO(elliot): Does not handle iterating strings.
				ins := &vm.MapSet{
					Map:   arrayOrMapResults[0],
					Key:   keyResults[0],
					Value: right[0],
				}
				compiledFunc.Append(ins)
			}

			// TODO(elliot): Return something more reasonable here.
			return "", "", nil
		}

		variable, ok := node.Left.(*ast.Identifier)
		if !ok {
			return "", "", fmt.Errorf("cannot assign to non-variable")
		}

		// Make sure we do not assign the wrong type to an existing variable.
		if v, ok := compiledFunc.Variables[variable.Name]; ok && rightKind[0] != v {
			return "", "", fmt.Errorf(
				"%s cannot assign %s to variable %s (expecting %s)",
				variable.Position(), rightKind[0], variable.Name, v)
		}

		switch node.Op {
		case lexer.TokenPlusAssign:
			switch {
			case kind.IsArray(rightKind[0]):
				compiledFunc.Append(&vm.Append{
					A:      vm.Register(variable.Name),
					B:      right[0],
					Result: vm.Register(variable.Name),
				})

			case rightKind[0] == "data":
				compiledFunc.Append(&vm.Combine{
					Left:   vm.Register(variable.Name),
					Right:  right[0],
					Result: vm.Register(variable.Name),
				})

			case rightKind[0] == "number":
				compiledFunc.Append(&vm.Add{
					Left:   vm.Register(variable.Name),
					Right:  vm.Register(right[0]),
					Result: vm.Register(variable.Name),
				})

			case rightKind[0] == "string":
				compiledFunc.Append(&vm.Concat{
					Left:   vm.Register(variable.Name),
					Right:  right[0],
					Result: vm.Register(variable.Name),
				})
			}

		case lexer.TokenMinusAssign:
			compiledFunc.Append(&vm.Subtract{
				Left:   vm.Register(variable.Name),
				Right:  right[0],
				Result: vm.Register(variable.Name),
			})

		case lexer.TokenTimesAssign:
			compiledFunc.Append(&vm.Multiply{
				Left:   vm.Register(variable.Name),
				Right:  right[0],
				Result: vm.Register(variable.Name),
			})

		case lexer.TokenDivideAssign:
			compiledFunc.Append(&vm.Divide{
				Left:   vm.Register(variable.Name),
				Right:  right[0],
				Result: vm.Register(variable.Name),
			})

		case lexer.TokenRemainderAssign:
			compiledFunc.Append(&vm.Remainder{
				Left:   vm.Register(variable.Name),
				Right:  right[0],
				Result: vm.Register(variable.Name),
			})
		}

		return vm.Register(variable.Name), rightKind[0], nil
	}

	_, _, returns, returnKind, err := compileComparison(compiledFunc, node, file)

	return returns, returnKind, err
}

func compileComparison(compiledFunc *vm.CompiledFunc, node *ast.Binary, file *Compiled) (vm.Register, vm.Register, vm.Register, string, error) {
	left, leftKind, err := compileExpr(compiledFunc, node.Left, file)
	if err != nil {
		return "", "", "", "", err
	}

	right, rightKind, err := compileExpr(compiledFunc, node.Right, file)
	if err != nil {
		return "", "", "", "", err
	}

	returns := compiledFunc.NextRegister()

	op := fmt.Sprintf("%s %s %s", leftKind[0], node.Op, rightKind[0])
	if bop, kind := getBinaryInstruction(op, left[0], right[0], returns); bop != nil {
		// TODO(elliot): It would be nice to be able to evaluate expressions
		//  involving literals here. So, 1 + 1 just becomes 2.

		compiledFunc.Append(bop)

		return left[0], right[0], returns, kind, nil
	}

	return left[0], right[0], returns, "",
		fmt.Errorf("%s cannot perform %s", node.Position(), op)
}
