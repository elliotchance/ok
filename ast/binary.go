package ast

// Binary is an binary operator operation.
type Binary struct {
	// Op is TokenPlus, TokenMinusAssign, etc.
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
