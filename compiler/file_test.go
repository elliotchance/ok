package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/parser"
	"ok/vm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileFile(t *testing.T) {
	for testName, test := range map[string]struct {
		f        *parser.File
		expected map[string]*vm.CompiledFunc
	}{
		//"no-functions": {
		//	&parser.File{},
		//	map[string]*vm.CompiledFunc{},
		//},
		//"one-function": {
		//	&parser.File{
		//		Funcs: map[string]*ast.Func{
		//			"main": {
		//				Statements: []ast.Node{
		//					&ast.Call{
		//						FunctionName: "print",
		//					},
		//				},
		//			},
		//		},
		//	},
		//	map[string]*vm.CompiledFunc{
		//		"main": {
		//			Instructions: []vm.Instruction{
		//				&vm.Print{Stdout: os.Stdout},
		//			},
		//			Variables: map[string]string{},
		//		},
		//	},
		//},
		//"call-function-zero-args-zero-returns": {
		//	&parser.File{
		//		Funcs: map[string]*ast.Func{
		//			"add": {
		//				Statements: []ast.Node{},
		//			},
		//			"main": {
		//				Statements: []ast.Node{
		//					&ast.Call{
		//						FunctionName: "add",
		//					},
		//				},
		//			},
		//		},
		//	},
		//	map[string]*vm.CompiledFunc{
		//		"main": {
		//			Instructions: []vm.Instruction{
		//				&vm.Call{
		//					FunctionName: "add",
		//				},
		//			},
		//			Variables: map[string]string{},
		//			Registers: 1,
		//		},
		//		"add": {
		//			Variables: map[string]string{},
		//		},
		//	},
		//},
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
			map[string]*vm.CompiledFunc{
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
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFuncs, err := compiler.CompileFile(test.f)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, compiledFuncs)
		})
	}
}
