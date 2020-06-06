package ast

// Comment represents a single line comment. All characters immediately
// following `//` are part of the comment (even the proceeding space) upto but
// not including the new line.
type Comment struct {
	Comment string
}
