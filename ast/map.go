package ast

import "github.com/elliotchance/ok/types"

// Map is zero or more elements.
type Map struct {
	Kind     *types.Type
	Elements []*KeyValue
	Pos      string
}

// Position returns the position.
func (node *Map) Position() string {
	return node.Pos
}
