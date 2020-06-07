package ast

// Unary is an unary operator operation.
type Unary struct {
	// Op is TokenMinus.
	Op string

	Expr *Literal
}
