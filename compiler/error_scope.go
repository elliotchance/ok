package compiler

import (
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileErrorScope(
	compiledFunc *vm.CompiledFunc,
	n *ast.ErrorScope,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	// Only activate the finally clause if there is one.
	if n.Finally != nil {
		compiledFunc.Append(&vm.Finally{
			Index: n.Finally.Index,
			Run:   true,
		})
	}

	// Try section.
	err := compileBlock(compiledFunc, n.Statements, nil, nil,
		file, scopeOverrides)
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
		// TODO(elliot): Check that the type is actually an interface.
		// TODO(elliot): Handle missing/invalid types more gracefully.

		parts := strings.Split(on.Type, ".")
		var typeRegister string
		if len(parts) == 2 {
			interfaceType := compiledFunc.Constants[parts[0]].Kind.Properties[parts[1]].Returns[0]
			typeRegister, err = file.Types.Add(interfaceType)
		} else {
			typeRegister, err = file.Types.Add(compiledFunc.Constants[on.Type].Kind.Returns[0])
		}

		if err != nil {
			panic(err)
		}

		compiledFunc.Append(&vm.On{
			Type: vm.TypeRegister(typeRegister),
		})

		// Provide the err variable. The runtime value will be provided by the
		// On instruction above.
		compiledFunc.NewVariable("err", file.Types.Get(typeRegister))

		err = compileBlock(compiledFunc, on.Statements, nil,
			nil, file, scopeOverrides)
		if err != nil {
			return err
		}

		compiledFunc.Append(done)
	}

	// An On with an empty Type signals to a VM trying to recover from an error
	// that there is no handler. It will pass the error up to the caller.
	anyRegister, err := file.Types.Add(types.Any)
	if err != nil {
		return err
	}
	compiledFunc.Append(&vm.On{
		Finished: true,

		// This isn't actually used when Finished - true, but we must provide a
		// valid type for all TypeRegisters.
		Type: vm.TypeRegister(anyRegister),
	})

	// Correct the jump after the error has been handled. The "-1" is to
	// correct for the "+1" that would happen after every instruction.
	done.To = len(compiledFunc.Instructions.Instructions) - 1

	// The optional finally clause exists as unrelated code after all the
	// clauses. On success the finally will run here, otherwise the block was
	// activated as soon as the try was entered and the VM will ensure it is run
	// before any return.
	if n.Finally != nil {
		beforeLen := len(compiledFunc.Instructions.Instructions)

		compiledFunc.Append(&vm.Finally{
			Index: n.Finally.Index,
			Run:   false,
		})

		// Finally section.
		err := compileBlock(compiledFunc, n.Finally.Statements, nil,
			nil, file, scopeOverrides)
		if err != nil {
			return err
		}

		finallyInstructions := compiledFunc.Instructions.Instructions[beforeLen:]
		compiledFunc.Finally = append(compiledFunc.Finally,
			vm.NewInstructions(finallyInstructions...))
	}

	return nil
}
