package lexer

// Options allows configuration of the lexer.
type Options struct {
	// IncludeComments will include TokenComment in the returned tokens.
	IncludeComments bool
}
