package vm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// File is the root structure that will be serialized into the okc file.
type File struct {
	Tests []*CompiledTest `json:",omitempty"`

	// Types contains the type descriptions that can be referenced by some
	// instructions at runtime.
	Types types.Registry `json:",omitempty"`

	// Symbols contains literal values that can be referenced by instructions.
	Symbols map[SymbolRegister]*Symbol `json:",omitempty"`

	// Globals describes the global registers and unique names of the functions
	// that will initialize each package.
	Globals map[string]string `json:",omitempty"`
}

func (f *File) FuncByName(name string) *CompiledFunc {
	for _, fn := range f.Symbols {
		if fn.Func != nil && fn.Func.Name == name {
			return fn.Func
		}
	}

	return nil
}

// Store will create or replace the okc file for the provided package name.
func Store(file *File, packageName string) error {
	err := os.MkdirAll(Directory, 0755)
	if err != nil {
		return err
	}

	filePath := PathForPackage(packageName)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func Load(packageName string) (*File, error) {
	filePath := PathForPackage(packageName)
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var okcFile *File
	err = json.Unmarshal(jsonData, &okcFile)
	if err != nil {
		return nil, err
	}

	return okcFile, nil
}

func (f *File) AddType(ty *types.Type) TypeRegister {
	newType, err := f.Types.Add(ty)
	if err != nil {
		panic(err)
	}

	return TypeRegister(newType)
}

func (f *File) AddSymbolLiteral(lit *ast.Literal) SymbolRegister {
	// TODO(elliot): Dedup symbols here.
	key := SymbolRegister(fmt.Sprintf("%d", len(f.Symbols)))

	f.Symbols[key] = &Symbol{
		Type:  f.AddType(lit.Kind),
		Value: lit.Value,
	}

	return key
}

func (f *File) AddSymbolFunc(fn *CompiledFunc) SymbolRegister {
	key := SymbolRegister(fmt.Sprintf("%d", len(f.Symbols)))
	f.Symbols[key] = &Symbol{
		Type:      fn.Type,
		Interface: f.Types.Get(string(fn.Type)).Interface(),
		Func:      fn,
	}

	return key
}
