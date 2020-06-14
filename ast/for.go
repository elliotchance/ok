package ast

// For represents a for loop.
type For struct {
	// Condition may be nil.
	Condition Node

	Statements []Node
}
