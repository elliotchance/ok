package lexer

import "fmt"

// Tokens defined here have a human-readable name used in error messages. You
// should not rely on these values staying the same, only that their value will
// be unique amongst the defined tokens.
const (
	TokenEOF = "end of file"

	// Dynamic
	TokenIdentifier = "identifier" // any non-keyword
	TokenString     = "string"     // "..."
	TokenComment    = "comment"    // "//..."

	// Keywords
	TokenFunc = "func"

	// Operators
	TokenParenOpen  = "("
	TokenParenClose = ")"
	TokenCurlyOpen  = "{"
	TokenCurlyClose = "}"
)

// Token represents a single token.
type Token struct {
	// Kind will be one of the Token* constants.
	Kind string

	// Value is captured from the original source code. It is only useful for
	// dynamic tokens such as TokenString or TokenIdentifier.
	Value string
}

// String returns a human-readable representation of the token.
func (t Token) String() string {
	if t.Kind == TokenEOF {
		return t.Kind
	}

	if t.Kind == TokenIdentifier {
		return fmt.Sprintf(`"%s"`, t.Value)
	}

	return fmt.Sprintf("'%s'", t.Kind)
}
