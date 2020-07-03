package asttest

import "github.com/elliotchance/ok/ast"

// NewCall produces a new call to functionName with any number of arguments.
func NewCall(functionName string, arguments ...ast.Node) *ast.Call {
	return &ast.Call{
		FunctionName: functionName,
		Arguments:    arguments,
	}
}
