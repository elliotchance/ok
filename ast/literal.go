package ast

import (
	"encoding/json"
)

// Literal represents a literal of any type.
type Literal struct {
	Kind  string
	Value string

	// Array is also used to hold the keys of the map. This is required for
	// iteration.
	Array []*Literal

	Map map[string]*Literal

	Pos string
}

// Position returns the position.
func (node *Literal) Position() string {
	return node.Pos
}

// String is the human-readable representation of the value. To make it easier
// we just output everything as JSON.
func (node *Literal) String() string {
	// Ignore error, it shouldn't happen, and if it does there is nothing
	// sensible we can do in that case.
	v, _ := json.Marshal(node.Value)

	// TODO(elliot): This doesn't handle non-scalar types. However, they can
	//  only exist at runtime so for what this function is used for it may not
	//  be an issue.

	return string(v)
}
