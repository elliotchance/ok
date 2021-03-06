package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileIdentifier(
	compiledFunc *vm.CompiledFunc,
	e *ast.Identifier,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) ([]vm.Register, []*types.Type, error) {
	if e.Name[0] == '^' {
		if compiledFunc.Parent == nil {
			return nil, nil, fmt.Errorf("%s does not exist in the parent scope", e.Name[1:])
		}

		parentVar, ok := compiledFunc.Parent.GetTypeForVariable(e.Name[1:], scopeOverrides)
		if !ok {
			return nil, nil, fmt.Errorf("%s does not exist in the parent scope", e.Name[1:])
		}

		// TODO(elliot): We must copy this out-of-scope value to a register
		//  so that it's found correctly in the stack. This doesn't have to
		//  be for in-scope variables because they are register names
		//  themselves. However, the tricky part is that we cannot copy to a
		//  register if the value is intended to be mutated (eg. ++^i) so
		//  there needs to be a new option passed down the compile functions
		//  to indicate this. Without this change code like "padZero(^foo)"
		//  will not work because the scope will be incorrect by the time
		//  ^foo is accessed.
		return []vm.Register{vm.Register(e.Name)},
			[]*types.Type{parentVar}, nil
	}

	// Constants (defined at the package-level) can be referenced from
	// anywhere. This only covers the case where we are referencing a
	// constant that belongs to the current package, as external constants
	// would be resolved through the package import variable.
	if c, ok := compiledFunc.Constants[e.Name]; ok {
		// We copy it locally to make sure it's value isn't changed. The
		// compiler will prevent a constant from being modified directly.
		//
		// TODO(elliot): The compiler needs to raise an error when trying to
		//  modify a constant.
		literalRegister := compiledFunc.NextRegister()

		if c.IsGlobal {
			compiledFunc.Append(&vm.Assign{
				Result:   literalRegister,
				Register: vm.Register("$" + c.Value),
			})
		} else {
			compiledFunc.Append(&vm.AssignSymbol{
				Result: literalRegister,
				Symbol: file.AddSymbolLiteral(c),
			})
		}

		return []vm.Register{literalRegister}, []*types.Type{c.Kind}, nil
	}

	if v, ok := compiledFunc.GetTypeForVariable(e.Name, scopeOverrides); ok {
		return []vm.Register{vm.Register(e.Name)}, []*types.Type{v}, nil
	}

	// It could also reference a package-level function.
	if fn := file.FuncByName(e.Name); fn != nil {
		literalRegister := compiledFunc.NextRegister()
		ty := file.Types.Get(string(fn.Type))
		compiledFunc.Append(&vm.AssignSymbol{
			Result: literalRegister,
			Symbol: file.AddSymbolLiteral(&ast.Literal{
				Kind:  ty,
				Value: fn.UniqueName,
			}),
		})

		return []vm.Register{literalRegister}, []*types.Type{ty}, nil
	}

	return nil, nil, fmt.Errorf("%s undefined variable: %s",
		e.Pos, e.Name)
}
