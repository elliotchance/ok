package ast

// Binary is an binary operator operation.
type Binary struct {
	// Op is TokenPlus, TokenMinusAssign, etc. It will never be TokenAssign,
	// this special case is handled in an Assign operation.
	Op string

	Left, Right Node
}

// NewBinary creates a binary operation.
func NewBinary(left Node, op string, right Node) *Binary {
	return &Binary{
		Left:  left,
		Op:    op,
		Right: right,
	}
}
