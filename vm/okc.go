package vm

import (
	"encoding/gob"
	"os"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// File is the root structure that will be serialized into the okc file.
type File struct {
	// Imports lists all the packages that this package relies on.
	Imports map[string]*types.Type

	Funcs map[string]*CompiledFunc
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
}

func (f *File) FuncByName(name string) *CompiledFunc {
	for _, fn := range f.Funcs {
		if fn.Name == name {
			return fn
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

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(file)
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
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	decoder := gob.NewDecoder(f)
	var okcFile File
	err = decoder.Decode(&okcFile)
	if err != nil {
		return nil, err
	}

	return &okcFile, nil
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
