package parser

import (
	"fmt"
	"reflect"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/pkg/errors"
)

type Parser struct {
	errors           []error
	funcs            map[string]*ast.Func
	tests            []*ast.Test
	finalizers       map[string][]*ast.Finally
	functionNames    []string
	anonFunctionName int
	imports          map[string]string
	comments         []*ast.Comment

	// Constants are variables defined at the package level. They cannot be
	// modified and only allow literals for values.
	//
	// TODO(elliot): We should allow for expressions that can be resolved at
	//  compile time, such as "3600 * 24".
	Constants map[string]*ast.Literal

	// tokens are reset with each Parse* function call.
	tokens []lexer.Token
}

// appendError adds an error to the stack.
func (parser *Parser) appendErrorAt(pos string, message string) {
	var err error
	if pos != "" {
		err = fmt.Errorf("%s %s", pos, message)
	} else {
		err = errors.New(message)
	}
	parser.errors = append(parser.errors, err)
}

// appendError adds an error to the stack.
func (parser *Parser) appendError(node ast.Node, message string) {
	if v := reflect.ValueOf(node); v.IsValid() && !v.IsNil() {
		parser.appendErrorAt(node.Position(), message)
	} else {
		parser.appendErrorAt("", message)
	}
}

// appendErrorf adds an error to the stack.
func (parser *Parser) appendErrorf(node ast.Node, format string, args ...interface{}) {
	parser.appendError(node, fmt.Sprintf(format, args...))
}

// Errors returns all errors.
func (parser *Parser) Errors() Errors {
	return parser.errors
}

// appendFinally will track finalizers until the function is finished, then it
// will be reset.
func (parser *Parser) appendFinally(finally *ast.Finally) {
	funcName := parser.functionNames[len(parser.functionNames)-1]
	finally.Index = len(parser.finalizers[funcName])
	parser.finalizers[funcName] = append(parser.finalizers[funcName], finally)
}

// nextFunctionName returns a unique name to be used internally for anonymous
// functions.
func (parser *Parser) nextFunctionName() string {
	parser.anonFunctionName++

	return fmt.Sprintf("%d", parser.anonFunctionName)
}

func (parser *Parser) Funcs() map[string]*ast.Func {
	return parser.funcs
}

func (parser *Parser) Tests() []*ast.Test {
	return parser.tests
}

func (parser *Parser) Imports() map[string]string {
	return parser.imports
}

// Comments returns all comments collected from parsing all inputs.
func (parser *Parser) Comments() []*ast.Comment {
	return parser.comments
}

// NewParser creates an empty parser.
func NewParser() *Parser {
	return &Parser{
		finalizers: map[string][]*ast.Finally{},
		Constants:  map[string]*ast.Literal{},
		imports:    map[string]string{},
		funcs:      map[string]*ast.Func{},
	}
}

// pos returns the rendered position of a token.
func (parser *Parser) pos(offset int) string {
	return parser.tokens[offset].Pos.String()
}
