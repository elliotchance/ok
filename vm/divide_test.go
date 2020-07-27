package vm_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestDivide_Execute(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right string
		expected    string
		err         error
	}{
		"success":     {"1.2200", "4.7", "0.25957446808510638298", nil},
		"divide-zero": {"1.2200", "0", "0", errors.New("division by zero")},
	} {
		t.Run(testName, func(t *testing.T) {
			registers := map[vm.Register]*ast.Literal{
				"0": asttest.NewLiteralNumber(test.left),
				"1": asttest.NewLiteralNumber(test.right),
			}
			ins := &vm.Divide{Left: "0", Right: "1", Result: "2"}
			assert.Equal(t, test.err, ins.Execute(registers, nil, nil))
			actual := number.NewNumber(registers[ins.Result].Value)
			assert.Equal(t, test.expected, number.Format(actual, -1))
		})
	}
}

func TestDivide_String(t *testing.T) {
	ins := &vm.Divide{Left: "1", Right: "2", Result: "3"}
	assert.Equal(t, "$3 = $1 / $2", ins.String())
}
