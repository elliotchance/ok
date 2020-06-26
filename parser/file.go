package parser

import (
	"ok/ast"
	"ok/lexer"
)

// File contains the output state of the parser for a single input.
type File struct {
	Funcs    map[string]*ast.Func
	Tests    []*ast.Test
	Imports  map[string]string
	Comments []*ast.Comment
	Tokens   []lexer.Token
}
