package ast

// Call represents a function call with zero or more arguments.
type Call struct {
	// Expr is the expression that returns the function to be called.
	Expr Node

	// Arguments contains zero or more elements that represent each of the
	// arguments respectively.
	Arguments []Node

	Pos string
}

// Position returns the position.
func (node *Call) Position() string {
	return node.Pos
}
