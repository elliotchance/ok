package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

type builtinFn func(
	compiledFunc *vm.CompiledFunc,
	args []vm.Register,
) (vm.Instruction, vm.Register, *types.Type, error)

// TODO(elliot): These needs to check function signatures.
var builtinFunctions = map[string]builtinFn{
	"__call":      funcCall,
	"__get":       funcGet,
	"__interface": funcInterface,
	"__len":       funcLen,
	"__log":       funcLog,
	"__pow":       funcPow,
	"__props":     funcProps,
	"__set":       funcSet,
	"__type":      funcType,
	"char":        funcChar,
	"len":         funcLen,
	"number":      funcNumber,
	"print":       funcPrint,
	"string":      funcString,
}

func compileCall(
	compiledFunc *vm.CompiledFunc,
	call *ast.Call,
	file *vm.File,
) ([]vm.Register, []*types.Type, error) {
	var argResults []vm.Register
	for _, arg := range call.Arguments {
		argResult, _, err := compileExpr(compiledFunc, arg, file)
		if err != nil {
			return nil, nil, err
		}

		argResults = append(argResults, argResult...)
	}

	if name, ok := call.Expr.(*ast.Identifier); ok {
		if fn, ok := builtinFunctions[name.Name]; ok {
			ins, result, returnType, err := fn(compiledFunc, argResults)
			if err != nil {
				return nil, nil, err
			}

			compiledFunc.Append(ins)

			return []vm.Register{result}, []*types.Type{returnType}, nil
		}
	}

	fnResult, fnType, err := compileExpr(compiledFunc, call.Expr, file)
	if err != nil {
		return nil, nil, err
	}

	// TODO(elliot): IMPORTANT: Check argument types are valid.
	if len(fnType) != 1 || fnType[0].Kind != types.KindFunc {
		// TODO(elliot): Make this error message cleaner and more helpful.
		return nil, nil, fmt.Errorf("%s cannot call %s",
			call.Expr.Position(), fnType)
	}

	// Prepare enough return registers.
	var returnRegisters []vm.Register
	for range fnType[0].Returns {
		returnRegisters = append(returnRegisters, compiledFunc.NextRegister())
	}

	ins := &vm.Call{
		FunctionName: fmt.Sprintf("*%s", string(fnResult[0])),
		Arguments:    argResults,
		Results:      returnRegisters,
	}

	compiledFunc.Append(ins)

	return returnRegisters, fnType[0].Returns, nil
}

func funcNumber(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastNumber{
		X:      args[0],
		Result: result,
	}

	return ins, result, types.Number, nil
}

func funcString(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastString{
		X:      args[0],
		Result: result,
	}

	return ins, result, types.String, nil
}

func funcChar(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastChar{
		X:      args[0],
		Result: result,
	}

	return ins, result, types.Char, nil
}

func funcInterface(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Interface{
		Value:  args[0],
		Result: result,
	}

	return ins, result, types.String, nil
}

func funcGet(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Get{
		Object: args[0],
		Prop:   args[1],
		Result: result,
	}

	return ins, result, types.String, nil
}

func funcSet(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Set{
		Object: args[0],
		Prop:   args[1],
		Value:  args[2],
		Result: result,
	}

	return ins, result, types.Bool, nil
}

func funcProps(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Props{
		Object: args[0],
		Result: result,
	}

	return ins, result, types.StringArray, nil
}

func funcType(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Type{
		Value:  args[0],
		Result: result,
	}

	return ins, result, types.String, nil
}

func funcCall(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()

	// TODO(elliot): This is not a safe operation because it assumes the
	//  provided inputs are sane since they cannot be checked by the compiler.
	//  We will need some runtime check in the instruction itself.
	ins := &vm.DynamicCall{
		Variable:  args[0],
		Arguments: args[1],
		Results:   result,
	}

	return ins, result, types.AnyArray, nil
}

func funcLog(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Log{
		X:      args[0],
		Result: result,
	}

	return ins, result, types.Number, nil
}

func funcPow(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Power{
		Base:   args[0],
		Power:  args[1],
		Result: result,
	}

	return ins, result, types.Number, nil
}

func funcLen(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Len{
		Argument: args[0],
		Result:   result,
	}

	return ins, result, types.Number, nil
}

func funcPrint(_ *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, *types.Type, error) {
	ins := &vm.Print{
		Arguments: args,
	}

	return ins, "", nil, nil
}
