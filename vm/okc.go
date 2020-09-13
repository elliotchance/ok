package vm

import (
	"encoding/gob"
	"os"

	"github.com/elliotchance/ok/ast"
)

// File is the root structure that will be serialized into the okc file.
type File struct {
	// Imports lists all the packages that this package relies on.
	Imports map[string]struct{}

	Funcs      map[string]*CompiledFunc
	FuncDefs   map[string]*ast.Func
	Tests      []*CompiledTest
	Interfaces map[string]map[string]string
	Constants  map[string]*ast.Literal

	// ImportedFuncs are ephemeral. They will be stripped out before saving to
	// the file and will be nil when the file is loaded. They provide a way to
	// tell the compiler about functions imported by other packages at compile
	// time.
	//
	// The VM will find them at runtime because the dependent package will have
	// already been loaded because of Imports.
	ImportedFuncs map[string]*ast.Func
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
	file.ImportedFuncs = nil
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
