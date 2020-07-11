package ast

// Raise will raise an error to be handled. An error must be in the form of a
// constructor call, this is for simplicity right now.
type Raise struct {
	Err *Call
	Pos string
}

// Position returns the position.
func (node *Raise) Position() string {
	return node.Pos
}
