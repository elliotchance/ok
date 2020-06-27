package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"

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
			registers := map[string]*ast.Literal{
				"0": ast.NewLiteralData(test.left),
				"1": ast.NewLiteralData(test.right),
			}
			ins := &vm.Combine{Left: "0", Right: "1", Result: "2"}
			assert.NoError(t, ins.Execute(registers, nil, nil))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}
