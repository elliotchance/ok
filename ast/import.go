package ast

// Import is used to include packages.
type Import struct {
	PackageName string
	Pos         string
}

// Position returns the position.
func (node *Import) Position() string {
	return node.Pos
}
