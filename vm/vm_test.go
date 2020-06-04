package vm_test

import (
	"bytes"
	"ok/ast"
	"ok/instruction"
	"ok/lexer"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVM_Run(t *testing.T) {
	for testName, test := range map[string]struct {
		instructions   []instruction.Instruction
		expectedStdout string
	}{
		"no-instructions": {},
		"one-instruction": {
			instructions: []instruction.Instruction{
				&instruction.Print{
					Value: &ast.Literal{Kind: lexer.TokenNumber, Value: "1.23"},
				},
			},
			expectedStdout: "1.23\n",
		},
		"two-instructions": {
			instructions: []instruction.Instruction{
				&instruction.Print{
					Value: &ast.Literal{Kind: lexer.TokenString, Value: "foo bar"},
				},
				&instruction.Print{
					Value: &ast.Literal{Kind: lexer.TokenNumber, Value: "1.23"},
				},
			},
			expectedStdout: "foo bar\n1.23\n",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			vm := vm.NewVM(test.instructions, buf)
			vm.Run()
			assert.Equal(t, test.expectedStdout, buf.String())
		})
	}
}
