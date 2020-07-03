package ast

// Continue represents a "continue" statement.
type Continue struct {
	Pos string
}

// Position returns the position.
func (node *Continue) Position() string {
	return node.Pos
}
