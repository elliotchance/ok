package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestCastString_String(t *testing.T) {
	ins := &vm.CastString{X: "1", Result: "2"}
	assert.Equal(t, "$2 = string $1", ins.String())
}

func TestCastNumber_String(t *testing.T) {
	ins := &vm.CastNumber{X: "1", Result: "2"}
	assert.Equal(t, "$2 = number $1", ins.String())
}

func TestCastChar_String(t *testing.T) {
	ins := &vm.CastChar{X: "1", Result: "2"}
	assert.Equal(t, "$2 = char $1", ins.String())
}
