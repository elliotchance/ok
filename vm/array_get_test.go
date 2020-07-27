package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestArrayGet_String(t *testing.T) {
	ins := &vm.ArrayGet{Array: "0", Index: "1", Result: "2"}
	assert.Equal(t, "$2 = $0[$1]", ins.String())
}
