package ast

// For represents a for loop.
type For struct {
	// All of Init, Condition and Next may be nil.
	Init, Condition, Next Node

	// Statements may be nil.
	Statements []Node

	Pos string
}

// Position returns the position.
func (node *For) Position() string {
	return node.Pos
}

// In represents an "in" expression in for loops.
type In struct {
	Value, Key string
	Expr       Node
	Pos        string
}

// Position returns the position.
func (node *In) Position() string {
	return node.Pos
}
