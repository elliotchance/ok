package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

type builtinFn func(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error)

// TODO(elliot): These needs to check function signatures.
var builtinFunctions = map[string]builtinFn{
	"__log":  funcLog,
	"__pow":  funcPow,
	"len":    funcLen,
	"print":  funcPrint,
	"string": funcString,
	"number": funcNumber,
	"char":   funcChar,
}

func compileCall(compiledFunc *vm.CompiledFunc, call *ast.Call, file *Compiled) ([]vm.Register, []string, error) {
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

		return []vm.Register{result}, []string{returnType}, nil
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

func findFunc(compiledFunc *vm.CompiledFunc, call *ast.Call, file *Compiled) (*ast.Func, error) {
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

	// It might be a built in function.
	//
	// TODO(elliot): This needs to only allow this usage if its imported.
	if internal := vm.Lib[call.FunctionName]; internal != nil {
		return internal.FuncDef, nil
	}

	// Is it a method being called on a variable?
	parts := strings.Split(call.FunctionName, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s no such function: %s",
			call.Position(), call.FunctionName)
	}

	if _, ok := compiledFunc.Variables[parts[0]]; !ok {
		return nil, fmt.Errorf("%s no such function: %s",
			call.Position(), call.FunctionName)
	}

	if ty, ok := file.FuncDefs[parts[1]]; ok {
		toCall = ast.NewFuncFromPrototype(ty.String())
		call.FunctionName = parts[1]

		return toCall, nil
	}

	return nil, fmt.Errorf("%s no such function: %s",
		call.Position(), call.FunctionName)
}

func funcNumber(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastNumber{
		X:      args[0],
		Result: result,
	}

	return ins, result, "number", nil
}

func funcString(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastString{
		X:      args[0],
		Result: result,
	}

	return ins, result, "string", nil
}

func funcChar(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.CastChar{
		X:      args[0],
		Result: result,
	}

	return ins, result, "char", nil
}

func funcLog(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Log{
		X:      args[0],
		Result: result,
	}

	return ins, result, "number", nil
}

func funcPow(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Power{
		Base:   args[0],
		Power:  args[1],
		Result: result,
	}

	return ins, result, "number", nil
}

func funcLen(compiledFunc *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	result := compiledFunc.NextRegister()
	ins := &vm.Len{
		Argument: args[0],
		Result:   result,
	}

	return ins, result, "number", nil
}

func funcPrint(_ *vm.CompiledFunc, args []vm.Register) (vm.Instruction, vm.Register, string, error) {
	ins := &vm.Print{
		Arguments: args,
	}

	return ins, "", "", nil
}
