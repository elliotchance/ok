package ast

// Group is an expression wrapped in "()".
type Group struct {
	Expr Node
	Pos  string
}

// Position returns the position.
func (node *Group) Position() string {
	return node.Pos
}
