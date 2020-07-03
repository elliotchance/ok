package ast

// Return is a statement to return values.
type Return struct {
	// Exprs can be zero or more elements.
	Exprs []Node
	Pos   string
}

// Position returns the position.
func (node *Return) Position() string {
	return node.Pos
}
