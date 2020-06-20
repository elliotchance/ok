package ast

import "sort"

// Map is zero or more elements.
type Map struct {
	Kind     string
	Elements []*KeyValue
}

// NewMapNumbers creates a Map with some number literal values. This function is
// for convenience and should not be used for general production code.
func NewMapNumbers(values map[string]string) *Map {
	// For predictability we load the keys in order.
	var keys []string
	for key := range values {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var elements []*KeyValue
	for _, key := range keys {
		value := values[key]
		elements = append(elements, &KeyValue{
			Key:   NewLiteralString(key),
			Value: NewLiteralNumber(value),
		})
	}

	return &Map{
		Kind:     "{}number",
		Elements: elements,
	}
}
