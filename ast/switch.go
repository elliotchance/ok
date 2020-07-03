package ast

// Case represents a switch case statement.
type Case struct {
	// Conditions will always contain at least one element.
	Conditions []Node

	// Statements may be nil.
	Statements []Node

	Pos string
}

// Position returns the position.
func (node *Case) Position() string {
	return node.Pos
}

// Switch represents a switch statement.
type Switch struct {
	// Expr may be nil.
	Expr Node

	// Cases may be nil.
	Cases []*Case

	// Else may be nil.
	Else []Node

	Pos string
}

// Position returns the position.
func (node *Switch) Position() string {
	return node.Pos
}
