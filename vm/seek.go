package vm

import (
	"fmt"
	"strconv"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Seek moves the cursor of a file handle.
type Seek struct {
	Fd, Offset, Whence Register // in
	NewOffset          Register // out
}

// Execute implements the Instruction interface for the VM.
func (ins *Seek) Execute(_ *int, vm *VM) error {
	f := vm.Get(ins.Fd).File
	offset := number.Int64(number.NewNumber(vm.Get(ins.Offset).Value))
	whence := number.Int(number.NewNumber(vm.Get(ins.Whence).Value))
	pos, err := f.Seek(offset, whence)
	if err != nil {
		vm.Raise(err.Error())
		return nil
	}

	vm.Set(ins.NewOffset, asttest.NewLiteralNumber(strconv.Itoa(int(pos))))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Seek) String() string {
	return fmt.Sprintf("%s = os.Seek(%s, %s, %s)",
		ins.NewOffset, ins.Fd, ins.Offset, ins.Whence)
}
