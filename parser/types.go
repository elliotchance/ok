package parser

import (
	"fmt"
)

func (parser *Parser) resolveInterfaces() error {
	for _, fn := range parser.funcs {
		if fn.IsConstructor() {
			ty, err := fn.Interface()
			if err != nil {
				return fmt.Errorf("%v %s", fn.Position(), err)
			}

			parser.Interfaces[fn.Name] = ty
		}
	}

	return nil
}
