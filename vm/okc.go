package vm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// File is the root structure that will be serialized into the okc file.
type File struct {
	// Imports lists all the packages that this package relies on.
	Imports map[string]*types.Type

	Tests []*CompiledTest

	// Constants contains the package-level constants. These would also appear
	// as values within PackageFunc, however it's not trivial for the compiler
	// to extract these. Also, since constants cannot be modified, there's no
	// need to create some complex logic to retrieved them at compile or
	// runtime.
	//
	// It's totally fine (and probably the best choice) to have the compiler
	// substitute these values as literals into expressions. Especially for when
	// the compiler starts supporting simplifying expressions at compile time.
	//
	// Constants can only be scalars and all scalars.
	Constants map[string]*ast.Literal

	PackageFunc *CompiledFunc

	// Types contains the type descriptions that can be referenced by some
	// instructions at runtime.
	Types map[TypeRegister]*types.Type

	// Symbols contains literal values that can be referenced by instructions.
	Symbols map[SymbolRegister]*Symbol
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
	// Ignore packages that are build in (standard library).
	if p, ok := Packages[packageName]; ok {
		return p, nil
	}

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

// Interface is useful when we need to lookup root elements (constants,
// functions etc) for the package.
func (f *File) Interface() *types.Type {
	return f.PackageFunc.Type.Returns[0]
}

func (f *File) AddType(kind *types.Type) TypeRegister {
	if kind == nil {
		return NoTypeRegister
	}

	// TODO(elliot): Dedup types here.
	key := TypeRegister(kind.String())
	f.Types[key] = kind

	return key
}

func (f *File) AddSymbolLiteral(lit *ast.Literal) SymbolRegister {
	// TODO(elliot): Dedup symbols here.
	// TODO(elliot): Remove the newline replace. This is just so the value
	//  doesn't cause libgen to produce mangled code.
	key := SymbolRegister(strings.ReplaceAll(lit.Kind.String()+lit.Value, "\n", "\\n"))

	// TODO(elliot): This is a nasty hack to get around some rounding that's
	//  happening at ultra high precisions, like the constants in the math
	//  library. Remove this code to see the failure.
	//if lit.Kind.Kind == types.KindNumber && len(key) > 20 {
	//	key = key[:20]
	//}

	f.Symbols[key] = &Symbol{
		Type:  lit.Kind.String(),
		Value: lit.Value,
	}

	return key
}

func (f *File) AddSymbolFunc(fn *CompiledFunc) {
	f.Symbols[SymbolRegister(fn.UniqueName)] = &Symbol{
		Type: fn.Type.String(),
		Func: fn,
	}
}
