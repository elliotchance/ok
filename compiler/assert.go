package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileAssert(
	compiledFunc *vm.CompiledFunc,
	n *ast.Assert,
	file *vm.File,
) error {
	left, right, returns, returnKind, err := compileComparison(compiledFunc, n.Expr, file)
	if err != nil {
		return err
	}

	if returnKind.Kind != types.KindBool {
		return fmt.Errorf("assert condition must be a bool but is %s", returnKind)
	}

	compiledFunc.Append(&vm.Assert{
		Left:  left,
		Op:    n.Expr.Op,
		Right: right,
		Final: returns,
		Pos:   n.Position(),
	})

	return nil
}

func compileAssertRaise(
	compiledFunc *vm.CompiledFunc,
	n *ast.AssertRaise,
	file *vm.File,
) error {
	// assert(<call> raise <type>) is just syntactic sugar for:
	//
	//    raised = false
	//    try {
	//        raiseAnError()
	//    } on error.Error {
	//        raised = true
	//    }
	//    assert(raised == true)
	//

	raisedVariable := "__raised" // compiledFunc.NextRegister()

	err := compileAssign(compiledFunc, &ast.Assign{
		Lefts:  []ast.Node{&ast.Identifier{Name: raisedVariable}},
		Rights: []ast.Node{asttest.NewLiteralBool(false)},
	}, file)
	if err != nil {
		return err
	}

	err = compileErrorScope(compiledFunc, &ast.ErrorScope{
		Statements: []ast.Node{n.Call},
		On: []*ast.On{
			{
				Type: types.ErrorInterface, // TODO(elliot): Fix me.
				Statements: []ast.Node{
					&ast.Assign{
						Lefts:  []ast.Node{&ast.Identifier{Name: raisedVariable}},
						Rights: []ast.Node{asttest.NewLiteralBool(true)},
					},
				},
				Pos: n.Position(),
			},
		},
		Pos: n.Position(),
	}, file)
	if err != nil {
		return err
	}

	err = compileAssert(compiledFunc, &ast.Assert{
		Expr: &ast.Binary{
			Op:    "==",
			Left:  &ast.Identifier{Name: raisedVariable},
			Right: asttest.NewLiteralBool(true),
		},
		Pos: n.Position(),
	}, file)
	if err != nil {
		return err
	}

	return nil
}
