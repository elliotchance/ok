package ast

// KeyValue represents a key-value pair used in maps and object initialization.
type KeyValue struct {
	Key, Value Node
}

// Position returns the position.
func (node *KeyValue) Position() string {
	return node.Key.Position()
}
