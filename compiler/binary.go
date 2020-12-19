package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func getBinaryInstruction(
	op string,
	left,
	right,
	result vm.Register,
) (vm.Instruction, *types.Type) {
	if strings.HasPrefix(op, "any == ") ||
		strings.HasSuffix(op, " == any") ||
		strings.HasPrefix(op, "any != ") ||
		strings.HasSuffix(op, " != any") {
		return &vm.Equal{Left: left, Right: right, Result: result}, types.Bool
	}

	switch op {
	// TODO(elliot): These below is not documented in the language spec.
	case "[]bool + []bool",
		"[]char + []char",
		"[]data + []data",
		"[]number + []number",
		"[]string + []string":
		return &vm.Append{A: left, B: right, Result: result},
			types.TypeFromString(strings.Split(op, " ")[0])

	case "data + data":
		return &vm.Combine{Left: left, Right: right, Result: result}, types.Data

	case "data += data":
		return &vm.Combine{Left: left, Right: right, Result: left}, types.Data

	case "number + number":
		return &vm.Add{Left: left, Right: right, Result: result}, types.Number

	case "number += number":
		return &vm.Add{Left: left, Right: right, Result: left}, types.Number

	case "number - number":
		return &vm.Subtract{Left: left, Right: right, Result: result}, types.Number

	case "number -= number":
		return &vm.Subtract{Left: left, Right: right, Result: left}, types.Number

	case "number * number":
		return &vm.Multiply{Left: left, Right: right, Result: result}, types.Number

	case "number *= number":
		return &vm.Multiply{Left: left, Right: right, Result: left}, types.Number

	case "number / number":
		return &vm.Divide{Left: left, Right: right, Result: result}, types.Number

	case "number /= number":
		return &vm.Divide{Left: left, Right: right, Result: left}, types.Number

	case "number % number":
		return &vm.Remainder{Left: left, Right: right, Result: result}, types.Number

	case "number %= number":
		return &vm.Remainder{Left: left, Right: right, Result: left}, types.Number

	case "string + string":
		return &vm.Concat{Left: left, Right: right, Result: result}, types.String

	case "string += string":
		return &vm.Concat{Left: left, Right: right, Result: left}, types.String

	case "bool and bool":
		return &vm.And{Left: left, Right: right, Result: result}, types.Bool

	case "bool or bool":
		return &vm.Or{Left: left, Right: right, Result: result}, types.Bool

	case "bool == bool",
		"char == char",
		"data == data",
		"string == string",

		// TODO(elliot): These below is not documented in the language spec.
		// arrays
		"[]bool == []bool",
		"[]char == []char",
		"[]data == []data",
		"[]number == []number",
		"[]string == []string",
		"[]any == []any",

		// maps
		"{}bool == {}bool",
		"{}char == {}char",
		"{}data == {}data",
		"{}number == {}number",
		"{}string == {}string",
		"{}any == {}any":
		return &vm.Equal{Left: left, Right: right, Result: result}, types.Bool

	case "number == number":
		return &vm.EqualNumber{Left: left, Right: right, Result: result}, types.Bool

	case "bool != bool",
		"char != char",
		"data != data",
		"string != string",

		// TODO(elliot): These below is not documented in the language spec.
		// arrays
		"[]bool != []bool",
		"[]char != []char",
		"[]data != []data",
		"[]number != []number",
		"[]string != []string",
		"[]any != []any",

		// maps
		"{}bool != {}bool",
		"{}char != {}char",
		"{}data != {}data",
		"{}number != {}number",
		"{}string != {}string",
		"{}any != {}any":
		return &vm.NotEqual{Left: left, Right: right, Result: result}, types.Bool

	case "number != number":
		return &vm.NotEqualNumber{Left: left, Right: right, Result: result}, types.Bool

	case "number > number":
		return &vm.GreaterThanNumber{Left: left, Right: right, Result: result}, types.Bool

	case "string > string", "char > char":
		return &vm.GreaterThanString{Left: left, Right: right, Result: result}, types.Bool

	case "number < number":
		return &vm.LessThanNumber{Left: left, Right: right, Result: result}, types.Bool

	case "string < string", "char < char":
		return &vm.LessThanString{Left: left, Right: right, Result: result}, types.Bool

	case "number >= number":
		return &vm.GreaterThanEqualNumber{Left: left, Right: right, Result: result}, types.Bool

	case "string >= string", "char >= char":
		return &vm.GreaterThanEqualString{Left: left, Right: right, Result: result}, types.Bool

	case "number <= number":
		return &vm.LessThanEqualNumber{Left: left, Right: right, Result: result}, types.Bool

	case "string <= string", "char <= char":
		return &vm.LessThanEqualString{Left: left, Right: right, Result: result}, types.Bool
	}

	return nil, nil
}

