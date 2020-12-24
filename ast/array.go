package ast

import (
	"github.com/elliotchance/ok/types"
)

// Array is zero or more elements.
type Array struct {
	Kind     *types.Type
	Elements []Node
	Pos      string
}

// Position returns the position.
func (node *Array) Position() string {
	return node.Pos
}
