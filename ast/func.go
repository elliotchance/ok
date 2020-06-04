package ast

// Func represents the definition of a function.
type Func struct {
	// Name is the name of the function being declared.
	Name string

	// Statements can have zero or more elements for each of the ordered
	// discreet statements in the function.
	Statements []Node
}
