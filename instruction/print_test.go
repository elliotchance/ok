package instruction_test

import (
	"bytes"
	"ok/ast"
	"ok/instruction"
	"ok/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrint_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		values         []*ast.Literal
		expectedStdout string
	}{
		"no-args": {nil, "\n"},
		"true": {
			[]*ast.Literal{{Kind: lexer.TokenBool, Value: "true"}},
			"true\n",
		},
		"false": {
			[]*ast.Literal{{Kind: lexer.TokenBool, Value: "false"}},
			"false\n",
		},
		"char": {
			[]*ast.Literal{{Kind: lexer.TokenCharacter, Value: "#"}},
			"#\n",
		},
		"data": {
			[]*ast.Literal{{Kind: lexer.TokenData, Value: "abc"}},
			"abc\n",
		},
		"number": {
			[]*ast.Literal{{Kind: lexer.TokenNumber, Value: "1.23"}},
			"1.23\n",
		},
		"string": {
			[]*ast.Literal{{Kind: lexer.TokenString, Value: "foo bar"}},
			"foo bar\n",
		},
		"multiple-args": {
			[]*ast.Literal{
				{Kind: lexer.TokenString, Value: "foo"},
				{Kind: lexer.TokenNumber, Value: "123"},
			},
			"foo 123\n",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			ins := &instruction.Print{
				Stdout: buf,
				Values: test.values,
			}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expectedStdout, buf.String())
		})
	}
}
