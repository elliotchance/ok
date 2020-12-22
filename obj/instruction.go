package obj

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

type Instruction string

func NewInstruction(ins vm.Instruction) Instruction {
	ty := reflect.TypeOf(ins).Elem()
	val := reflect.ValueOf(ins).Elem()
	instruction := []string{ty.String()[3:]}

	for i := 0; i < ty.NumField(); i++ {
		instruction = append(instruction, render(val.Field(i).Interface()))
	}

	return Instruction(strings.Join(instruction, " "))
}

func render(x interface{}) string {
	return fmt.Sprintf("%v", x)
	switch a := x.(type) {
	case nil:
		return "nil"

	case vm.Register:
		return string(a)

	case vm.Registers:
		registers := make([]string, len(a))
		for i, register := range a {
			registers[i] = string(register)
		}

		return strings.Join(registers, " ")

	case *ast.Literal:
		// TODO(elliot): Remove me
		return "" // a.Value
	}

	panic(reflect.TypeOf(x).String())
}
