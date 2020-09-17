package vm

import "github.com/elliotchance/ok/ast"
import "github.com/elliotchance/ok/types"

func init() {
	Packages = map[string]*File{
		"error": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"Error": &CompiledFunc{
					Arguments: []string{"Error"},
					Instructions: []Instruction{
						&Return{Registers{"0"}},
					},
					Registers: 1,
					Variables: map[string]*types.Type{
						"Error": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: map[string]map[string]*types.Type{
						"Error": map[string]*types.Type{
							"Error": &types.Type{
								Kind: 7,
							},
						},
					},
				},
			},
			FuncDefs: map[string]*ast.Func{
				"Error": &ast.Func{
					Name: "Error",
					Arguments: []*ast.Argument{
						&ast.Argument{"Error", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Name: "Error",
						},
					},
					Pos: "lib/error/error.ok:2:1",
				},
			},
			Interfaces: map[string]map[string]*types.Type{
				"Error": map[string]*types.Type{
					"Error": &types.Type{
						Kind: 7,
					},
				},
			},
			Constants: nil,
		},
		"math": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"Abs": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/abs.ok:3:12"}, ""},
						&LessThanNumber{"x", "2", "3"},
						&JumpUnless{"3", 5},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&Subtract{"4", "x", "5"},
						&Return{Registers{"5"}},
						&Return{Registers{"x"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Cbrt": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/powers.ok:21:21"}, ""},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3", nil, nil, "lib/math/powers.ok:21:23"}, ""},
						&Divide{"2", "3", "4"},
						&Power{"x", "4", "5"},
						&Return{Registers{"5"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Ceil": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:3:16"}, ""},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:4:16"}, ""},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:8:12"}, ""},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 11},
						&Subtract{"x", "frac", "8"},
						&Return{Registers{"8"}},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:12:17"}, ""},
						&Subtract{"9", "frac", "10"},
						&Add{"x", "10", "11"},
						&Return{Registers{"11"}},
					},
					Registers: 11,
					Variables: map[string]*types.Type{
						"frac": &types.Type{
							Kind: 6,
						},
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Exp": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "2.71828182845904523536028747135266249775724709369995", nil, nil, "lib/math/powers.ok:4:9"}, ""},
						&Assign{"e", nil, "2"},
						&Power{"e", "x", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"e": &types.Type{
							Kind: 6,
						},
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Floor": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:17:16"}, ""},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:18:16"}, ""},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:22:12"}, ""},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 13},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:23:27"}, ""},
						&Add{"frac", "8", "9"},
						&Subtract{"x", "9", "10"},
						&Return{Registers{"10"}},
						&Subtract{"x", "frac", "11"},
						&Return{Registers{"11"}},
					},
					Registers: 11,
					Variables: map[string]*types.Type{
						"frac": &types.Type{
							Kind: 6,
						},
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Log10": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Log{"x", "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "10", nil, nil, "lib/math/log.ok:8:29"}, ""},
						&Log{"3", "4"},
						&Divide{"2", "4", "5"},
						&Return{Registers{"5"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"LogE": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Log{"x", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Variables: map[string]*types.Type{
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Pow": &CompiledFunc{
					Arguments: []string{"base", "power"},
					Instructions: []Instruction{
						&Power{"base", "power", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"base": &types.Type{
							Kind: 6,
						},
						"power": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Round": &CompiledFunc{
					Arguments: []string{"x", "prec"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "10", nil, nil, "lib/math/rounding.ok:32:15"}, ""},
						&Power{"3", "prec", "4"},
						&Assign{"p", nil, "4"},
						&Multiply{"x", "p", "5"},
						&Assign{"y", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:35:16"}, ""},
						&Remainder{"y", "6", "7"},
						&Assign{"diff", nil, "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.5", nil, nil, "lib/math/rounding.ok:36:16"}, ""},
						&GreaterThanEqualNumber{"diff", "8", "9"},
						&JumpUnless{"9", 15},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:37:22"}, ""},
						&Subtract{"10", "diff", "11"},
						&Add{"y", "11", "12"},
						&Divide{"12", "p", "13"},
						&Return{Registers{"13"}},
						&Subtract{"y", "diff", "14"},
						&Divide{"14", "p", "15"},
						&Return{Registers{"15"}},
					},
					Registers: 15,
					Variables: map[string]*types.Type{
						"diff": &types.Type{
							Kind: 6,
						},
						"p": &types.Type{
							Kind: 6,
						},
						"prec": &types.Type{
							Kind: 6,
						},
						"x": &types.Type{
							Kind: 6,
						},
						"y": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Sqrt": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.5", nil, nil, "lib/math/powers.ok:16:21"}, ""},
						&Power{"x", "2", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"x": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
			},
			FuncDefs: map[string]*ast.Func{
				"Abs": &ast.Func{
					Name: "Abs",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/abs.ok:2:1",
				},
				"Cbrt": &ast.Func{
					Name: "Cbrt",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/powers.ok:20:1",
				},
				"Ceil": &ast.Func{
					Name: "Ceil",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/rounding.ok:2:1",
				},
				"Exp": &ast.Func{
					Name: "Exp",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/powers.ok:2:1",
				},
				"Floor": &ast.Func{
					Name: "Floor",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/rounding.ok:16:1",
				},
				"Log10": &ast.Func{
					Name: "Log10",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/log.ok:7:1",
				},
				"LogE": &ast.Func{
					Name: "LogE",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/log.ok:2:1",
				},
				"Pow": &ast.Func{
					Name: "Pow",
					Arguments: []*ast.Argument{
						&ast.Argument{"base", &types.Type{
							Kind: 6,
						}},
						&ast.Argument{"power", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/powers.ok:10:1",
				},
				"Round": &ast.Func{
					Name: "Round",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
						&ast.Argument{"prec", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/rounding.ok:31:1",
				},
				"Sqrt": &ast.Func{
					Name: "Sqrt",
					Arguments: []*ast.Argument{
						&ast.Argument{"x", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/math/powers.ok:15:1",
				},
			},
			Interfaces: nil,
			Constants: map[string]*ast.Literal{
				"E": &ast.Literal{&types.Type{
					Kind: 6,
				}, "2.71828182845904523536028747135266249775724709369995957496696763", nil, nil, "lib/math/constants.ok:1:7"},
				"Ln10": &ast.Literal{&types.Type{
					Kind: 6,
				}, "2.30258509299404568401799145468436420760110148862877297603332790", nil, nil, "lib/math/constants.ok:11:8"},
				"Ln2": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.693147180559945309417232121458176568075500134360255254120680009", nil, nil, "lib/math/constants.ok:10:8"},
				"Phi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.61803398874989484820458683436563811772030917980576286213544862", nil, nil, "lib/math/constants.ok:3:7"},
				"Pi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "3.14159265358979323846264338327950288419716939937510582097494459", nil, nil, "lib/math/constants.ok:2:7"},
				"Sqrt2": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.41421356237309504880168872420969807856967187537694807317667974", nil, nil, "lib/math/constants.ok:5:11"},
				"SqrtE": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.64872127070012814684865078781416357165377610071014801157507931", nil, nil, "lib/math/constants.ok:6:11"},
				"SqrtPhi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.27201964951406896425242246173749149171560804184009624861664038", nil, nil, "lib/math/constants.ok:8:11"},
				"SqrtPi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.77245385090551602729816748334114518279754945612238712821380779", nil, nil, "lib/math/constants.ok:7:11"},
			},
		},
		"reflect": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"Call": &CompiledFunc{
					Arguments: []string{"fn", "args"},
					Instructions: []Instruction{
						&DynamicCall{"fn", "args", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"args": &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 2,
							},
						},
						"fn": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Get": &CompiledFunc{
					Arguments: []string{"obj", "prop"},
					Instructions: []Instruction{
						&Get{"obj", "prop", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"obj": &types.Type{
							Kind: 2,
						},
						"prop": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Interface": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Interface{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Variables: map[string]*types.Type{
						"value": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Kind": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Call{"Type", Registers{"value"}, Registers{"2"}},
						&Assign{"type", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "[]", nil, nil, "lib/reflect/kind.ok:7:30"}, ""},
						&Call{"hasPrefix", Registers{"type", "3"}, Registers{"4"}},
						&JumpUnless{"4", 7},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "array", nil, nil, "lib/reflect/kind.ok:8:20"}, ""},
						&Return{Registers{"5"}},
						&Jump{20},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 7,
						}, "{}", nil, nil, "lib/reflect/kind.ok:11:31"}, ""},
						&Interpolate{"6", Registers{"7"}},
						&Call{"hasPrefix", Registers{"type", "6"}, Registers{"8"}},
						&JumpUnless{"8", 14},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 7,
						}, "map", nil, nil, "lib/reflect/kind.ok:12:20"}, ""},
						&Return{Registers{"9"}},
						&Jump{20},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func(", nil, nil, "lib/reflect/kind.ok:15:30"}, ""},
						&Call{"hasPrefix", Registers{"type", "10"}, Registers{"11"}},
						&JumpUnless{"11", 20},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func", nil, nil, "lib/reflect/kind.ok:16:20"}, ""},
						&Return{Registers{"12"}},
						&Jump{20},
						&Return{Registers{"type"}},
					},
					Registers: 12,
					Variables: map[string]*types.Type{
						"type": &types.Type{
							Kind: 7,
						},
						"value": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Len": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Len{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Variables: map[string]*types.Type{
						"value": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Properties": &CompiledFunc{
					Arguments: []string{"obj"},
					Instructions: []Instruction{
						&Props{"obj", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Variables: map[string]*types.Type{
						"obj": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Set": &CompiledFunc{
					Arguments: []string{"obj", "prop", "value"},
					Instructions: []Instruction{
						&Set{"obj", "prop", "value", "4"},
						&Return{Registers{"4"}},
					},
					Registers: 4,
					Variables: map[string]*types.Type{
						"obj": &types.Type{
							Kind: 2,
						},
						"prop": &types.Type{
							Kind: 2,
						},
						"value": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"Type": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Type{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Variables: map[string]*types.Type{
						"value": &types.Type{
							Kind: 2,
						},
					},
					Interfaces: nil,
				},
				"hasPrefix": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Len{"s", "3"},
						&Len{"prefix", "4"},
						&LessThanNumber{"3", "4", "5"},
						&JumpUnless{"5", 5},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/reflect/strings.ok:6:16"}, ""},
						&Return{Registers{"6"}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/reflect/strings.ok:9:13"}, ""},
						&Assign{"i", nil, "7"},
						&Len{"prefix", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 19},
						&StringIndex{"s", "i", "10"},
						&StringIndex{"prefix", "i", "11"},
						&NotEqual{"10", "11", "12"},
						&JumpUnless{"12", 16},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/reflect/strings.ok:11:20"}, ""},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "14", "i"},
						&Jump{8},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/reflect/strings.ok:15:12"}, ""},
						&Return{Registers{"15"}},
					},
					Registers: 15,
					Variables: map[string]*types.Type{
						"i": &types.Type{
							Kind: 6,
						},
						"prefix": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
			},
			FuncDefs: map[string]*ast.Func{
				"Call": &ast.Func{
					Name: "Call",
					Arguments: []*ast.Argument{
						&ast.Argument{"fn", &types.Type{
							Kind: 2,
						}},
						&ast.Argument{"args", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 2,
							},
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 2,
							},
						},
					},
					Pos: "lib/reflect/call.ok:16:1",
				},
				"Get": &ast.Func{
					Name: "Get",
					Arguments: []*ast.Argument{
						&ast.Argument{"obj", &types.Type{
							Kind: 2,
						}},
						&ast.Argument{"prop", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 2,
						},
					},
					Pos: "lib/reflect/get.ok:15:1",
				},
				"Interface": &ast.Func{
					Name: "Interface",
					Arguments: []*ast.Argument{
						&ast.Argument{"value", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/reflect/interface.ok:10:1",
				},
				"Kind": &ast.Func{
					Name: "Kind",
					Arguments: []*ast.Argument{
						&ast.Argument{"value", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/reflect/kind.ok:3:1",
				},
				"Len": &ast.Func{
					Name: "Len",
					Arguments: []*ast.Argument{
						&ast.Argument{"value", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/reflect/len.ok:3:1",
				},
				"Properties": &ast.Func{
					Name: "Properties",
					Arguments: []*ast.Argument{
						&ast.Argument{"obj", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						},
					},
					Pos: "lib/reflect/props.ok:3:1",
				},
				"Set": &ast.Func{
					Name: "Set",
					Arguments: []*ast.Argument{
						&ast.Argument{"obj", &types.Type{
							Kind: 2,
						}},
						&ast.Argument{"prop", &types.Type{
							Kind: 2,
						}},
						&ast.Argument{"value", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 2,
						},
					},
					Pos: "lib/reflect/set.ok:16:1",
				},
				"Type": &ast.Func{
					Name: "Type",
					Arguments: []*ast.Argument{
						&ast.Argument{"value", &types.Type{
							Kind: 2,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/reflect/type.ok:8:1",
				},
				"hasPrefix": &ast.Func{
					Name: "hasPrefix",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"prefix", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 3,
						},
					},
					Pos: "lib/reflect/strings.ok:4:1",
				},
			},
			Interfaces: nil,
			Constants:  nil,
		},
		"strings": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"Contains": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Call{"Index", Registers{"s", "substr"}, Registers{"3"}},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/contains.ok:3:32"}, ""},
						&NotEqualNumber{"3", "4", "5"},
						&Return{Registers{"5"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"HasPrefix": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Len{"s", "3"},
						&Len{"prefix", "4"},
						&LessThanNumber{"3", "4", "5"},
						&JumpUnless{"5", 5},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:9:16"}, ""},
						&Return{Registers{"6"}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/contains.ok:12:13"}, ""},
						&Assign{"i", nil, "7"},
						&Len{"prefix", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 19},
						&StringIndex{"s", "i", "10"},
						&StringIndex{"prefix", "i", "11"},
						&NotEqual{"10", "11", "12"},
						&JumpUnless{"12", 16},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:14:20"}, ""},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "14", "i"},
						&Jump{8},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/contains.ok:18:12"}, ""},
						&Return{Registers{"15"}},
					},
					Registers: 15,
					Variables: map[string]*types.Type{
						"i": &types.Type{
							Kind: 6,
						},
						"prefix": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"HasSuffix": &CompiledFunc{
					Arguments: []string{"s", "suffix"},
					Instructions: []Instruction{
						&Len{"s", "3"},
						&Len{"suffix", "4"},
						&LessThanNumber{"3", "4", "5"},
						&JumpUnless{"5", 5},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:24:16"}, ""},
						&Return{Registers{"6"}},
						&Len{"s", "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/contains.ok:27:18"}, ""},
						&Subtract{"7", "8", "9"},
						&Assign{"j", nil, "9"},
						&Len{"suffix", "10"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/contains.ok:28:27"}, ""},
						&Subtract{"10", "11", "12"},
						&Assign{"i", nil, "12"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/contains.ok:28:35"}, ""},
						&GreaterThanEqualNumber{"i", "13", "14"},
						&JumpUnless{"14", 27},
						&StringIndex{"s", "j", "15"},
						&StringIndex{"suffix", "i", "16"},
						&NotEqual{"15", "16", "17"},
						&JumpUnless{"17", 22},
						&Assign{"18", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:30:20"}, ""},
						&Return{Registers{"18"}},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Subtract{"j", "19", "j"},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Subtract{"i", "20", "i"},
						&Jump{14},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/contains.ok:36:12"}, ""},
						&Return{Registers{"21"}},
					},
					Registers: 21,
					Variables: map[string]*types.Type{
						"i": &types.Type{
							Kind: 6,
						},
						"j": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"suffix": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Index": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:3:34"}, ""},
						&Call{"IndexAfter", Registers{"s", "substr", "3"}, Registers{"4"}},
						&Return{Registers{"4"}},
					},
					Registers: 4,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"IndexAfter": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:18:26"}, ""},
						&Call{"max", Registers{"offset", "4"}, Registers{"5"}},
						&Assign{"offset", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:20:22"}, ""},
						&Add{"offset", "6", "7"},
						&Assign{"i", nil, "7"},
						&Len{"s", "8"},
						&Len{"substr", "9"},
						&Subtract{"8", "9", "10"},
						&LessThanEqualNumber{"i", "10", "11"},
						&JumpUnless{"11", 33},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/index.ok:21:17"}, ""},
						&Assign{"found", nil, "12"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/index.ok:23:17"}, ""},
						&Assign{"j", nil, "13"},
						&Len{"substr", "14"},
						&LessThanNumber{"j", "14", "15"},
						&JumpUnless{"15", 28},
						&Add{"i", "j", "16"},
						&StringIndex{"s", "16", "17"},
						&StringIndex{"substr", "j", "18"},
						&NotEqual{"17", "18", "19"},
						&JumpUnless{"19", 25},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/index.ok:25:25"}, ""},
						&Assign{"found", nil, "20"},
						&Jump{28},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"j", "21", "j"},
						&Jump{15},
						&JumpUnless{"found", 30},
						&Return{Registers{"i"}},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "22", "i"},
						&Jump{8},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:35:12"}, ""},
						&Return{Registers{"23"}},
					},
					Registers: 23,
					Variables: map[string]*types.Type{
						"found": &types.Type{
							Kind: 3,
						},
						"i": &types.Type{
							Kind: 6,
						},
						"j": &types.Type{
							Kind: 6,
						},
						"offset": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Join": &CompiledFunc{
					Arguments: []string{"strings", "glue"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/join.ok:5:14"}, ""},
						&Assign{"result", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&NextArray{"strings", "4", "i", "s", "5"},
						&JumpUnless{"5", 10},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/join.ok:7:16"}, ""},
						&GreaterThanNumber{"i", "6", "7"},
						&JumpUnless{"7", 8},
						&Concat{"result", "glue", "result"},
						&Concat{"result", "s", "result"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 7,
					Variables: map[string]*types.Type{
						"glue": &types.Type{
							Kind: 7,
						},
						"i": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"strings": &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						},
					},
					Interfaces: nil,
				},
				"LastIndex": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Call{"Reverse", Registers{"s"}, Registers{"3"}},
						&Call{"Reverse", Registers{"substr"}, Registers{"4"}},
						&Call{"Index", Registers{"3", "4"}, Registers{"5"}},
						&Assign{"index", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:59:17"}, ""},
						&EqualNumber{"index", "6", "7"},
						&JumpUnless{"7", 8},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:60:16"}, ""},
						&Return{Registers{"8"}},
						&Len{"s", "9"},
						&Len{"substr", "10"},
						&Add{"index", "10", "11"},
						&Subtract{"9", "11", "12"},
						&Return{Registers{"12"}},
					},
					Registers: 12,
					Variables: map[string]*types.Type{
						"index": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"LastIndexBefore": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&Len{"s", "5"},
						&Call{"min", Registers{"offset", "5"}, Registers{"6"}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:79:45"}, ""},
						&Add{"6", "7", "8"},
						&Subtract{"4", "8", "9"},
						&Assign{"offset", nil, "9"},
						&Call{"Reverse", Registers{"s"}, Registers{"10"}},
						&Call{"Reverse", Registers{"substr"}, Registers{"11"}},
						&Call{"IndexAfter", Registers{"10", "11", "offset"}, Registers{"12"}},
						&Assign{"index", nil, "12"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:82:17"}, ""},
						&EqualNumber{"index", "13", "14"},
						&JumpUnless{"14", 15},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:83:16"}, ""},
						&Return{Registers{"15"}},
						&Len{"s", "16"},
						&Len{"substr", "17"},
						&Add{"index", "17", "18"},
						&Subtract{"16", "18", "19"},
						&Return{Registers{"19"}},
					},
					Registers: 19,
					Variables: map[string]*types.Type{
						"index": &types.Type{
							Kind: 6,
						},
						"offset": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"PadLeft": &CompiledFunc{
					Arguments: []string{"s", "pad", "toLen"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&GreaterThanEqualNumber{"4", "toLen", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/pad.ok:10:34"}, ""},
						&Equal{"pad", "6", "7"},
						&Or{"5", "7", "8"},
						&JumpUnless{"8", 6},
						&Return{Registers{"s"}},
						&Len{"s", "9"},
						&Subtract{"toLen", "9", "10"},
						&Call{"createPad", Registers{"pad", "10"}, Registers{"11"}},
						&Concat{"11", "s", "12"},
						&Return{Registers{"12"}},
					},
					Registers: 12,
					Variables: map[string]*types.Type{
						"pad": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"toLen": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"PadRight": &CompiledFunc{
					Arguments: []string{"s", "pad", "toLen"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&GreaterThanEqualNumber{"4", "toLen", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/pad.ok:20:34"}, ""},
						&Equal{"pad", "6", "7"},
						&Or{"5", "7", "8"},
						&JumpUnless{"8", 6},
						&Return{Registers{"s"}},
						&Len{"s", "9"},
						&Subtract{"toLen", "9", "10"},
						&Call{"createPad", Registers{"pad", "10"}, Registers{"11"}},
						&Concat{"s", "11", "12"},
						&Return{Registers{"12"}},
					},
					Registers: 12,
					Variables: map[string]*types.Type{
						"pad": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"toLen": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"Repeat": &CompiledFunc{
					Arguments: []string{"str", "times"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/repeat.ok:4:14"}, ""},
						&Assign{"result", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/repeat.ok:5:13"}, ""},
						&Assign{"i", nil, "4"},
						&LessThanNumber{"i", "times", "5"},
						&JumpUnless{"5", 9},
						&Concat{"result", "str", "result"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "6", "i"},
						&Jump{3},
						&Return{Registers{"result"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"i": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"str": &types.Type{
							Kind: 7,
						},
						"times": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"ReplaceAll": &CompiledFunc{
					Arguments: []string{"s", "find", "replace"},
					Instructions: []Instruction{
						&Call{"Split", Registers{"s", "find"}, Registers{"4"}},
						&Call{"Join", Registers{"4", "replace"}, Registers{"5"}},
						&Return{Registers{"5"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"find": &types.Type{
							Kind: 7,
						},
						"replace": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Reverse": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/reverse.ok:3:14"}, ""},
						&Assign{"result", nil, "2"},
						&Len{"s", "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/reverse.ok:4:22"}, ""},
						&Subtract{"3", "4", "5"},
						&Assign{"i", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/reverse.ok:4:30"}, ""},
						&GreaterThanEqualNumber{"i", "6", "7"},
						&JumpUnless{"7", 14},
						&StringIndex{"s", "i", "8"},
						&CastString{"8", "9"},
						&Concat{"result", "9", "result"},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Subtract{"i", "10", "i"},
						&Jump{6},
						&Return{Registers{"result"}},
					},
					Registers: 10,
					Variables: map[string]*types.Type{
						"i": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Split": &CompiledFunc{
					Arguments: []string{"s", "delimiter"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArrayAlloc{"3", "4", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"elements", nil, "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:10:21"}, ""},
						&Equal{"delimiter", "5", "6"},
						&JumpUnless{"6", 21},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:13:17"}, ""},
						&Assign{"i", nil, "7"},
						&Len{"s", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 20},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"10", "11", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&StringIndex{"s", "i", "13"},
						&CastString{"13", "14"},
						&ArraySet{"11", "12", "14"},
						&Append{"elements", "11", "elements"},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "15", "i"},
						&Jump{8},
						&Jump{58},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:17:19"}, ""},
						&Assign{"element", nil, "16"},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:18:17"}, ""},
						&Assign{"i", nil, "17"},
						&Len{"s", "18"},
						&LessThanNumber{"i", "18", "19"},
						&JumpUnless{"19", 53},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:19:45"}, ""},
						&Subtract{"i", "20", "21"},
						&Call{"IndexAfter", Registers{"s", "delimiter", "21"}, Registers{"22"}},
						&Subtract{"22", "i", "23"},
						&Assign{"24", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:19:55"}, ""},
						&EqualNumber{"23", "24", "25"},
						&JumpUnless{"25", 47},
						&Assign{"26", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"26", "27", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"28", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"27", "28", "element"},
						&Append{"elements", "27", "elements"},
						&Assign{"29", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:21:27"}, ""},
						&Assign{"element", nil, "29"},
						&Len{"delimiter", "30"},
						&Assign{"31", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:22:39"}, ""},
						&Subtract{"30", "31", "32"},
						&Add{"i", "32", "i"},
						&Jump{50},
						&StringIndex{"s", "i", "33"},
						&CastString{"33", "34"},
						&Concat{"element", "34", "element"},
						&Assign{"35", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "35", "i"},
						&Jump{26},
						&Assign{"36", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"36", "37", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"38", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"37", "38", "element"},
						&Append{"elements", "37", "elements"},
						&Return{Registers{"elements"}},
					},
					Registers: 38,
					Variables: map[string]*types.Type{
						"delimiter": &types.Type{
							Kind: 7,
						},
						"element": &types.Type{
							Kind: 7,
						},
						"elements": &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						},
						"i": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Substr": &CompiledFunc{
					Arguments: []string{"s", "fromIndex", "toIndex"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/substr.ok:4:9"}, ""},
						&Assign{"r", nil, "4"},
						&Assign{"i", nil, "fromIndex"},
						&LessThanNumber{"i", "toIndex", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "i", "6"},
						&CastString{"6", "7"},
						&Concat{"r", "7", "r"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "8", "i"},
						&Jump{2},
						&Return{Registers{"r"}},
					},
					Registers: 8,
					Variables: map[string]*types.Type{
						"fromIndex": &types.Type{
							Kind: 6,
						},
						"i": &types.Type{
							Kind: 6,
						},
						"r": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
						"toIndex": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"ToLower": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/case.ok:5:14"}, ""},
						&Assign{"result", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 4,
						}, "A", nil, nil, "lib/strings/case.ok:8:24"}, ""},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 4,
						}, "Z", nil, nil, "lib/strings/case.ok:8:44"}, ""},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "32", nil, nil, "lib/strings/case.ok:9:40"}, ""},
						&Add{"n", "13", "14"},
						&CastChar{"14", "15"},
						&CastString{"15", "16"},
						&Concat{"result", "16", "result"},
						&Jump{22},
						&CastString{"c", "17"},
						&Concat{"result", "17", "result"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 17,
					Variables: map[string]*types.Type{
						"c": &types.Type{
							Kind: 4,
						},
						"n": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"ToUpper": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/case.ok:22:14"}, ""},
						&Assign{"result", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 4,
						}, "a", nil, nil, "lib/strings/case.ok:25:24"}, ""},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 4,
						}, "z", nil, nil, "lib/strings/case.ok:25:44"}, ""},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "32", nil, nil, "lib/strings/case.ok:26:40"}, ""},
						&Subtract{"n", "13", "14"},
						&CastChar{"14", "15"},
						&CastString{"15", "16"},
						&Concat{"result", "16", "result"},
						&Jump{22},
						&CastString{"c", "17"},
						&Concat{"result", "17", "result"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 17,
					Variables: map[string]*types.Type{
						"c": &types.Type{
							Kind: 4,
						},
						"n": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"Trim": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Call{"TrimLeft", Registers{"s", "cutset"}, Registers{"3"}},
						&Call{"TrimRight", Registers{"3", "cutset"}, Registers{"4"}},
						&Return{Registers{"4"}},
					},
					Registers: 4,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"TrimLeft": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/trim.ok:4:18"}, ""},
						&Assign{"offset", nil, "3"},
						&Len{"s", "4"},
						&LessThanNumber{"offset", "4", "5"},
						&JumpUnless{"5", 15},
						&StringIndex{"s", "offset", "6"},
						&CastString{"6", "7"},
						&Call{"Index", Registers{"cutset", "7"}, Registers{"8"}},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/trim.ok:5:47"}, ""},
						&EqualNumber{"8", "9", "10"},
						&JumpUnless{"10", 12},
						&Call{"substrFrom", Registers{"s", "offset"}, Registers{"11"}},
						&Return{Registers{"11"}},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"offset", "12", "offset"},
						&Jump{2},
						&Return{Registers{"s"}},
					},
					Registers: 12,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"offset": &types.Type{
							Kind: 6,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"TrimPrefix": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Call{"HasPrefix", Registers{"s", "prefix"}, Registers{"3"}},
						&JumpUnless{"3", 4},
						&Len{"prefix", "4"},
						&Call{"substrFrom", Registers{"s", "4"}, Registers{"5"}},
						&Return{Registers{"5"}},
						&Return{Registers{"s"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"prefix": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"TrimRight": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Call{"Reverse", Registers{"s"}, Registers{"3"}},
						&Call{"TrimLeft", Registers{"3", "cutset"}, Registers{"4"}},
						&Call{"Reverse", Registers{"4"}, Registers{"5"}},
						&Return{Registers{"5"}},
					},
					Registers: 5,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"TrimSuffix": &CompiledFunc{
					Arguments: []string{"s", "suffix"},
					Instructions: []Instruction{
						&Call{"Reverse", Registers{"s"}, Registers{"3"}},
						&Call{"Reverse", Registers{"suffix"}, Registers{"4"}},
						&Call{"TrimPrefix", Registers{"3", "4"}, Registers{"5"}},
						&Call{"Reverse", Registers{"5"}, Registers{"6"}},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"suffix": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
				"createPad": &CompiledFunc{
					Arguments: []string{"pad", "toLen"},
					Instructions: []Instruction{
						&Len{"pad", "3"},
						&Divide{"toLen", "3", "4"},
						&Call{"Repeat", Registers{"pad", "4"}, Registers{"5"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/pad.ok:28:50"}, ""},
						&Call{"Substr", Registers{"5", "6", "toLen"}, Registers{"7"}},
						&Return{Registers{"7"}},
					},
					Registers: 7,
					Variables: map[string]*types.Type{
						"pad": &types.Type{
							Kind: 7,
						},
						"toLen": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"max": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&GreaterThanNumber{"a", "b", "3"},
						&JumpUnless{"3", 2},
						&Return{Registers{"a"}},
						&Return{Registers{"b"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"a": &types.Type{
							Kind: 6,
						},
						"b": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"min": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&LessThanNumber{"a", "b", "3"},
						&JumpUnless{"3", 2},
						&Return{Registers{"a"}},
						&Return{Registers{"b"}},
					},
					Registers: 3,
					Variables: map[string]*types.Type{
						"a": &types.Type{
							Kind: 6,
						},
						"b": &types.Type{
							Kind: 6,
						},
					},
					Interfaces: nil,
				},
				"substrFrom": &CompiledFunc{
					Arguments: []string{"s", "index"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/trim.ok:53:14"}, ""},
						&Assign{"result", nil, "3"},
						&Len{"s", "4"},
						&LessThanNumber{"index", "4", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "index", "6"},
						&CastString{"6", "7"},
						&Concat{"result", "7", "result"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"index", "8", "index"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 8,
					Variables: map[string]*types.Type{
						"index": &types.Type{
							Kind: 6,
						},
						"result": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Interfaces: nil,
				},
			},
			FuncDefs: map[string]*ast.Func{
				"Contains": &ast.Func{
					Name: "Contains",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"substr", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 3,
						},
					},
					Pos: "lib/strings/contains.ok:2:1",
				},
				"HasPrefix": &ast.Func{
					Name: "HasPrefix",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"prefix", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 3,
						},
					},
					Pos: "lib/strings/contains.ok:7:1",
				},
				"HasSuffix": &ast.Func{
					Name: "HasSuffix",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"suffix", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 3,
						},
					},
					Pos: "lib/strings/contains.ok:22:1",
				},
				"Index": &ast.Func{
					Name: "Index",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"substr", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:2:1",
				},
				"IndexAfter": &ast.Func{
					Name: "IndexAfter",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"substr", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"offset", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:17:1",
				},
				"Join": &ast.Func{
					Name: "Join",
					Arguments: []*ast.Argument{
						&ast.Argument{"strings", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&ast.Argument{"glue", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/join.ok:4:1",
				},
				"LastIndex": &ast.Func{
					Name: "LastIndex",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"substr", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:57:1",
				},
				"LastIndexBefore": &ast.Func{
					Name: "LastIndexBefore",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"substr", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"offset", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:76:1",
				},
				"PadLeft": &ast.Func{
					Name: "PadLeft",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"pad", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"toLen", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/pad.ok:9:1",
				},
				"PadRight": &ast.Func{
					Name: "PadRight",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"pad", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"toLen", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/pad.ok:19:1",
				},
				"Repeat": &ast.Func{
					Name: "Repeat",
					Arguments: []*ast.Argument{
						&ast.Argument{"str", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"times", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/repeat.ok:3:1",
				},
				"ReplaceAll": &ast.Func{
					Name: "ReplaceAll",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"find", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"replace", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/replace.ok:5:1",
				},
				"Reverse": &ast.Func{
					Name: "Reverse",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/reverse.ok:2:1",
				},
				"Split": &ast.Func{
					Name: "Split",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"delimiter", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						},
					},
					Pos: "lib/strings/split.ok:7:1",
				},
				"Substr": &ast.Func{
					Name: "Substr",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"fromIndex", &types.Type{
							Kind: 6,
						}},
						&ast.Argument{"toIndex", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/substr.ok:3:1",
				},
				"ToLower": &ast.Func{
					Name: "ToLower",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/case.ok:4:1",
				},
				"ToUpper": &ast.Func{
					Name: "ToUpper",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/case.ok:21:1",
				},
				"Trim": &ast.Func{
					Name: "Trim",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"cutset", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:21:1",
				},
				"TrimLeft": &ast.Func{
					Name: "TrimLeft",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"cutset", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:3:1",
				},
				"TrimPrefix": &ast.Func{
					Name: "TrimPrefix",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"prefix", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:32:1",
				},
				"TrimRight": &ast.Func{
					Name: "TrimRight",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"cutset", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:15:1",
				},
				"TrimSuffix": &ast.Func{
					Name: "TrimSuffix",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"suffix", &types.Type{
							Kind: 7,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:47:1",
				},
				"createPad": &ast.Func{
					Name: "createPad",
					Arguments: []*ast.Argument{
						&ast.Argument{"pad", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"toLen", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/pad.ok:27:1",
				},
				"max": &ast.Func{
					Name: "max",
					Arguments: []*ast.Argument{
						&ast.Argument{"a", &types.Type{
							Kind: 6,
						}},
						&ast.Argument{"b", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:48:1",
				},
				"min": &ast.Func{
					Name: "min",
					Arguments: []*ast.Argument{
						&ast.Argument{"a", &types.Type{
							Kind: 6,
						}},
						&ast.Argument{"b", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
					Pos: "lib/strings/index.ok:39:1",
				},
				"substrFrom": &ast.Func{
					Name: "substrFrom",
					Arguments: []*ast.Argument{
						&ast.Argument{"s", &types.Type{
							Kind: 7,
						}},
						&ast.Argument{"index", &types.Type{
							Kind: 6,
						}},
					},
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
					Pos: "lib/strings/trim.ok:52:1",
				},
			},
			Interfaces: nil,
			Constants:  nil,
		},
	}
}
