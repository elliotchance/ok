package ast

import "fmt"

func TypeOf(node Node) (string, error) {
	switch n := node.(type) {
	case *Literal:
		return n.Kind, nil

	case *Func:
		return n.Type(), nil

	case *Interpolate:
		// The result of an interpolation is always a string.
		return "string", nil
	}

	return "", fmt.Errorf("cannot resolve type for %v (%T)", node, node)
}
