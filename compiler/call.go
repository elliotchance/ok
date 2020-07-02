package compiler

import (
	"fmt"
	"os"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileCall(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error) {
	builtinFunctions := map[string]func(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error){
		"__log": compileCallLog,
		"__pow": compileCallPow,
		"len":   compileCallLen,
		"print": compileCallPrint,
	}

	if fn, ok := builtinFunctions[call.FunctionName]; ok {
		return fn(compiledFunc, call, fns)
	}

	var argResults []string
	for _, arg := range call.Arguments {
		argResult, _, err := compileExpr(compiledFunc, arg, fns)
		if err != nil {
			return nil, nil, err
		}

		argResults = append(argResults, argResult...)
	}

	toCall := fns[call.FunctionName]
	if toCall == nil {
		// It might be a built in function.
		//
		// TODO(elliot): This needs to only allow this usage if its imported.
		if internal := vm.Lib[call.FunctionName]; internal != nil {
			toCall = internal.FuncDef
		}

		if toCall == nil {
			return nil, nil,
				fmt.Errorf("no such function: %s", call.FunctionName)
		}
	}

	// Prepare enough return registers.
	var returnRegisters []string
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

func compileCallLog(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error) {
	x, _, err := compileExpr(compiledFunc, call.Arguments[0], fns)
	if err != nil {
		return nil, nil, err
	}

	result := compiledFunc.NextRegister()
	ins := &vm.Log{
		X:      x[0],
		Result: result,
	}

	compiledFunc.Append(ins)

	return []string{result}, []string{"number"}, nil
}

func compileCallPow(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error) {
	base, _, err := compileExpr(compiledFunc, call.Arguments[0], fns)
	if err != nil {
		return nil, nil, err
	}

	power, _, err := compileExpr(compiledFunc, call.Arguments[1], fns)
	if err != nil {
		return nil, nil, err
	}

	result := compiledFunc.NextRegister()
	ins := &vm.Power{
		Base:   base[0],
		Power:  power[0],
		Result: result,
	}

	compiledFunc.Append(ins)

	return []string{result}, []string{"number"}, nil
}

func compileCallLen(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error) {
	argumentResults, _, err := compileExpr(compiledFunc, call.Arguments[0], fns)
	if err != nil {
		return nil, nil, err
	}

	result := compiledFunc.NextRegister()
	ins := &vm.Len{
		Argument: argumentResults[0],
		Result:   result,
	}

	compiledFunc.Append(ins)

	return []string{result}, []string{"number"}, nil
}

func compileCallPrint(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, []string, error) {
	ins := &vm.Print{
		Stdout: os.Stdout,
	}
	for _, arg := range call.Arguments {
		returns, _, err := compileExpr(compiledFunc, arg, fns)
		if err != nil {
			return nil, nil, err
		}

		ins.Arguments = append(ins.Arguments, returns...)
	}

	compiledFunc.Append(ins)

	return nil, nil, nil
}
