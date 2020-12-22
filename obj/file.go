package obj

import "github.com/elliotchance/ok/types"

// File represents an object file.
type File struct {
	Symbols Symbols
	Types   map[string]*types.Type
}

// NewFile creates a new empty object file.
func NewFile() *File {
	return &File{
		Symbols: Symbols{},
		Types:   Types{},
	}
}
