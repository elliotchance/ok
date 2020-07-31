package ast

import (
	"encoding/json"

	"github.com/elliotchance/ok/compiler/kind"
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
	var x interface{} = node.Value

	switch {
	case kind.IsArray(node.Kind):
		x = node.Array

	case kind.IsMap(node.Kind):
		x = node.Map
	}

	// Ignore error, it shouldn't happen, and if it does there is nothing
	// sensible we can do in that case.
	v, _ := json.Marshal(x)

	return string(v)
}
