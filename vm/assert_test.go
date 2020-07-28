package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestAssert_String(t *testing.T) {
	ins := &vm.Assert{Left: "0", Right: "1", Final: "2", Op: "==", Pos: "pos"}
	assert.Equal(t, "assert($0 == $1)", ins.String())
}
