package ast

// Call represents a function call with zero or more arguments.
type Call struct {
	// FunctionName is the name of the function being invoked.
	FunctionName string

	// Arguments contains zero or more elements that represent each of the
	// arguments respectively.
	Arguments []Node

	Pos string
}

// Position returns the position.
func (node *Call) Position() string {
	return node.Pos
}
