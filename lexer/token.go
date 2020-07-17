package lexer

import "fmt"

// Tokens defined here have a human-readable name used in error messages. You
// should not rely on these values staying the same, only that their value will
// be unique amongst the defined tokens.
const (
	TokenEOF = "end of file"

	// Dynamic
	TokenBoolLiteral   = "bool literal"   // boolean literal, eg. true
	TokenCharLiteral   = "char literal"   // char literal, eg. 'a'
	TokenComment       = "comment"        // eg. "//..."
	TokenDataLiteral   = "data literal"   // data literal, eg. `foo`
	TokenIdentifier    = "identifier"     // any non-keyword
	TokenNumberLiteral = "number literal" // number literal, eg. 12.3
	TokenStringLiteral = "string literal" // string literal, eg. "hello"

	// Keywords
	TokenAnd      = "and"
	TokenAny      = "any"
	TokenAssert   = "assert"
	TokenBool     = "bool"
	TokenBreak    = "break"
	TokenCase     = "case"
	TokenChar     = "char"
	TokenContinue = "continue"
	TokenData     = "data"
	TokenElse     = "else"
	TokenFinally  = "finally"
	TokenFor      = "for"
	TokenFunc     = "func"
	TokenIf       = "if"
	TokenImport   = "import"
	TokenIn       = "in"
	TokenNot      = "not"
	TokenNumber   = "number"
	TokenOn       = "on"
	TokenOr       = "or"
	TokenRaise    = "raise"
	TokenReturn   = "return"
	TokenString   = "string"
	TokenSwitch   = "switch"
	TokenTest     = "test"
	TokenTry      = "try"

	// Operators
	TokenAssign           = "="
	TokenColon            = ":"
	TokenComma            = ","
	TokenCurlyClose       = "}"
	TokenCurlyOpen        = "{"
	TokenDecrement        = "--"
	TokenDivide           = "/"
	TokenDivideAssign     = "/="
	TokenDot              = "."
	TokenEqual            = "=="
	TokenGreaterThan      = ">"
	TokenGreaterThanEqual = ">="
	TokenIncrement        = "++"
	TokenLessThan         = "<"
	TokenLessThanEqual    = "<="
	TokenMinus            = "-"
	TokenMinusAssign      = "-="
	TokenNotEqual         = "!="
	TokenParenClose       = ")"
	TokenParenOpen        = "("
	TokenPlus             = "+"
	TokenPlusAssign       = "+="
	TokenRemainder        = "%"
	TokenRemainderAssign  = "%="
	TokenSemiColon        = ";"
	TokenSquareClose      = "]"
	TokenSquareOpen       = "["
	TokenTimes            = "*"
	TokenTimesAssign      = "*="

	// Interpolation tokens works like brackets (using TokenComma to separate
	// each part) around literals and expressions that becomes one interpolation
	// expression.
	TokenInterpolateStart = "interpolate start"
	TokenInterpolateEnd   = "interpolate end"
)

// Token represents a single token.
type Token struct {
	// Kind will be one of the Token* constants.
	Kind string

	// Value is captured from the original source code. It is only useful for
	// dynamic tokens such as TokenString or TokenIdentifier.
	Value string

	// IsEndOfLine will be true if there is at least one new line character
	// after this token (ignoring other whitespace). This is needed by some
	// grammars to determine the end of the line, but newlines have no effect in
	// between most tokens.
	//
	// One exception to this is comments. Comment always have to be terminated
	// by a new line, however, IsEndOfLine will only be true if the comment is
	// followed by an empty line. This is because IsEndOfLine is also used by
	// the lexer to determine if comments are part of the same block as their
	// previous lines and/or if they might be attached to functions are
	// documentation.
	IsEndOfLine bool

	// Pos is the location of the token.
	Pos Pos
}

// NewToken initializes a token with a kind and value and other
// defaults.
func NewToken(kind, value string, pos Pos) Token {
	return Token{
		Kind:        kind,
		Value:       value,
		IsEndOfLine: false,
		Pos:         pos,
	}
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
