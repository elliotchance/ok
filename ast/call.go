package ast

// Call represents a function call with zero or more arguments.
type Call struct {
	// FunctionName is the name of the function being invoked.
	FunctionName string

	// Arguments contains zero or more elements that represent each of the
	// arguments respectively.
	Arguments []*Literal
}
