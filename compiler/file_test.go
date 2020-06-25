package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/lexer"
	"ok/parser"
	"ok/vm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
							&vm.Print{Stdout: os.Stdout},
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
									ast.NewLiteralNumber("123"),
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
								Value:        ast.NewLiteralNumber("123"),
							},
							&vm.Call{
								FunctionName: "add",
								Arguments:    []string{"1"},
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
								Stdout:    os.Stdout,
								Arguments: []string{"x"},
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
								Expr: ast.NewBinary(
									ast.NewLiteralNumber("1"),
									lexer.TokenLessThan,
									ast.NewLiteralNumber("2"),
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
							&vm.Print{Stdout: os.Stdout},
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
									Value:        ast.NewLiteralNumber("1"),
								},
								&vm.Assign{
									VariableName: "2",
									Value:        ast.NewLiteralNumber("2"),
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
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFile, err := compiler.CompileFile(test.f)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, compiledFile)
		})
	}
}
