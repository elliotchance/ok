package ast

// Binary is an binary operator operation.
type Binary struct {
	// Op is TokenPlus, TokenMinusAssign, etc. It will never be TokenAssign,
	// this special case is handled in an Assign operation.
	Op string

	Left, Right Node
}

// Position returns the position.
func (node *Binary) Position() string {
	return node.Left.Position()
}
