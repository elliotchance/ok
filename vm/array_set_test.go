package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestArraySet_String(t *testing.T) {
	ins := &vm.ArraySet{Array: "0", Index: "1", Value: "2"}
	assert.Equal(t, "$0[$1] = $2", ins.String())
}
