package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
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
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralData(test.left),
				"1": asttest.NewLiteralData(test.right),
			}
			ins := &vm.Combine{Left: "0", Right: "1", Result: "2"}
			vm := &vm.VM{
				Stack: []map[vm.Register]*ast.Literal{registers},
			}
			assert.NoError(t, ins.Execute(nil, vm))
			assert.Equal(t, test.expected, registers[ins.Result].Value)
		})
	}
}

func TestCombine_String(t *testing.T) {
	ins := &vm.Combine{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = combine($1, $2)", ins.String())
}
