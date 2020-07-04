package ast

// Interpolate is a string literal that contains expressions.
type Interpolate struct {
	// Parts will have at least one element. Each element will be one of
	// StringLiteral or Group. However, they can appear in any order.
	Parts []Node

	Pos string
}

// Position returns the position.
func (node *Interpolate) Position() string {
	return node.Pos
}
