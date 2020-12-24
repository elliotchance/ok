package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

// compileBlock translates list of statements into a set of instructions. The
// number of instructions returned may be zero.
func compileBlock(
	compiledFunc *vm.CompiledFunc,
	stmts []ast.Node,
	breakIns,
	continueIns vm.Instruction,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	// Before we can compile the statements we need to lookahead to extract the
	// types of the variables. Specifically, this is for functions that were
	// hoisted.
	for _, statement := range stmts {
		if a, ok := statement.(*ast.Assign); ok &&
			len(a.Lefts) == 1 && len(a.Rights) == 1 {
			if c, ok := a.Rights[0].(*ast.Func); ok && c.Name != "" {
				// TODO(elliot): Should not mutate scopeOverrides here. Copy
				//  instead.
				scopeOverrides[c.Name] = c.Type()
			}
		}
	}

	for _, statement := range stmts {
		err := compileStatement(compiledFunc, statement, breakIns, continueIns,
			file, scopeOverrides)
		if err != nil {
			return err
		}
	}

	return nil
}