func compileBinary(
	compiledFunc *vm.CompiledFunc,
	node *ast.Binary,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, *types.Type, error) {
	// Type check.
	if node.Op == lexer.TokenIs {
		left, _, err := compileExpr(compiledFunc, node.Left, file,
			scopeOverrides)
		if err != nil {
			return "", nil, err
		}

		typeRegister := compiledFunc.NextRegister()
		tyName := asttest.NewLiteralString(node.Right.(*ast.Identifier).Name)
		compiledFunc.Append(&vm.Assign{
			Value:        tyName,
			VariableName: typeRegister,
		})

		resultRegister := compiledFunc.NextRegister()

		compiledFunc.Append(&vm.Is{
			Value:  left[0],
			Type:   typeRegister,
			Result: resultRegister,
		})

		return resultRegister, types.Bool, nil
	}

	// TokenAssign is not possible here because that is handled by an Assign
	// operation.
	if node.Op == lexer.TokenPlusAssign ||
		node.Op == lexer.TokenMinusAssign ||
		node.Op == lexer.TokenTimesAssign ||
		node.Op == lexer.TokenDivideAssign ||
		node.Op == lexer.TokenRemainderAssign {
		right, rightKind, err := compileExpr(compiledFunc, node.Right, file,
			scopeOverrides)
		if err != nil {
			return "", nil, err
		}

		// TODO(elliot): Check +=, etc.
		if key, ok := node.Left.(*ast.Key); ok {
			arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc,
				key.Expr, file, scopeOverrides)
			if err != nil {
				return "", nil, err
			}

			// TODO(elliot): Check this is a sane operation.
			keyResults, _, err := compileExpr(compiledFunc, key.Key, file,
				scopeOverrides)
			if err != nil {
				return "", nil, err
			}

			if arrayOrMapKind[0].Kind == types.KindArray {
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
			return "", nil, nil
		}

		variable, ok := node.Left.(*ast.Identifier)
		if !ok {
			return "", nil, fmt.Errorf("cannot assign to non-variable")
		}

		// Make sure we do not assign the wrong type to an existing variable.
		if v, ok := compiledFunc.GetTypeForVariable(variable.Name, scopeOverrides); ok && rightKind[0].String() != v.String() {
			return "", nil, fmt.Errorf(
				"%s cannot assign %s to variable %s (expecting %s)",
				variable.Position(), rightKind[0], variable.Name, v)
		}

		switch node.Op {
		case lexer.TokenPlusAssign:
			switch {
			case rightKind[0].Kind == types.KindArray:
				compiledFunc.Append(&vm.Append{
					A:      vm.Register(variable.Name),
					B:      right[0],
					Result: vm.Register(variable.Name),
				})

			case rightKind[0].Kind == types.KindData:
				compiledFunc.Append(&vm.Combine{
					Left:   vm.Register(variable.Name),
					Right:  right[0],
					Result: vm.Register(variable.Name),
				})

			case rightKind[0].Kind == types.KindNumber:
				compiledFunc.Append(&vm.Add{
					Left:   vm.Register(variable.Name),
					Right:  vm.Register(right[0]),
					Result: vm.Register(variable.Name),
				})

			case rightKind[0].Kind == types.KindString:
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

	_, _, returns, returnKind, err := compileComparison(compiledFunc, node,
		file, scopeOverrides)

	return returns, returnKind, err
}

func compileComparison(
	compiledFunc *vm.CompiledFunc,
	node *ast.Binary,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, vm.Register, vm.Register, *types.Type, error) {
	left, leftKind, err := compileExpr(compiledFunc, node.Left, file,
		scopeOverrides)
	if err != nil {
		return "", "", "", nil, err
	}

	right, rightKind, err := compileExpr(compiledFunc, node.Right, file,
		scopeOverrides)
	if err != nil {
		return "", "", "", nil, err
	}

	returns := compiledFunc.NextRegister()

	op := fmt.Sprintf("%s %s %s", leftKind[0], node.Op, rightKind[0])
	if bop, kind := getBinaryInstruction(op, left[0], right[0], returns); bop != nil {
		// TODO(elliot): It would be nice to be able to evaluate expressions
		//  involving literals here. So, 1 + 1 just becomes 2.

		compiledFunc.Append(bop)

		return left[0], right[0], returns, kind, nil
	}

	return left[0], right[0], returns, nil,
		fmt.Errorf("%s cannot perform %s", node.Position(), op)
}
