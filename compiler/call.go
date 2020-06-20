package compiler

import (
	"ok/ast"
	"ok/vm"
	"os"
)

func compileCall(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, error) {
	builtinFunctions := map[string]func(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, error){
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
			return nil, err
		}

		argResults = append(argResults, argResult...)
	}

	// TODO(elliot): Check function exists.
	toCall := fns[call.FunctionName]

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

	return returnRegisters, nil
}

func compileCallLen(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, error) {
	argumentResults, _, err := compileExpr(compiledFunc, call.Arguments[0], fns)
	if err != nil {
		return nil, err
	}

	result := compiledFunc.NextRegister()
	ins := &vm.Len{
		Argument: argumentResults[0],
		Result:   result,
	}

	compiledFunc.Append(ins)

	return []string{result}, nil
}

func compileCallPrint(compiledFunc *vm.CompiledFunc, call *ast.Call, fns map[string]*ast.Func) ([]string, error) {
	ins := &vm.Print{
		Stdout: os.Stdout,
	}
	for _, arg := range call.Arguments {
		returns, _, err := compileExpr(compiledFunc, arg, fns)
		if err != nil {
			return nil, err
		}

		ins.Arguments = append(ins.Arguments, returns...)
	}

	compiledFunc.Append(ins)

	return nil, nil
}
