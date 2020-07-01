package ast

import "strings"

// Comment represents a single or multiline comment. All characters immediately
// following `//` are part of the comment (even the proceeding space) up to but
// not including the new line.
type Comment struct {
	Comment string

	// Func will be the name of the function this comment is attached to;
	// otherwise it will be empty.
	Func string
}

// String returns the cleaner presentation version of the comment.
func (c *Comment) String() string {
	lines := strings.Split(c.Comment, "\n")
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, strings.TrimSpace(line))
	}

	return strings.Join(newLines, "\n")
}
