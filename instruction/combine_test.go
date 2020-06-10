package instruction_test

import (
	"ok/ast"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right []byte
		expected    string
	}{
		"both-empty":     {[]byte(""), []byte(""), ""},
		"both-non-empty": {[]byte("foo"), []byte("bar"), "foobar"},
	} {
		t.Run(testName, func(t *testing.T) {
			left := ast.NewLiteralData(test.left)
			right := ast.NewLiteralData(test.right)
			ins := &instruction.Combine{Left: left, Right: right}
			assert.NoError(t, ins.Execute())
			assert.Equal(t, test.expected, ins.Result.Value)
		})
	}
}
