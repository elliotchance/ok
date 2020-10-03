package parser

import (
	"fmt"

	"github.com/elliotchance/ok/types"
)

func (parser *Parser) resolveInterfaces() error {
	for _, fn := range parser.funcs {
		// Argument types
		for _, arg := range fn.Arguments {
			if arg.Type.Kind == types.KindUnresolvedInterface {
				// We may not know all the interfaces yet, so only replace those
				// that are known. resolveInterfaces() will be called again
				// later. If anything is still unresolved, the compile will have
				// to handle that.
				if lookupFn, ok := parser.funcs[arg.Type.Name]; ok {
					ty, err := lookupFn.Interface()
					if err != nil {
						return fmt.Errorf("%v %s", fn.Position(), err)
					}

					arg.Type = types.NewInterface(arg.Type.Name, ty)
				}
			}
		}

		// Return types
		for i, ret := range fn.Returns {
			if ret.Kind == types.KindUnresolvedInterface {
				// We may not know all the interfaces yet, so only replace those
				// that are known. resolveInterfaces() will be called again
				// later. If anything is still unresolved, the compile will have
				// to handle that.
				if lookupFn, ok := parser.funcs[ret.Name]; ok {
					ty, err := lookupFn.Interface()
					if err != nil {
						return fmt.Errorf("%v %s", fn.Position(), err)
					}

					fn.Returns[i] = types.NewInterface(ret.Name, ty)
				}
			}
		}
	}

	return nil
}
