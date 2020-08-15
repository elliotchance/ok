package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompileFile(t *testing.T) {
	for testName, test := range map[string]struct {
		f        *parser.File
		expected *compiler.Compiled
	}{
		"no-functions": {
			&parser.File{},
			&compiler.Compiled{
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
			&compiler.Compiled{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Instructions: []vm.Instruction{
							&vm.Print{},
						},
						Variables: map[string]string{},
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
			&compiler.Compiled{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Instructions: []vm.Instruction{
							&vm.Call{
								FunctionName: "add",
							},
						},
						Variables: map[string]string{},
						Registers: 0,
					},
					"add": {
						Variables: map[string]string{},
					},
				},
			},
		},
		"call-function-one-arg-zero-returns": {
			&parser.File{
				Funcs: map[string]*ast.Func{
					"add": {
						Arguments: []*ast.Argument{
							{Name: "x", Type: "number"},
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
			&compiler.Compiled{
				Funcs: map[string]*vm.CompiledFunc{
					"main": {
						Variables: map[string]string{},
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
						Variables: map[string]string{
							"x": "number",
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
			&compiler.Compiled{
				Funcs: map[string]*vm.CompiledFunc{
					"foo": {
						Instructions: []vm.Instruction{
							&vm.Print{},
						},
						Variables: map[string]string{},
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
							Variables: map[string]string{},
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
						Name:    "Person",
						Returns: []string{"Person"},
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
			&compiler.Compiled{
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
						Variables: map[string]string{
							"p": "Person",
						},
						Registers: 1,
					},
					"Person": {
						Variables: map[string]string{},
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
			compiledFile, err := compiler.CompileFile(test.f, nil)
			require.NoError(t, err)

			// It is not worth testing these because the Statements make it very
			// verbose.
			compiledFile.FuncDefs = nil

			assert.Equal(t, test.expected, compiledFile)
		})
	}
}
