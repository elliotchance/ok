package asttest

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func AssertEqual(t *testing.T, fns1, fns2 interface{}) bool {
	return assert.Empty(t, cmp.Diff(fns1, fns2, cmpOptions()))
}

func cmpOptions() cmp.Options {
	return []cmp.Option{
		cmpopts.IgnoreFields(ast.Array{}, "Pos"),
		cmpopts.IgnoreFields(ast.Assert{}, "Pos"),
		cmpopts.IgnoreFields(ast.Break{}, "Pos"),
		cmpopts.IgnoreFields(ast.Call{}, "Pos"),
		cmpopts.IgnoreFields(ast.Case{}, "Pos"),
		cmpopts.IgnoreFields(ast.Comment{}, "Pos"),
		cmpopts.IgnoreFields(ast.Continue{}, "Pos"),
		cmpopts.IgnoreFields(ast.ErrorScope{}, "Pos"),
		cmpopts.IgnoreFields(ast.Finally{}, "Pos"),
		cmpopts.IgnoreFields(ast.For{}, "Pos"),
		cmpopts.IgnoreFields(ast.Func{}, "Pos"),
		cmpopts.IgnoreFields(ast.Group{}, "Pos"),
		cmpopts.IgnoreFields(ast.Identifier{}, "Pos"),
		cmpopts.IgnoreFields(ast.If{}, "Pos"),
		cmpopts.IgnoreFields(ast.In{}, "Pos"),
		cmpopts.IgnoreFields(ast.Interpolate{}, "Pos"),
		cmpopts.IgnoreFields(ast.Key{}, "Pos"),
		cmpopts.IgnoreFields(ast.Literal{}, "Pos"),
		cmpopts.IgnoreFields(ast.Map{}, "Pos"),
		cmpopts.IgnoreFields(ast.On{}, "Pos"),
		cmpopts.IgnoreFields(ast.Raise{}, "Pos"),
		cmpopts.IgnoreFields(ast.Return{}, "Pos"),
		cmpopts.IgnoreFields(ast.Switch{}, "Pos"),
		cmpopts.IgnoreFields(ast.Test{}, "Pos"),
		cmpopts.IgnoreFields(ast.Unary{}, "Pos"),
	}
}
