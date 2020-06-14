package ast

// If represents an if/else combination.
type If struct {
	Condition Node

	// Either or both is allowed to be nil.
	True, False []Node
}
