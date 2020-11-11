package ast

import (
	"encoding/json"

	"github.com/elliotchance/ok/types"
)

// Literal represents a literal of any type.
type Literal struct {
	Kind  *types.Type
	Value string

	// Array is also used to hold the keys of the map. This is required for
	// iteration.
	Array []*Literal

	// If the literal is a function Map will be the parent scope.
	Map map[string]*Literal

	Pos string

	// Used for file handles. Even though it is always a *os.File, we need to
	// keep it as a interface{} so that gob doesn't complain.
	File interface{}

	// We can only open one reader for a file handle as to not reset the
	// placement. This is always a *bufio.Reader, but we need to keep it an
	// interface{} for the same reason as File.
	Reader interface{}
}

// Position returns the position.
func (node *Literal) Position() string {
	return node.Pos
}

// String is the human-readable representation of the value. To make it easier
// we just output everything as JSON.
func (node *Literal) String() string {
	var x interface{} = node.Value

	switch node.Kind.Kind {
	case types.KindArray:
		x = node.Array

	case types.KindMap:
		x = node.Map

	case types.KindFunc:
		return "func " + node.Value
	}

	// Ignore error, it shouldn't happen, and if it does there is nothing
	// sensible we can do in that case.
	v, _ := json.Marshal(x)

	return string(v)
}
