package parser

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

func (parser *Parser) resolveInterfaces() error {
	var err error

	for _, fn := range parser.funcs {
		// Argument types
		for _, arg := range fn.Arguments {
			arg.Type, err = parser.resolveInterface(fn, arg.Type)
			if err != nil {
				return err
			}
		}

		// Return types
		for i, ret := range fn.Returns {
			fn.Returns[i], err = parser.resolveInterface(fn, ret)
			if err != nil {
				return err
			}
		}

		err = parser.resolveInterfacesInStatements(fn.Statements)
		if err != nil {
			return err
		}
	}

	for _, fn := range parser.tests {
		err = parser.resolveInterfacesInStatements(fn.Statements)
		if err != nil {
			return err
		}
	}

	return nil
}

func (parser *Parser) resolveInterface(node ast.Node, typ *types.Type) (*types.Type, error) {
	// We may not know all the interfaces yet, so only replace those that are
	// known. resolveInterfaces() will be called again later. If anything is
	// still unresolved, the compile will have to handle that.
	if typ.Kind == types.KindUnresolvedInterface {
		// TODO(elliot): This is a hack to get around the fact we have no linker
		//  yet. Remove in the future.
		if typ.Name == "error.Error" {
			return types.ErrorInterface, nil
		}

		if lookupFn, ok := parser.funcs[typ.Name]; ok {
			ty, err := lookupFn.Interface()
			if err != nil {
				return typ, fmt.Errorf("%v %s", node.Position(), err)
			}

			return types.NewInterface(typ.Name, ty), nil
		}
	}

	return typ, nil
}

func (parser *Parser) resolveInterfacesInStatements(stmts []ast.Node) error {
	var err error

	for _, stmt := range stmts {
		if es, ok := stmt.(*ast.ErrorScope); ok {
			for _, on := range es.On {
				on.Type, err = parser.resolveInterface(on, on.Type)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
