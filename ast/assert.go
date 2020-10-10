package ast

// Assert is for the `assert(BinaryExpr)` syntax.
type Assert struct {
	Expr *Binary
	Pos  string
}

// Position returns the position.
func (node *Assert) Position() string {
	return node.Pos
}

// AssertRaise is for the `assert(Call raise TypeOrValue)` syntax.
type AssertRaise struct {
	Call        *Call
	TypeOrValue Node
	Pos         string
}

// Position returns the position.
func (node *AssertRaise) Position() string {
	return node.Pos
}
