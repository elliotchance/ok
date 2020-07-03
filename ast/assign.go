package ast

// Assign is a specific case of Binary, just for "=".
type Assign struct {
	// There will always be at least one Lefts element.
	Lefts []Node

	// There may be one Right element, or the same number of elements as Lefts.
	Rights []Node
}

// Position returns the position.
func (node *Assign) Position() string {
	return node.Lefts[0].Position()
}
