package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestArrayAlloc_String(t *testing.T) {
	ins := &vm.ArrayAlloc{Size: "0", Result: "1", Kind: types.StringArray}
	assert.Equal(t, "$1 = []string with $0 elements", ins.String())
}
