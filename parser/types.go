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
				ty, err := parser.funcs[arg.Type.Name].Interface()
				if err != nil {
					return fmt.Errorf("%v %s", fn.Position(), err)
				}

				arg.Type = types.NewInterface(arg.Type.Name, ty)
			}
		}

		// Return types
		for i, ret := range fn.Returns {
			if ret.Kind == types.KindUnresolvedInterface {
				ty, err := parser.funcs[ret.Name].Interface()
				if err != nil {
					return fmt.Errorf("%v %s", fn.Position(), err)
				}

				fn.Returns[i] = types.NewInterface(ret.Name, ty)
			}
		}
	}

	return nil
}
