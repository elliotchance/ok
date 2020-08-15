package parser

import "fmt"

func (p *Parser) resolveInterfaces() error {
	p.Interfaces = make(map[string]map[string]string)

	for _, fn := range p.File.Funcs {
		if fn.IsConstructor() {
			ty, err := fn.Interface()
			if err != nil {
				return fmt.Errorf("%v %s", fn.Position(), err)
			}

			p.Interfaces[fn.Name] = ty
		}
	}

	return nil
}
