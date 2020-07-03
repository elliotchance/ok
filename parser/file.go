package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// File contains the output state of the parser for a single input.
type File struct {
	Funcs    map[string]*ast.Func
	Tests    []*ast.Test
	Imports  map[string]string
	Comments []*ast.Comment
	Tokens   []lexer.Token
}

// Pos returns the rendered position of a token.
func (file *File) Pos(offset int) string {
	return file.Tokens[offset].Pos.String()
}
