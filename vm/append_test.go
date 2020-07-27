package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestAppend_String(t *testing.T) {
	ins := &vm.Append{A: "0", B: "1", Result: "2"}
	assert.Equal(t, "$2 = append($0, $1)", ins.String())
}
