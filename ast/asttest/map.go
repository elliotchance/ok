package asttest

import (
	"sort"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// NewMapNumbers creates a Map with some number literal values. This function is
// for convenience and should not be used for general production code.
func NewMapNumbers(values map[string]string) *ast.Map {
	// For predictability we load the keys in order.
	var keys []string
	for key := range values {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var elements []*ast.KeyValue
	for _, key := range keys {
		value := values[key]
		elements = append(elements, &ast.KeyValue{
			Key:   NewLiteralString(key),
			Value: NewLiteralNumber(value),
		})
	}

	return &ast.Map{
		Kind:     types.NumberMap,
		Elements: elements,
	}
}
