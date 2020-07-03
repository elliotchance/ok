package ast

// Key returns the value based on the index.
type Key struct {
	Expr Node
	Key  Node
	Pos  string
}

// Position returns the position.
func (node *Key) Position() string {
	return node.Pos
}
