package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileErrorScope(compiledFunc *vm.CompiledFunc, n *ast.ErrorScope, fns map[string]*ast.Func) error {
	// Try section.
	err := compileBlock(compiledFunc, n.Statements, nil, nil, fns)
	if err != nil {
		return err
	}

	// The done jump will be correct later. It is called after all the try
	// statements (ie. there was no error raised) and will jump to after all the
	// error handlers so the problem can continue.
	//
	// It is also placed at the end of each error handler so that the program
	// can continue.
	done := &vm.Jump{}

	compiledFunc.Append(done)

	// Each of the On clauses.
	for _, on := range n.On {
		compiledFunc.Append(&vm.On{
			Type: on.Type,
		})

		err := compileBlock(compiledFunc, on.Statements, nil, nil, fns)
		if err != nil {
			return err
		}

		compiledFunc.Append(done)
	}

	// An On with an empty Type signals to a VM trying to recover from an error
	// that there is no handler. It will pass the error up to the caller.
	compiledFunc.Append(&vm.On{Type: ""})

	// Correct the jump after the error has been handled. The "-1" is to
	// correct for the "+1" that would happen after every instruction.
	done.To = len(compiledFunc.Instructions) - 1

	return nil
}
