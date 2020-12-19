package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileInterpolate(
	compiledFunc *vm.CompiledFunc,
	n *ast.Interpolate,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, error) {
	ins := &vm.Interpolate{
		Result: compiledFunc.NextRegister(),
	}

	for _, part := range n.Parts {
		partRegisters, _, err := compileExpr(compiledFunc, part, file, scopeOverrides)
		if err != nil {
			return ins.Result, err
		}

		ins.Args = append(ins.Args, partRegisters[0])
	}

	compiledFunc.Append(ins)

	return ins.Result, nil
}
