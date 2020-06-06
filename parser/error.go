package parser

import (
	"fmt"
)

type tokenMismatch struct {
	expected, after, found string
}

func newTokenMismatch(expected, after, found string) *tokenMismatch {
	return &tokenMismatch{
		expected: expected,
		after:    after,
		found:    found,
	}
}

// Error implements the error interface.
func (err *tokenMismatch) Error() string {
	return fmt.Sprintf(
		"expecting %s after %s, but found %s",
		err.expected, err.after, err.found)
}
