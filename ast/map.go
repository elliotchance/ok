package ast

// Map is zero or more elements.
type Map struct {
	Kind     string
	Elements []*KeyValue
	Pos      string
}

// Position returns the position.
func (node *Map) Position() string {
	return node.Pos
}
