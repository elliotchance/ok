package ast

// Unary is an unary operator operation.
type Unary struct {
	// Op is TokenMinus or TokenNot.
	Op string

	Expr Node
}
