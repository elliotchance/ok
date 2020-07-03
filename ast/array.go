package ast

// Array is zero or more elements.
type Array struct {
	Kind     string
	Elements []Node
	Pos      string
}

// Position returns the position.
func (node *Array) Position() string {
	return node.Pos
}
