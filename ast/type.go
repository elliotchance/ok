package ast

import (
	"fmt"

	"github.com/elliotchance/ok/types"
)

func TypeOf(node Node) (*types.Type, error) {
	switch n := node.(type) {
	case *Literal:
		return n.Kind, nil

	case *Func:
		return n.Type(), nil

	case *Interpolate:
		// The result of an interpolation is always a string.
		return types.String, nil
	}

	return nil, fmt.Errorf("cannot resolve type for %v (%T)", node, node)
}
