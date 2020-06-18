package ast

// Array is zero or more elements.
type Array struct {
	Kind     string
	Elements []Node
}
