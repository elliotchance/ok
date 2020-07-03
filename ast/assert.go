package ast

// Assert is used in tests.
type Assert struct {
	Expr *Binary
	Pos  string
}

// Position returns the position.
func (node *Assert) Position() string {
	return node.Pos
}
