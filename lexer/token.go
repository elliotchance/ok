package lexer

import "fmt"

// Tokens defined here have a human-readable name used in error messages. You
// should not rely on these values staying the same, only that their value will
// be unique amongst the defined tokens.
const (
	TokenEOF = "end of file"

	// Dynamic
	TokenBool       = "bool"       // bool literal, eg. true
	TokenCharacter  = "character"  // char literal, eg. 'a'
	TokenComment    = "comment"    // eg. "//..."
	TokenData       = "data"       // data literal, eg. `foo`
	TokenIdentifier = "identifier" // any non-keyword
	TokenNumber     = "number"     // number literal, eg. 12.3
	TokenString     = "string"     // string literal, eg. "hello"

	// Keywords
	TokenFunc = "func"

	// Operators
	TokenCurlyClose = "}"
	TokenCurlyOpen  = "{"
	TokenParenClose = ")"
	TokenParenOpen  = "("
	TokenPlus       = "+"
	TokenMinus      = "-"
	TokenTimes      = "*"
	TokenDivide     = "/"
	TokenRemainder  = "%"
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
