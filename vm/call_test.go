package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestCall_String(t *testing.T) {
	ins := &vm.Call{
		FunctionName: "foo",
		Arguments:    []vm.Register{"1", "2"},
		Results:      []vm.Register{"4", "5"},
	}
	assert.Equal(t, "($4, $5) = foo($1, $2)", ins.String())
}
