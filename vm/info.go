package vm

import (
	"fmt"
	"os"

	"github.com/elliotchance/ok/ast/asttest"
)

// Info fetches file information or raises and error if the file does not exist.
type Info struct {
	Path    Register   // In
	Name    Register   // Out
	Size    Register   // Out
	Mode    Register   // Out
	ModTime []Register // Out (6 elements)
	IsDir   Register   // Out
}

// Execute implements the Instruction interface for the VM.
func (ins *Info) Execute(_ *int, vm *VM) error {
	info, err := os.Stat(vm.Get(ins.Path).Value)
	if err != nil {
		vm.Raise(err.Error())
		return nil
	}

	vm.Set(ins.Name, asttest.NewLiteralString(info.Name()))
	vm.Set(ins.Size, asttest.NewLiteralNumber(fmt.Sprintf("%d", info.Size())))
	vm.Set(ins.Mode, asttest.NewLiteralString(info.Mode().String()))
	setTimeLiterals(vm, info.ModTime(),
		ins.ModTime[0], ins.ModTime[1], ins.ModTime[2],
		ins.ModTime[3], ins.ModTime[4], ins.ModTime[5])
	vm.Set(ins.IsDir, asttest.NewLiteralBool(info.IsDir()))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Info) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s = os.Info(%s)",
		ins.Name, ins.Size, ins.Mode,
		ins.ModTime[0], ins.ModTime[1], ins.ModTime[2],
		ins.ModTime[3], ins.ModTime[4], ins.ModTime[5],
		ins.IsDir, ins.Path,
	)
}
