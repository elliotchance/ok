package parser

import (
	"fmt"
	"reflect"

	"github.com/elliotchance/ok/ast"
	"github.com/pkg/errors"
)

type Parser struct {
	errors        []error
	File          *File
	finalizers    map[string][]*ast.Finally
	functionNames []string
}

// AppendError adds an error to the stack.
func (p *Parser) AppendErrorAt(pos string, message string) {
	var err error
	if pos != "" {
		err = fmt.Errorf("%s %s", pos, message)
	} else {
		err = errors.New(message)
	}
	p.errors = append(p.errors, err)
}

// AppendError adds an error to the stack.
func (p *Parser) AppendError(node ast.Node, message string) {
	if v := reflect.ValueOf(node); v.IsValid() && !v.IsNil() {
		p.AppendErrorAt(node.Position(), message)
	} else {
		p.AppendErrorAt("", message)
	}
}

// AppendErrorf adds an error to the stack.
func (p *Parser) AppendErrorf(node ast.Node, format string, args ...interface{}) {
	p.AppendError(node, fmt.Sprintf(format, args...))
}

// Errors returns all errors.
func (p *Parser) Errors() []error {
	return p.errors
}

// AppendFinally will track finalizers until the function is finished, then it
// will be reset.
func (p *Parser) AppendFinally(finally *ast.Finally) {
	funcName := p.functionNames[len(p.functionNames)-1]
	finally.Index = len(p.finalizers[funcName])
	p.finalizers[funcName] = append(p.finalizers[funcName], finally)
}
