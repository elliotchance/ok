package ast

import "fmt"

// Call represents a function call with zero or more arguments.
type Call struct {
	// Expr is the expression that returns the function to be called.
	Expr Node

	// Arguments contains zero or more elements that represent each of the
	// arguments respectively.
	Arguments []Node

	Pos string
}

// Position returns the position.
func (node *Call) Position() string {
	return node.Pos
}

// FunctionName is a temporary migration step. It will return what uses to be
// the FunctionName string.
//
// TODO(elliot): This can be removed once the compiler properly understands how
//  to deal with the Expr in all cases.
func (node *Call) FunctionName() string {
	return stringify(node.Expr)
}

func stringify(node Node) string {
	switch e := node.(type) {
	case *Identifier:
		return e.Name

	case *Key:
		return fmt.Sprintf("%s.%s", stringify(e.Expr), stringify(e.Key))
	}

	panic(fmt.Sprintf("%+#v", node))
}
