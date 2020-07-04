package ast

// Key returns the value based on the index. It is used for arrays, maps and
// objects.
type Key struct {
	Expr Node
	Key  Node
	Pos  string
}

// Position returns the position.
func (node *Key) Position() string {
	return node.Pos
}
