package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompileFile(t *testing.T) {
	for testName, test := range map[string]struct {
		f        *parser.File
		expected *vm.File
	}{
		"no-functions": {
			&parser.File{},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{},
			},
		},
		"one-function": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"main": {
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "print",
							},
						},
					},
				},
			},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Instructions: []vm.Instruction{
							&vm.Print{},
						},
						Variables: map[string]*types.Type{},
					},
				},
			},
		},
		"call-function-zero-args-zero-returns": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"add": {
						Statements: []ast.Node{},
					},
					"main": {
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "add",
							},
						},
					},
				},
			},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Instructions: []vm.Instruction{
							&vm.Call{
								FunctionName: "add",
							},
						},
						Variables: map[string]*types.Type{},
						Registers: 0,
					},
					"add": {
						Variables: map[string]*types.Type{},
					},
				},
			},
		},
		"call-function-one-arg-zero-returns": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"add": {
						Arguments: []*ast.Argument{
							{Name: "x", Type: types.Number},
						},
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "print",
								Arguments: []ast.Node{
									&ast.Identifier{Name: "x"},
								},
							},
						},
					},
					"main": {
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "add",
								Arguments: []ast.Node{
									asttest.NewLiteralNumber("123"),
								},
							},
						},
					},
				},
			},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Variables: map[string]*types.Type{},
						Registers: 1,
						Instructions: []vm.Instruction{
							&vm.Assign{
								VariableName: "1",
								Value:        asttest.NewLiteralNumber("123"),
							},
							&vm.Call{
								FunctionName: "add",
								Arguments:    []vm.Register{"1"},
							},
						},
					},
					"add": {
						Arguments: []string{"x"},
						Variables: map[string]*types.Type{
							"x": types.Number,
						},
						Registers: 1,
						Instructions: []vm.Instruction{
							&vm.Print{
								Arguments: []vm.Register{"x"},
							},
						},
					},
				},
			},
		},
		"one-function-and-one-test": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"foo": {
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "print",
							},
						},
					},
				},
				Tests: []*ast.Test{
					{
						Name: "test foo",
						Statements: []ast.Node{
							&ast.Assert{
								Expr: asttest.NewBinary(
									asttest.NewLiteralNumber("1"),
									lexer.TokenLessThan,
									asttest.NewLiteralNumber("2"),
								),
							},
						},
					},
				},
			},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{
					"foo": {
						Instructions: []vm.Instruction{
							&vm.Print{},
						},
						Variables: map[string]*types.Type{},
					},
				},
				Tests: []*vm.CompiledTest{
					{
						CompiledFunc: &vm.CompiledFunc{
							Instructions: []vm.Instruction{
								&vm.Assign{
									VariableName: "1",
									Value:        asttest.NewLiteralNumber("1"),
								},
								&vm.Assign{
									VariableName: "2",
									Value:        asttest.NewLiteralNumber("2"),
								},
								&vm.LessThanNumber{
									Left:   "1",
									Right:  "2",
									Result: "3",
								},
								&vm.Assert{
									Left:  "1",
									Op:    "<",
									Right: "2",
									Final: "3",
								},
							},
							Variables: map[string]*types.Type{},
							Registers: 3,
						},
						TestName: "test foo",
					},
				},
			},
		},
		"call-constructor-zero-args": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"Person": {
						Name: "Person",
						Returns: []*types.Type{
							types.NewUnresolvedInterface("Person"),
						},
					},
					"main": {
						Statements: []ast.Node{
							&ast.Assign{
								Lefts: []ast.Node{
									&ast.Identifier{Name: "p"},
								},
								Rights: []ast.Node{
									&ast.Call{
										FunctionName: "Person",
									},
								},
							},
						},
					},
				},
			},
			&vm.File{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Instructions: []vm.Instruction{
							&vm.Call{
								FunctionName: "Person",
								Results:      []vm.Register{"1"},
							},
							&vm.Assign{
								VariableName: "p",
								Register:     "1",
							},
						},
						Variables: map[string]*types.Type{
							"p": types.NewUnresolvedInterface("Person"),
						},
						Registers: 1,
					},
					"Person": {
						Variables: map[string]*types.Type{},
						Instructions: []vm.Instruction{
							// return instance
							&vm.Return{
								Results: []vm.Register{vm.StateRegister},
							},
						},
					},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFile, err := compiler.CompileFile(test.f, nil, nil)
			require.NoError(t, err)

			// It is not worth testing these because the Statements make it very
			// verbose.
			compiledFile.FuncDefs = nil

			assert.Equal(t, test.expected, compiledFile)
		})
	}
}
