package parser

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
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

func (parser *Parser) Package(packageAlias string) *ast.Func {
	properties := map[string]*types.Type{}

	var constantNames []string
	for name := range parser.Constants {
		constantNames = append(constantNames, name)
	}
	sort.Strings(constantNames)

	var statements []ast.Node
	for _, name := range constantNames {
		c := parser.Constants[name]
		properties[name] = c.Kind
		statements = append(statements, &ast.Assign{
			Lefts:  []ast.Node{&ast.Identifier{Name: name}},
			Rights: []ast.Node{c},
		})
	}

	var funcNames []string
	for name := range parser.funcs {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)

	for _, name := range funcNames {
		fn := parser.funcs[name]
		properties[name] = fn.Type()
		statements = append(statements, &ast.Assign{
			Lefts:  []ast.Node{&ast.Identifier{Name: name}},
			Rights: []ast.Node{fn},
		})
	}

	return &ast.Func{
		Name: packageAlias,
		Returns: []*types.Type{
			types.NewInterface(packageAlias, properties),
		},
		Statements: statements,
		Pos:        parser.pos(0),
	}
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
func NewParser(anonFunctionName int) *Parser {
	return &Parser{
		finalizers:       map[string][]*ast.Finally{},
		Constants:        map[string]*ast.Literal{},
		imports:          map[string]string{},
		funcs:            map[string]*ast.Func{},
		anonFunctionName: anonFunctionName,
	}
}

// pos returns the rendered position of a token.
func (parser *Parser) pos(offset int) string {
	return parser.tokens[offset].Pos.String()
}
