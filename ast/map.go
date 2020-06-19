package ast

// Map is zero or more elements.
type Map struct {
	Kind     string
	Elements []*KeyValue
}
