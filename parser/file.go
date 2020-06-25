package parser

import (
	"ok/ast"
	"ok/lexer"
)

// File contains the output state of the parser for a single input.
type File struct {
	Funcs    map[string]*ast.Func
	Tests    []*ast.Test
	Comments []*ast.Comment
	Tokens   []lexer.Token
}
