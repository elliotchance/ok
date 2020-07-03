package asttest

import "github.com/elliotchance/ok/ast"

// NewArrayNumbers creates an Array with some number literal values.
func NewArrayNumbers(values []string) *ast.Array {
	var elements []ast.Node
	for _, value := range values {
		elements = append(elements, NewLiteralNumber(value))
	}

	return &ast.Array{
		Kind:     "[]number",
		Elements: elements,
		// No pos needed for testing.
	}
}
