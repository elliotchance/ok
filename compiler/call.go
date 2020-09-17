package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
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

	if fn, ok := builtinFunctions[call.FunctionName]; ok {
		ins, result, returnType, err := fn(compiledFunc, argResults)
		if err != nil {
			return nil, nil, err
		}

		compiledFunc.Append(ins)

		return []vm.Register{result}, []*types.Type{returnType}, nil
	}

	toCall, err := findFunc(compiledFunc, call, file)
	if err != nil {
		return nil, nil, err
	}

	// Prepare enough return registers.
	var returnRegisters []vm.Register
	for range toCall.Returns {
		returnRegisters = append(returnRegisters, compiledFunc.NextRegister())
	}

	ins := &vm.Call{
		FunctionName: call.FunctionName,
		Arguments:    argResults,
		Results:      returnRegisters,
	}

	compiledFunc.Append(ins)

	return returnRegisters, toCall.Returns, nil
}

func findFunc(
	compiledFunc *vm.CompiledFunc,
	call *ast.Call,
	file *vm.File,
) (*ast.Func, error) {
	// Before we look for the function by name, we should first ensure that it's
	// not a variable being called as a function.
	//
	// TODO(elliot): Don't let a variable shadow a function of the same name.
	//
	// TODO(elliot): Check variable is indeed a function and the arguments and
	//  return values are legal.
	if ty, ok := compiledFunc.Variables[call.FunctionName]; ok {
		toCall := ast.NewFuncFromPrototype(ty)

		// TODO(elliot): This is a pretty nasty hack for now. This will tell the
		//  VM at runtime to resolve this variable to the real function name.
		call.FunctionName = "*" + call.FunctionName

		return toCall, nil
	}

	toCall := file.FuncDefs[call.FunctionName]
	if toCall != nil {
		return toCall, nil
	}

	toCall = file.ImportedFuncs[call.FunctionName]
	if toCall != nil {
		return toCall, nil
	}

	// Is it a method being called on a variable?
	parts := strings.Split(call.FunctionName, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s no such function: %s",
			call.Position(), call.FunctionName)
	}

	ty, ok := compiledFunc.Variables[parts[0]]
	if !ok {
		return nil, fmt.Errorf("%s no such function %s on variable %s",
			call.Position(), call.FunctionName, parts[0])
	}

	methodType, ok := file.Interfaces[ty.Name][parts[1]]
	if !ok {
		return nil, fmt.Errorf("%s no such function %s on %s",
			call.Position(), parts[1], ty)
	}

	keyRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: keyRegister,
		Value:        asttest.NewLiteralString(parts[1]),
	})

	callRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.MapGet{
		Map:    vm.Register(parts[0]),
		Key:    keyRegister,
		Result: callRegister,
	})

	toCall = ast.NewFuncFromPrototype(methodType)
	call.FunctionName = "*" + string(callRegister)

	return toCall, nil
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
