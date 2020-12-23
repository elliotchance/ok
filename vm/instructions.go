package vm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Instructions struct {
	Instructions []Instruction
}

func NewInstructions(instructions ...Instruction) *Instructions {
	return &Instructions{
		Instructions: instructions,
	}
}

func (ins *Instructions) MarshalJSON() ([]byte, error) {
	out := make([]string, len(ins.Instructions))
	for i := range ins.Instructions {
		out[i] = marshalInstruction(ins.Instructions[i])
	}

	return json.Marshal(out)
}

var decodeTypes = []interface{}{
	Add{},
	And{},
	Append{},
	ArrayAlloc{},
	ArrayGet{},
	ArraySet{},
	Assert{},
	Assign{},
	AssignFunc{},
	AssignSymbol{},
	Call{},
	CastChar{},
	CastData{},
	CastNumber{},
	CastString{},
	Close{},
	Combine{},
	Concat{},
	Divide{},
	DynamicCall{},
	EnvGet{},
	EnvSet{},
	EnvUnset{},
	EqualNumber{},
	Equal{},
	Exit{},
	Finally{},
	FromUnix{},
	Get{},
	GreaterThanEqualNumber{},
	GreaterThanEqualString{},
	GreaterThanNumber{},
	GreaterThanString{},
	Info{},
	Interface{},
	Interpolate{},
	Is{},
	JumpUnless{},
	Jump{},
	Len{},
	LessThanEqualNumber{},
	LessThanEqualString{},
	LessThanNumber{},
	LessThanString{},
	LoadPackage{},
	Log{},
	MapAlloc{},
	MapGet{},
	MapSet{},
	Mkdir{},
	Multiply{},
	NextArray{},
	NextMap{},
	NextString{},
	NotEqualNumber{},
	NotEqual{},
	Not{},
	Now{},
	On{},
	Open{},
	Or{},
	ParentScope{},
	Power{},
	Print{},
	Props{},
	Raise{},
	Rand{},
	ReadData{},
	ReadString{},
	Remainder{},
	Remove{},
	Rename{},
	Return{},
	Seek{},
	Set{},
	Sleep{},
	Stack{},
	StringIndex{},
	Subtract{},
	Type{},
	Unix{},
	Write{},
}

func (ins *Instructions) UnmarshalJSON(data []byte) error {
	var raw []string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for i := range raw {
		instruction := unmarshalInstruction(raw[i])
		ins.Instructions = append(ins.Instructions, instruction)
	}

	return nil
}

func marshalInstruction(ins Instruction) string {
	ty := reflect.TypeOf(ins).Elem()
	val := reflect.ValueOf(ins).Elem()
	parts := make([]string, 1+ty.NumField())
	parts[0] = ty.Name()

	for i := 0; i < ty.NumField(); i++ {
		parts[i+1] = argString(val.Field(i).Interface())
	}

	return strings.Join(parts, " ")
}

func argString(x interface{}) string {
	switch a := x.(type) {
	case Register:
		return string(a)
	case TypeRegister:
		return string(a)
	case SymbolRegister:
		return string(a)
	case Registers:
		registers := make([]string, len(a))
		for i := range a {
			registers[i] = string(a[i])
		}

		return strings.Join(registers, ",")
	}

	return fmt.Sprintf("%v", x)
}

func unmarshalInstruction(ins string) Instruction {
	parts := strings.Split(ins, " ")

	for j := range decodeTypes {
		ty := reflect.TypeOf(decodeTypes[j])
		if ty.Name() == parts[0] {
			dest := reflect.New(ty).Elem()
			for i := 0; i < ty.NumField(); i++ {
				f := dest.Field(i)
				switch f.Type().Name() {
				case "Register":
					f.Set(reflect.ValueOf(Register(parts[i+1])))
				case "TypeRegister":
					f.Set(reflect.ValueOf(TypeRegister(parts[i+1])))
				case "SymbolRegister":
					f.Set(reflect.ValueOf(SymbolRegister(parts[i+1])))

				case "Registers":
					if parts[i+1] != "" {
						var registers Registers
						for _, register := range strings.Split(parts[i+1], ",") {
							registers = append(registers, Register(register))
						}
						f.Set(reflect.ValueOf(registers))
					}

				case "int":
					intVal, err := strconv.Atoi(parts[i+1])
					if err != nil {
						panic(err)
					}
					f.Set(reflect.ValueOf(intVal))

				case "bool":
					f.Set(reflect.ValueOf(parts[i+1] == "true"))

				default:
					f.Set(reflect.ValueOf(parts[i+1]))
				}
			}

			return dest.Addr().Interface().(Instruction)
		}
	}

	panic(ins)
}
