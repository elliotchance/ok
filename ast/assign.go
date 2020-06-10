package ast

// Assign is used to set a variable to the result of an expression.
type Assign struct {
	VariableName string
	Expr         Node
}
