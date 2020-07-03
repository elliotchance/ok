package ast

// Identifier could refer to a variable, function, etc.
type Identifier struct {
	Name string
	Pos  string
}

// Position returns the position.
func (node *Identifier) Position() string {
	return node.Pos
}
