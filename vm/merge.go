package vm

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/elliotchance/ok/types"
)

// Merge produces a new file with all symbols merged. Merge ignores imports and
// tests.
//
// NOTE: This will mangle the input files so they cannot be used after going
// through this process.
func Merge(files ...*File) *File {
	merged := &File{
		Types:   types.Registry{},
		Symbols: map[SymbolRegister]*Symbol{},
		Globals: map[string]string{},
	}

	for _, file := range files {
		// Simple merge of globals, these might conflict but they should always
		// resolve to the same unique function, so let's ignore any collisions
		// for now.
		for k, v := range file.Globals {
			merged.Globals[k] = v
		}

		// For a deterministic output we need to merge types in a predictable
		// order. This step is not needed once all registers are integers.
		typeKeys := make([]string, len(file.Types))
		i := 0
		for typeKey := range file.Types {
			typeKeys[i] = typeKey
			i++
		}
		sort.Strings(typeKeys)

		// Merge all of the types, creating a map of replacements.
		mappedTypes := map[TypeRegister]TypeRegister{}
		for _, existingTypeRegister := range typeKeys {
			ty := file.Types.Get(existingTypeRegister)
			newType, err := merged.Types.Add(ty)
			if err != nil {
				panic(err)
			}

			newTypeRegister := TypeRegister(newType)
			mappedTypes[TypeRegister(existingTypeRegister)] = newTypeRegister
		}

		// Once again we need to extract the symbol keys so that the merge is
		// deterministic.
		symbolKeys := make([]string, len(file.Symbols))
		i = 0
		for symbolKey := range file.Symbols {
			symbolKeys[i] = string(symbolKey)
			i++
		}
		sort.Strings(symbolKeys)

		// Now merge the symbols. We cannot translate the symbol registers in
		// this step because function symbols may reference symbols loaded in
		// the future. Instead, we will track the symbol translation in the same
		// way as type, but also record all the merged symbols so we can go back
		// an correct them in the next step.
		mappedSymbols := map[SymbolRegister]SymbolRegister{}
		var symbolsToRevisit []SymbolRegister
		for _, symbolKey := range symbolKeys {
			symbol := file.Symbols[SymbolRegister(symbolKey)]

			// TODO(elliot): It would be great to dedup the symbols here. That's
			//  not actually trivial because the symbols and types are changing
			//  underneath.

			newSymbolKey := SymbolRegister(fmt.Sprintf("%d", len(merged.Symbols)))
			mappedSymbols[SymbolRegister(symbolKey)] = newSymbolKey

			// Translate instructions.
			compileFunc := symbol.Func
			if symbol.Func != nil {
				// We only need to revisit functions.
				symbolsToRevisit = append(symbolsToRevisit, newSymbolKey)
			}

			ty := file.Types.Get(string(symbol.Type))
			newType, err := merged.Types.Add(ty)
			if err != nil {
				panic(err)
			}

			newTypeRegister := TypeRegister(newType)
			mappedTypes[symbol.Type] = newTypeRegister

			merged.Symbols[newSymbolKey] = &Symbol{
				Type:  newTypeRegister,
				Value: symbol.Value,
				Func:  compileFunc,
			}
		}

		// Finally, fix up the symbol and type references within the
		// instructions.
		for _, symbolRegister := range symbolsToRevisit {
			symbol := merged.Symbols[symbolRegister]
			instructions := NewInstructions(symbol.Func.Instructions.Instructions...)
			for i := 0; i < len(instructions.Instructions); i++ {
				ty := reflect.TypeOf(instructions.Instructions[i]).Elem()
				val := reflect.ValueOf(instructions.Instructions[i]).Elem()
				for j := 0; j < ty.NumField(); j++ {
					field := val.Field(j)
					switch ty.Field(j).Type.Name() {
					case "TypeRegister":
						existing := field.Interface().(TypeRegister)
						field.Set(reflect.ValueOf(mappedTypes[existing]))

					case "SymbolRegister":
						existing := field.Interface().(SymbolRegister)
						field.Set(reflect.ValueOf(mappedSymbols[existing]))
					}
				}
			}

			symbol.Func.Instructions = instructions
		}
	}

	return merged
}
