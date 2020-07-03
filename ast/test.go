package ast

// Test is a named test.
type Test struct {
	Name       string
	Statements []Node
	Pos        string
}

// Position returns the position.
func (node *Test) Position() string {
	return node.Pos
}
