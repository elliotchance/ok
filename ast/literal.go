package ast

// Literal represents a literal of any type.
type Literal struct {
	Kind  string
	Value string

	// Array is also used to hold the keys of the map. This is required for
	// iteration.
	Array []*Literal

	Map map[string]*Literal

	Pos string
}

// Position returns the position.
func (node *Literal) Position() string {
	return node.Pos
}
