package obj

import (
	"fmt"

	"github.com/elliotchance/ok/vm"
)

type Symbol struct {
	Type         string
	Value        string        `json:",omitempty"`
	Instructions []Instruction `json:",omitempty"`
}

type Symbols map[string]Symbol

func (s Symbols) Add(symbol Symbol) {
	key := fmt.Sprintf("%d", len(s))
	s[key] = symbol
}

func NewSymbolFromCompiledFunc(fn *vm.CompiledFunc) Symbol {
	instructions := make([]Instruction, len(fn.Instructions))
	for i, ins := range fn.Instructions {
		instructions[i] = NewInstruction(ins)
	}

	return Symbol{
		Type:         fn.Type.String(),
		Instructions: instructions,
	}
}
