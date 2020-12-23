package vm

import (
	"encoding/json"
	"reflect"
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
	out := make([]map[string]interface{}, len(ins.Instructions))
	for i := range ins.Instructions {
		mapData, err := json.Marshal(ins.Instructions[i])
		if err != nil {
			panic(err)
		}

		var m map[string]interface{}
		err = json.Unmarshal(mapData, &m)
		if err != nil {
			panic(err)
		}

		m["Instruction"] = reflect.TypeOf(ins.Instructions[i]).Elem().Name()
		out[i] = m
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
	var raw []map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for i := range raw {
		// TODO(elliot): Horrible, I know. Just temporary.
		for j := range decodeTypes {
			ty := reflect.TypeOf(decodeTypes[j])
			if ty.Name() == raw[i]["Instruction"] {
				insData, err := json.Marshal(raw[i])
				if err != nil {
					return err
				}

				dest := reflect.New(ty).Interface()
				err = json.Unmarshal(insData, dest)
				if err != nil {
					return err
				}

				ins.Instructions = append(ins.Instructions, dest.(Instruction))
			}
		}
	}

	return nil
}
