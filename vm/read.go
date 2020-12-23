package vm

import (
	"bufio"
	"fmt"
	"io"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// ReadData reads upto a fixed amount of bytes from a file.
type ReadData struct {
	Fd, Size Register // In
	Data     Register // Out
}

// Execute implements the Instruction interface for the VM.
func (ins *ReadData) Execute(_ *int, vm *VM) error {
	f := vm.Get(ins.Fd)
	initReader(f)

	size := number.Int64(number.NewNumber(vm.Get(ins.Size).Value))
	buf := make([]byte, size)
	n, err := f.Reader.Read(buf)
	if err != nil && err != io.EOF {
		vm.Raise(err.Error())
		return nil
	}

	vm.Set(ins.Data, asttest.NewLiteralData(buf[:n]))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ReadData) String() string {
	return fmt.Sprintf("%s = %s.ReadData(%s)", ins.Data, ins.Fd, ins.Size)
}

// ReadString reads upto a fixed amount of characters from a file.
type ReadString struct {
	Fd, Size Register // In
	Str      Register // Out
}

func initReader(fd *ast.Literal) {
	if fd.Reader == nil {
		fd.Reader = bufio.NewReader(fd.File)
	}
}

// Execute implements the Instruction interface for the VM.
func (ins *ReadString) Execute(_ *int, vm *VM) error {
	f := vm.Get(ins.Fd)
	initReader(f)

	size := number.Int64(number.NewNumber(vm.Get(ins.Size).Value))

	buf := make([]rune, size)
	charsRead := int64(0)
	for ; charsRead < size; charsRead++ {
		ch, _, err := f.Reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			vm.Raise(err.Error())
			return nil
		}

		buf[charsRead] = ch
	}

	vm.Set(ins.Str, asttest.NewLiteralString(string(buf[:charsRead])))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ReadString) String() string {
	return fmt.Sprintf("%s = %s.ReadString(%s)", ins.Str, ins.Fd, ins.Size)
}
