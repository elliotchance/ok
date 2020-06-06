package parser

import (
	"ok/ast"
	"ok/lexer"
)

// File contains the output state of the parser for a single input.
type File struct {
	Root     *ast.Func
	Comments []*ast.Comment
	Tokens   []lexer.Token
}
