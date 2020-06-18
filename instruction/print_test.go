package instruction_test

import (
	"bytes"
	"fmt"
	"ok/ast"
	"ok/instruction"
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
			[]*ast.Literal{{Kind: "bool", Value: "true"}},
			"true\n",
		},
		"false": {
			[]*ast.Literal{{Kind: "bool", Value: "false"}},
			"false\n",
		},
		"char": {
			[]*ast.Literal{{Kind: "char", Value: "#"}},
			"#\n",
		},
		"data": {
			[]*ast.Literal{{Kind: "data", Value: "abc"}},
			"abc\n",
		},
		"number": {
			[]*ast.Literal{{Kind: "number", Value: "1.23"}},
			"1.23\n",
		},
		"string": {
			[]*ast.Literal{{Kind: "string", Value: "foo bar"}},
			"foo bar\n",
		},
		"multiple-args": {
			[]*ast.Literal{
				{Kind: "string", Value: "foo"},
				{Kind: "number", Value: "123"},
			},
			"foo 123\n",
		},
		"number-array": {
			[]*ast.Literal{
				{
					Kind: "[]number",
					Array: []*ast.Literal{
						ast.NewLiteralNumber("123"),
						ast.NewLiteralNumber("456"),
						ast.NewLiteralNumber("789"),
					},
				},
			},
			"[123, 456, 789]\n",
		},
		"any-array": {
			[]*ast.Literal{
				{
					Kind: "[]any",
					Array: []*ast.Literal{
						ast.NewLiteralBool(true),
						ast.NewLiteralChar('a'),
						ast.NewLiteralData([]byte("data")),
						ast.NewLiteralNumber("123"),
						ast.NewLiteralString("789"),
					},
				},
			},
			"[true, \"a\", \"data\", 123, \"789\"]\n",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[string]*ast.Literal{}
			var arguments []string
			for i, value := range test.values {
				register := fmt.Sprintf("%d", i)
				registers[register] = value
				arguments = append(arguments, register)
			}

			buf := bytes.NewBuffer(nil)
			ins := &instruction.Print{
				Stdout:    buf,
				Arguments: arguments,
			}
			assert.NoError(t, ins.Execute(registers, nil))
			assert.Equal(t, test.expectedStdout, buf.String())
		})
	}
}
