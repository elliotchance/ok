package vm_test

import (
	"ok/instruction"
	"ok/vm"
	"testing"
)

func TestVM_Run(t *testing.T) {
	for testName, test := range map[string]struct {
		instructions []instruction.Instruction
	}{
		"no-instructions": {nil},
		"one-instruction": {
			[]instruction.Instruction{
				&instruction.Print{},
			},
		},
		"two-instructions": {
			[]instruction.Instruction{
				&instruction.Print{},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			vm := vm.NewVM(test.instructions)
			vm.Run()
		})
	}
}
