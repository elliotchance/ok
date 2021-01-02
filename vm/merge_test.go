package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	file1 := &vm.File{
		Globals: map[string]string{
			"$foo__bar__mypkg": "1234",
		},
		Types: types.Registry{
			"0": types.Number,
			"1": types.String,
			"2": types.NewFunc(nil, []*types.Type{types.NewRef("0")}),
		},
		Symbols: map[vm.SymbolRegister]*vm.Symbol{
			"0": {
				Type: "2",
				Func: &vm.CompiledFunc{
					Instructions: vm.NewInstructions(
						&vm.AssignSymbol{
							Result: "0",
							Symbol: "2", // number: 17
						},
						&vm.ArrayAlloc{
							Size: "0",
							Kind: "0", // types.Number
						},
					),
					UniqueName: "unique-1",
				},
			},
			"1": {
				Type:  "0",
				Value: "3.14",
			},
			"2": {
				Type:  "0",
				Value: "17",
			},
			"3": {
				Type:  "1",
				Value: "17",
			},
		},
	}

	file2 := &vm.File{
		Globals: map[string]string{
			"$foo__bar__secondpkg": "456",
		},
		Types: types.Registry{
			"0": types.String,
			"1": types.Number,
			"2": types.NewArray(types.NewRef("1")),
			"3": types.NewFunc(nil, []*types.Type{types.NewRef("0")}),
		},
		Symbols: map[vm.SymbolRegister]*vm.Symbol{
			"0": { // func() string
				Type: "3",
				Func: &vm.CompiledFunc{
					Instructions: vm.NewInstructions(
						&vm.AssignSymbol{
							Result: "0",
							Symbol: "1", // number: 17
						},
						&vm.ArrayAlloc{
							Size: "0",
							Kind: "0", // types.String
						},
						&vm.AssignSymbol{
							Result: "1",
							Symbol: "2", // number: 2.718
						},
						&vm.ArrayAlloc{
							Size: "1",
							Kind: "2", // types.NumberArray
						},
					),
					UniqueName: "unique-2",
				},
			},
			"1": {
				Type:  "1",
				Value: "17",
			},
			"2": {
				Type:  "1",
				Value: "2.718",
			},
		},
	}

	file := vm.Merge(file1, file2)
	assert.Empty(t, cmp.Diff(&vm.File{
		Globals: map[string]string{
			"$foo__bar__mypkg":     "1234",
			"$foo__bar__secondpkg": "456",
		},
		Types: types.Registry{
			"0": types.Number,
			"1": types.String,
			"2": types.NewFunc(nil, []*types.Type{types.NewRef("0")}),
			"3": types.NewArray(types.NewRef("0")),
			"4": types.NewFunc(nil, []*types.Type{types.NewRef("1")}),
		},
		Symbols: map[vm.SymbolRegister]*vm.Symbol{
			"0": { // func() number
				Type: "2",
				Func: &vm.CompiledFunc{
					Instructions: vm.NewInstructions(
						&vm.AssignSymbol{
							Result: "0",
							Symbol: "2", // number: 17
						},
						&vm.ArrayAlloc{
							Size: "0",
							Kind: "0", // types.Number
						},
					),
					UniqueName: "unique-1",
				},
			},
			"1": {
				Type:  "0",
				Value: "3.14",
			},
			"2": {
				Type:  "0",
				Value: "17",
			},
			"3": {
				Type:  "1",
				Value: "17",
			},
			"4": { // func() string
				Type: "4",
				Func: &vm.CompiledFunc{
					Instructions: vm.NewInstructions(
						&vm.AssignSymbol{
							Result: "0",
							Symbol: "5", // CHANGED: number: 17
						},
						&vm.ArrayAlloc{
							Size: "0",
							Kind: "1", // CHANGED: types.String
						},
						&vm.AssignSymbol{
							Result: "1",
							Symbol: "6", // CHANGED: number: 2.718
						},
						&vm.ArrayAlloc{
							Size: "1",
							Kind: "3", // CHANGED: types.NumberArray
						},
					),
					UniqueName: "unique-2",
				},
			},
			"5": {
				Type:  "0",
				Value: "17",
			},
			"6": {
				Type:  "0",
				Value: "2.718",
			},
		},
	}, file, cmpopts.IgnoreUnexported(vm.CompiledFunc{})))
}
