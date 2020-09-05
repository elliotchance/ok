package asttest

import (
	"strconv"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func AssertEqual(t *testing.T, fns1, fns2 interface{}) bool {
	// TODO(elliot): This just ignores some transient data that will be removed
	//  in the future. Remove this code to see what breaks (and fix it).
	if m, ok := fns2.(map[string]*ast.Func); ok {
		delete(m, "")
		for k := range m {
			if _, err := strconv.Atoi(k); err == nil {
				delete(m, k)
			}
		}
	}

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
		cmpopts.IgnoreFields(ast.Func{}, "UniqueName"),
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

		// TODO(elliot): The function name must be ignored for now because it's
		//  based off a random value. See anonFunctionName. This should be
		//  removed once that hack is fixed.
		cmpopts.IgnoreFields(ast.Func{}, "Name"),
	}
}
