package ast

// Case represents a switch case statement.
type Case struct {
	// Conditions will always contain at least one element.
	Conditions []Node

	// Statements may be nil.
	Statements []Node
}

// Switch represents a switch statement.
type Switch struct {
	// Expr may be nil.
	Expr Node

	// Cases may be nil.
	Cases []*Case

	// Else may be nil.
	Else []Node
}
