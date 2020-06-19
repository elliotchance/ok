package ast

// Array is zero or more elements.
type Array struct {
	Kind     string
	Elements []Node
}

// NewArrayNumbers creates an Array with some number literal values.
func NewArrayNumbers(values []string) *Array {
	var elements []Node
	for _, value := range values {
		elements = append(elements, NewLiteralNumber(value))
	}

	return &Array{
		Kind:     "[]number",
		Elements: elements,
	}
}
