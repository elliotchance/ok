package ast

// Case represents a switch case statement.
type Case struct {
	Condition Node

	// Statements may be nil.
	Statements []Node
}

// Switch represents a switch statement.
type Switch struct {
	// Cases may be nil.
	Cases []*Case

	// Else may be nil.
	Else []Node
}
