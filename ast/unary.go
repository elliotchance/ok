package ast

// Unary is an unary operator operation.
type Unary struct {
	// Op is TokenMinus, TokenNot, TokenIncrement or TokenDecrement.
	Op   string
	Expr Node
	Pos  string
}

// Position returns the position.
func (node *Unary) Position() string {
	return node.Pos
}
