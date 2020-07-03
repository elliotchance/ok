package ast

// Break represents a "break" statement.
type Break struct {
	Pos string
}

// Position returns the position.
func (node *Break) Position() string {
	return node.Pos
}
