package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
)

func TestRegister_String(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.Equal(t, "()", vm.Registers{}.String())
	})

	t.Run("one", func(t *testing.T) {
		assert.Equal(t, "($1)", vm.Registers{"1"}.String())
	})

	t.Run("two", func(t *testing.T) {
		assert.Equal(t, "($123, a7)", vm.Registers{"123", "a7"}.String())
	})
}

func TestRegisters_String(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, "_", vm.Register("").String())
	})

	t.Run("register", func(t *testing.T) {
		assert.Equal(t, "$0", vm.Register("0").String())
		assert.Equal(t, "$123", vm.Register("123").String())
	})

	// TODO(elliot): Using variable names as registers will be removed in the
	//  future.
	t.Run("variable-name", func(t *testing.T) {
		assert.Equal(t, "foo", vm.Register("foo").String())
		assert.Equal(t, "r1", vm.Register("r1").String())
	})
}
