package vm

import "github.com/elliotchance/ok/ast"
import "github.com/elliotchance/ok/types"

func init() {
	Packages = map[string]*File{
		"error": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					},
					Name:       "Error",
					UniqueName: "1",
				},
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					},
					Name:       "Error",
					UniqueName: "1",
				},
			},
			Constants: nil,
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"1"},
					&Assign{"1", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"2"},
					&Assign{"Error", nil, "2"},
					&Return{Registers{"0"}},
				},
				Registers: 2,
				Variables: map[string]*types.Type{
					"1": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					},
					"Error": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Error",
								Properties: map[string]*types.Type{
									"Error": &types.Type{
										Kind: 7,
									},
								},
							},
						},
					},
				},
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__error",
							Properties: map[string]*types.Type{
								"Error": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Error",
											Properties: map[string]*types.Type{
												"Error": &types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Name: "..__error",
			},
		},
		"math": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Abs",
					UniqueName: "1",
				},
				"10": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Round",
					UniqueName: "10",
				},
				"2": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LogE",
					UniqueName: "2",
				},
				"3": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Log10",
					UniqueName: "3",
				},
				"4": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Exp",
					UniqueName: "4",
				},
				"5": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Pow",
					UniqueName: "5",
				},
				"6": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Sqrt",
					UniqueName: "6",
				},
				"7": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Cbrt",
					UniqueName: "7",
				},
				"8": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Ceil",
					UniqueName: "8",
				},
				"9": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Floor",
					UniqueName: "9",
				},
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Abs",
					UniqueName: "1",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Cbrt",
					UniqueName: "7",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Ceil",
					UniqueName: "8",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Exp",
					UniqueName: "4",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Floor",
					UniqueName: "9",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Log10",
					UniqueName: "3",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LogE",
					UniqueName: "2",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Pow",
					UniqueName: "5",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Round",
					UniqueName: "10",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Sqrt",
					UniqueName: "6",
				},
			},
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
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 6,
					}, "2.71828182845904523536028747135266249775724709369995957496696763", nil, nil, "lib/math/constants.ok:1:7"}, ""},
					&Assign{"E", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 6,
					}, "2.30258509299404568401799145468436420760110148862877297603332790", nil, nil, "lib/math/constants.ok:11:8"}, ""},
					&Assign{"Ln10", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 6,
					}, "0.693147180559945309417232121458176568075500134360255254120680009", nil, nil, "lib/math/constants.ok:10:8"}, ""},
					&Assign{"Ln2", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.61803398874989484820458683436563811772030917980576286213544862", nil, nil, "lib/math/constants.ok:3:7"}, ""},
					&Assign{"Phi", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 6,
					}, "3.14159265358979323846264338327950288419716939937510582097494459", nil, nil, "lib/math/constants.ok:2:7"}, ""},
					&Assign{"Pi", nil, "5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.41421356237309504880168872420969807856967187537694807317667974", nil, nil, "lib/math/constants.ok:5:11"}, ""},
					&Assign{"Sqrt2", nil, "6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.64872127070012814684865078781416357165377610071014801157507931", nil, nil, "lib/math/constants.ok:6:11"}, ""},
					&Assign{"SqrtE", nil, "7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.27201964951406896425242246173749149171560804184009624861664038", nil, nil, "lib/math/constants.ok:8:11"}, ""},
					&Assign{"SqrtPhi", nil, "8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.77245385090551602729816748334114518279754945612238712821380779", nil, nil, "lib/math/constants.ok:7:11"}, ""},
					&Assign{"SqrtPi", nil, "9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"10"},
					&Assign{"1", nil, "10"},
					&Assign{"11", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "10", nil, nil, ""}, ""},
					&ParentScope{"11"},
					&Assign{"10", nil, "11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"12"},
					&Assign{"2", nil, "12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"13"},
					&Assign{"3", nil, "13"},
					&Assign{"14", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"14"},
					&Assign{"4", nil, "14"},
					&Assign{"15", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"15"},
					&Assign{"5", nil, "15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"16"},
					&Assign{"6", nil, "16"},
					&Assign{"17", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"17"},
					&Assign{"7", nil, "17"},
					&Assign{"18", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"18"},
					&Assign{"8", nil, "18"},
					&Assign{"19", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"19"},
					&Assign{"9", nil, "19"},
					&Assign{"20", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"20"},
					&Assign{"Abs", nil, "20"},
					&Assign{"21", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"21"},
					&Assign{"Cbrt", nil, "21"},
					&Assign{"22", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"22"},
					&Assign{"Ceil", nil, "22"},
					&Assign{"23", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"23"},
					&Assign{"Exp", nil, "23"},
					&Assign{"24", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"24"},
					&Assign{"Floor", nil, "24"},
					&Assign{"25", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"25"},
					&Assign{"Log10", nil, "25"},
					&Assign{"26", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"26"},
					&Assign{"LogE", nil, "26"},
					&Assign{"27", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"27"},
					&Assign{"Pow", nil, "27"},
					&Assign{"28", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "10", nil, nil, ""}, ""},
					&ParentScope{"28"},
					&Assign{"Round", nil, "28"},
					&Assign{"29", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"29"},
					&Assign{"Sqrt", nil, "29"},
					&Return{Registers{"0"}},
				},
				Registers: 29,
				Variables: map[string]*types.Type{
					"1": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"10": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"2": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"3": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"4": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"5": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"6": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"7": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"8": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"9": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Abs": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Cbrt": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Ceil": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"E": &types.Type{
						Kind: 6,
					},
					"Exp": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Floor": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Ln10": &types.Type{
						Kind: 6,
					},
					"Ln2": &types.Type{
						Kind: 6,
					},
					"Log10": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"LogE": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Phi": &types.Type{
						Kind: 6,
					},
					"Pi": &types.Type{
						Kind: 6,
					},
					"Pow": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Round": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Sqrt": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Sqrt2": &types.Type{
						Kind: 6,
					},
					"SqrtE": &types.Type{
						Kind: 6,
					},
					"SqrtPhi": &types.Type{
						Kind: 6,
					},
					"SqrtPi": &types.Type{
						Kind: 6,
					},
				},
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__math",
							Properties: map[string]*types.Type{
								"Abs": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Cbrt": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Ceil": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"E": &types.Type{
									Kind: 6,
								},
								"Exp": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Floor": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Ln10": &types.Type{
									Kind: 6,
								},
								"Ln2": &types.Type{
									Kind: 6,
								},
								"Log10": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"LogE": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Phi": &types.Type{
									Kind: 6,
								},
								"Pi": &types.Type{
									Kind: 6,
								},
								"Pow": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Round": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Sqrt": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Sqrt2": &types.Type{
									Kind: 6,
								},
								"SqrtE": &types.Type{
									Kind: 6,
								},
								"SqrtPhi": &types.Type{
									Kind: 6,
								},
								"SqrtPi": &types.Type{
									Kind: 6,
								},
							},
						},
					},
				},
				Name: "..__math",
			},
		},
		"reflect": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					},
					Name:       "Call",
					UniqueName: "1",
				},
				"2": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					Name:       "Get",
					UniqueName: "2",
				},
				"3": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Interface",
					UniqueName: "3",
				},
				"4": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 2,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "9", nil, nil, ""}, ""},
						&Call{"*2", Registers{"value"}, Registers{"3"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"type", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "[]", nil, nil, "lib/reflect/kind.ok:7:30"}, ""},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*5", Registers{"type", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"6", 9},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 7,
						}, "array", nil, nil, "lib/reflect/kind.ok:8:20"}, ""},
						&Return{Registers{"7"}},
						&Jump{24},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 7,
						}, "{}", nil, nil, "lib/reflect/kind.ok:11:31"}, ""},
						&Interpolate{"8", Registers{"9"}},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*10", Registers{"type", "8"}, Registers{"11"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"11", 17},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 7,
						}, "map", nil, nil, "lib/reflect/kind.ok:12:20"}, ""},
						&Return{Registers{"12"}},
						&Jump{24},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func(", nil, nil, "lib/reflect/kind.ok:15:30"}, ""},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*14", Registers{"type", "13"}, Registers{"15"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"15", 24},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func", nil, nil, "lib/reflect/kind.ok:16:20"}, ""},
						&Return{Registers{"16"}},
						&Jump{24},
						&Return{Registers{"type"}},
					},
					Registers: 16,
					Variables: map[string]*types.Type{
						"type": &types.Type{
							Kind: 7,
						},
						"value": &types.Type{
							Kind: 2,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Kind",
					UniqueName: "4",
				},
				"5": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Len",
					UniqueName: "5",
				},
				"6": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					Name:       "Properties",
					UniqueName: "6",
				},
				"7": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					Name:       "Set",
					UniqueName: "7",
				},
				"8": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "hasPrefix",
					UniqueName: "8",
				},
				"9": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Type",
					UniqueName: "9",
				},
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					},
					Name:       "Call",
					UniqueName: "1",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					Name:       "Get",
					UniqueName: "2",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Interface",
					UniqueName: "3",
				},
				"Kind": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 2,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "9", nil, nil, ""}, ""},
						&Call{"*2", Registers{"value"}, Registers{"3"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"type", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "[]", nil, nil, "lib/reflect/kind.ok:7:30"}, ""},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*5", Registers{"type", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"6", 9},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 7,
						}, "array", nil, nil, "lib/reflect/kind.ok:8:20"}, ""},
						&Return{Registers{"7"}},
						&Jump{24},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 7,
						}, "{}", nil, nil, "lib/reflect/kind.ok:11:31"}, ""},
						&Interpolate{"8", Registers{"9"}},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*10", Registers{"type", "8"}, Registers{"11"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"11", 17},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 7,
						}, "map", nil, nil, "lib/reflect/kind.ok:12:20"}, ""},
						&Return{Registers{"12"}},
						&Jump{24},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func(", nil, nil, "lib/reflect/kind.ok:15:30"}, ""},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*14", Registers{"type", "13"}, Registers{"15"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"15", 24},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func", nil, nil, "lib/reflect/kind.ok:16:20"}, ""},
						&Return{Registers{"16"}},
						&Jump{24},
						&Return{Registers{"type"}},
					},
					Registers: 16,
					Variables: map[string]*types.Type{
						"type": &types.Type{
							Kind: 7,
						},
						"value": &types.Type{
							Kind: 2,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Kind",
					UniqueName: "4",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Len",
					UniqueName: "5",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					Name:       "Properties",
					UniqueName: "6",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					Name:       "Set",
					UniqueName: "7",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Type",
					UniqueName: "9",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "hasPrefix",
					UniqueName: "8",
				},
			},
			Constants: nil,
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"1"},
					&Assign{"1", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"2"},
					&Assign{"2", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"3"},
					&Assign{"3", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"4"},
					&Assign{"4", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"5"},
					&Assign{"5", nil, "5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"6"},
					&Assign{"6", nil, "6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"7"},
					&Assign{"7", nil, "7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"8"},
					&Assign{"8", nil, "8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"9"},
					&Assign{"9", nil, "9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"10"},
					&Assign{"Call", nil, "10"},
					&Assign{"11", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"11"},
					&Assign{"Get", nil, "11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"12"},
					&Assign{"Interface", nil, "12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"13"},
					&Assign{"Kind", nil, "13"},
					&Assign{"14", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"14"},
					&Assign{"Len", nil, "14"},
					&Assign{"15", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"15"},
					&Assign{"Properties", nil, "15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"16"},
					&Assign{"Set", nil, "16"},
					&Assign{"17", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"17"},
					&Assign{"Type", nil, "17"},
					&Assign{"18", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"18"},
					&Assign{"hasPrefix", nil, "18"},
					&Return{Registers{"0"}},
				},
				Registers: 18,
				Variables: map[string]*types.Type{
					"1": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					},
					"2": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					"3": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"4": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"5": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"6": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					"7": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					"8": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"9": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Call": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 2,
								},
							},
						},
					},
					"Get": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					"Interface": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Kind": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Len": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Properties": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					"Set": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
					},
					"Type": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 2,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"hasPrefix": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
				},
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__reflect",
							Properties: map[string]*types.Type{
								"Call": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Kind: 2,
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Kind: 2,
											},
										},
									},
								},
								"Get": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
								},
								"Interface": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Kind": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Len": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Properties": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Kind: 7,
											},
										},
									},
								},
								"Set": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
										&types.Type{
											Kind: 2,
										},
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
								},
								"Type": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 2,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						},
					},
				},
				Name: "..__reflect",
			},
		},
		"strings": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ToLower",
					UniqueName: "1",
				},
				"10": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*5", Registers{"substr"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:59:17"}, ""},
						&EqualNumber{"index", "9", "10"},
						&JumpUnless{"10", 11},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:60:16"}, ""},
						&Return{Registers{"11"}},
						&Len{"s", "12"},
						&Len{"substr", "13"},
						&Add{"index", "13", "14"},
						&Subtract{"12", "14", "15"},
						&Return{Registers{"15"}},
					},
					Registers: 15,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LastIndex",
					UniqueName: "10",
				},
				"11": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&Len{"s", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*6", Registers{"offset", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:79:45"}, ""},
						&Add{"7", "8", "9"},
						&Subtract{"4", "9", "10"},
						&Assign{"offset", nil, "10"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*11", Registers{"s"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*13", Registers{"substr"}, Registers{"14"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*15", Registers{"12", "14", "offset"}, Registers{"16"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "16"},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:82:17"}, ""},
						&EqualNumber{"index", "17", "18"},
						&JumpUnless{"18", 19},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:83:16"}, ""},
						&Return{Registers{"19"}},
						&Len{"s", "20"},
						&Len{"substr", "21"},
						&Add{"index", "21", "22"},
						&Subtract{"20", "22", "23"},
						&Return{Registers{"23"}},
					},
					Registers: 23,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LastIndexBefore",
					UniqueName: "11",
				},
				"12": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Join",
					UniqueName: "12",
				},
				"13": &CompiledFunc{
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
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "15", nil, nil, ""}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"12", "s", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "PadLeft",
					UniqueName: "13",
				},
				"14": &CompiledFunc{
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
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "15", nil, nil, ""}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"s", "12", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "PadRight",
					UniqueName: "14",
				},
				"15": &CompiledFunc{
					Arguments: []string{"pad", "toLen"},
					Instructions: []Instruction{
						&Len{"pad", "3"},
						&Divide{"toLen", "3", "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "16", nil, nil, ""}, ""},
						&Call{"*5", Registers{"pad", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/pad.ok:28:50"}, ""},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "20", nil, nil, ""}, ""},
						&Call{"*8", Registers{"6", "7", "toLen"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Variables: map[string]*types.Type{
						"pad": &types.Type{
							Kind: 7,
						},
						"toLen": &types.Type{
							Kind: 6,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "createPad",
					UniqueName: "15",
				},
				"16": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Repeat",
					UniqueName: "16",
				},
				"17": &CompiledFunc{
					Arguments: []string{"s", "find", "replace"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Kind: 7,
									},
								},
							},
						}, "19", nil, nil, ""}, ""},
						&Call{"*4", Registers{"s", "find"}, Registers{"5"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Kind: 7,
									},
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "12", nil, nil, ""}, ""},
						&Call{"*6", Registers{"5", "replace"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
					},
					Registers: 7,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ReplaceAll",
					UniqueName: "17",
				},
				"18": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Reverse",
					UniqueName: "18",
				},
				"19": &CompiledFunc{
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
						&Jump{59},
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
						&JumpUnless{"19", 54},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:19:45"}, ""},
						&Subtract{"i", "20", "21"},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*22", Registers{"s", "delimiter", "21"}, Registers{"23"}, &types.Type{
							Kind: 2,
						}},
						&Subtract{"23", "i", "24"},
						&Assign{"25", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:19:55"}, ""},
						&EqualNumber{"24", "25", "26"},
						&JumpUnless{"26", 48},
						&Assign{"27", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"27", "28", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"29", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"28", "29", "element"},
						&Append{"elements", "28", "elements"},
						&Assign{"30", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:21:27"}, ""},
						&Assign{"element", nil, "30"},
						&Len{"delimiter", "31"},
						&Assign{"32", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:22:39"}, ""},
						&Subtract{"31", "32", "33"},
						&Add{"i", "33", "i"},
						&Jump{51},
						&StringIndex{"s", "i", "34"},
						&CastString{"34", "35"},
						&Concat{"element", "35", "element"},
						&Assign{"36", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "36", "i"},
						&Jump{26},
						&Assign{"37", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"37", "38", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"39", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"38", "39", "element"},
						&Append{"elements", "38", "elements"},
						&Return{Registers{"elements"}},
					},
					Registers: 39,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					Name:       "Split",
					UniqueName: "19",
				},
				"2": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ToUpper",
					UniqueName: "2",
				},
				"20": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Substr",
					UniqueName: "20",
				},
				"21": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/trim.ok:4:18"}, ""},
						&Assign{"offset", nil, "3"},
						&Len{"s", "4"},
						&LessThanNumber{"offset", "4", "5"},
						&JumpUnless{"5", 17},
						&StringIndex{"s", "offset", "6"},
						&CastString{"6", "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*8", Registers{"cutset", "7"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/trim.ok:5:47"}, ""},
						&EqualNumber{"9", "10", "11"},
						&JumpUnless{"11", 14},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "26", nil, nil, ""}, ""},
						&Call{"*12", Registers{"s", "offset"}, Registers{"13"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"offset", "14", "offset"},
						&Jump{2},
						&Return{Registers{"s"}},
					},
					Registers: 14,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimLeft",
					UniqueName: "21",
				},
				"22": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "21", nil, nil, ""}, ""},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*7", Registers{"6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"8"}},
					},
					Registers: 8,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimRight",
					UniqueName: "22",
				},
				"23": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "21", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "cutset"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "22", nil, nil, ""}, ""},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Trim",
					UniqueName: "23",
				},
				"24": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "4", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "prefix"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"4", 6},
						&Len{"prefix", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "26", nil, nil, ""}, ""},
						&Call{"*6", Registers{"s", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
						&Return{Registers{"s"}},
					},
					Registers: 7,
					Variables: map[string]*types.Type{
						"prefix": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimPrefix",
					UniqueName: "24",
				},
				"25": &CompiledFunc{
					Arguments: []string{"s", "suffix"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*5", Registers{"suffix"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "24", nil, nil, ""}, ""},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*9", Registers{"8"}, Registers{"10"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"10"}},
					},
					Registers: 10,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"suffix": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimSuffix",
					UniqueName: "25",
				},
				"26": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "substrFrom",
					UniqueName: "26",
				},
				"3": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "substr"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/contains.ok:3:32"}, ""},
						&NotEqualNumber{"4", "5", "6"},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "Contains",
					UniqueName: "3",
				},
				"4": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "HasPrefix",
					UniqueName: "4",
				},
				"5": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "HasSuffix",
					UniqueName: "5",
				},
				"6": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:3:34"}, ""},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*4", Registers{"s", "substr", "3"}, Registers{"5"}, &types.Type{
							Kind: 2,
						}},
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Index",
					UniqueName: "6",
				},
				"7": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:18:26"}, ""},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "9", nil, nil, ""}, ""},
						&Call{"*5", Registers{"offset", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"offset", nil, "6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:20:22"}, ""},
						&Add{"offset", "7", "8"},
						&Assign{"i", nil, "8"},
						&Len{"s", "9"},
						&Len{"substr", "10"},
						&Subtract{"9", "10", "11"},
						&LessThanEqualNumber{"i", "11", "12"},
						&JumpUnless{"12", 34},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/index.ok:21:17"}, ""},
						&Assign{"found", nil, "13"},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/index.ok:23:17"}, ""},
						&Assign{"j", nil, "14"},
						&Len{"substr", "15"},
						&LessThanNumber{"j", "15", "16"},
						&JumpUnless{"16", 29},
						&Add{"i", "j", "17"},
						&StringIndex{"s", "17", "18"},
						&StringIndex{"substr", "j", "19"},
						&NotEqual{"18", "19", "20"},
						&JumpUnless{"20", 26},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/index.ok:25:25"}, ""},
						&Assign{"found", nil, "21"},
						&Jump{29},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"j", "22", "j"},
						&Jump{16},
						&JumpUnless{"found", 31},
						&Return{Registers{"i"}},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "23", "i"},
						&Jump{9},
						&Assign{"24", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:35:12"}, ""},
						&Return{Registers{"24"}},
					},
					Registers: 24,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "IndexAfter",
					UniqueName: "7",
				},
				"8": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "min",
					UniqueName: "8",
				},
				"9": &CompiledFunc{
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "max",
					UniqueName: "9",
				},
				"Contains": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "substr"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/contains.ok:3:32"}, ""},
						&NotEqualNumber{"4", "5", "6"},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"substr": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "Contains",
					UniqueName: "3",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "HasPrefix",
					UniqueName: "4",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "HasSuffix",
					UniqueName: "5",
				},
				"Index": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:3:34"}, ""},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*4", Registers{"s", "substr", "3"}, Registers{"5"}, &types.Type{
							Kind: 2,
						}},
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Index",
					UniqueName: "6",
				},
				"IndexAfter": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:18:26"}, ""},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "9", nil, nil, ""}, ""},
						&Call{"*5", Registers{"offset", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"offset", nil, "6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:20:22"}, ""},
						&Add{"offset", "7", "8"},
						&Assign{"i", nil, "8"},
						&Len{"s", "9"},
						&Len{"substr", "10"},
						&Subtract{"9", "10", "11"},
						&LessThanEqualNumber{"i", "11", "12"},
						&JumpUnless{"12", 34},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/index.ok:21:17"}, ""},
						&Assign{"found", nil, "13"},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/index.ok:23:17"}, ""},
						&Assign{"j", nil, "14"},
						&Len{"substr", "15"},
						&LessThanNumber{"j", "15", "16"},
						&JumpUnless{"16", 29},
						&Add{"i", "j", "17"},
						&StringIndex{"s", "17", "18"},
						&StringIndex{"substr", "j", "19"},
						&NotEqual{"18", "19", "20"},
						&JumpUnless{"20", 26},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/index.ok:25:25"}, ""},
						&Assign{"found", nil, "21"},
						&Jump{29},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"j", "22", "j"},
						&Jump{16},
						&JumpUnless{"found", 31},
						&Return{Registers{"i"}},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "23", "i"},
						&Jump{9},
						&Assign{"24", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:35:12"}, ""},
						&Return{Registers{"24"}},
					},
					Registers: 24,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "IndexAfter",
					UniqueName: "7",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Join",
					UniqueName: "12",
				},
				"LastIndex": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*5", Registers{"substr"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:59:17"}, ""},
						&EqualNumber{"index", "9", "10"},
						&JumpUnless{"10", 11},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:60:16"}, ""},
						&Return{Registers{"11"}},
						&Len{"s", "12"},
						&Len{"substr", "13"},
						&Add{"index", "13", "14"},
						&Subtract{"12", "14", "15"},
						&Return{Registers{"15"}},
					},
					Registers: 15,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LastIndex",
					UniqueName: "10",
				},
				"LastIndexBefore": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&Len{"s", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "8", nil, nil, ""}, ""},
						&Call{"*6", Registers{"offset", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:79:45"}, ""},
						&Add{"7", "8", "9"},
						&Subtract{"4", "9", "10"},
						&Assign{"offset", nil, "10"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*11", Registers{"s"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*13", Registers{"substr"}, Registers{"14"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*15", Registers{"12", "14", "offset"}, Registers{"16"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "16"},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:82:17"}, ""},
						&EqualNumber{"index", "17", "18"},
						&JumpUnless{"18", 19},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:83:16"}, ""},
						&Return{Registers{"19"}},
						&Len{"s", "20"},
						&Len{"substr", "21"},
						&Add{"index", "21", "22"},
						&Subtract{"20", "22", "23"},
						&Return{Registers{"23"}},
					},
					Registers: 23,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "LastIndexBefore",
					UniqueName: "11",
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
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "15", nil, nil, ""}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"12", "s", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "PadLeft",
					UniqueName: "13",
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
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "15", nil, nil, ""}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"s", "12", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "PadRight",
					UniqueName: "14",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Repeat",
					UniqueName: "16",
				},
				"ReplaceAll": &CompiledFunc{
					Arguments: []string{"s", "find", "replace"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Kind: 7,
									},
								},
							},
						}, "19", nil, nil, ""}, ""},
						&Call{"*4", Registers{"s", "find"}, Registers{"5"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Kind: 7,
									},
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "12", nil, nil, ""}, ""},
						&Call{"*6", Registers{"5", "replace"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
					},
					Registers: 7,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ReplaceAll",
					UniqueName: "17",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Reverse",
					UniqueName: "18",
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
						&Jump{59},
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
						&JumpUnless{"19", 54},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:19:45"}, ""},
						&Subtract{"i", "20", "21"},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "7", nil, nil, ""}, ""},
						&Call{"*22", Registers{"s", "delimiter", "21"}, Registers{"23"}, &types.Type{
							Kind: 2,
						}},
						&Subtract{"23", "i", "24"},
						&Assign{"25", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:19:55"}, ""},
						&EqualNumber{"24", "25", "26"},
						&JumpUnless{"26", 48},
						&Assign{"27", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"27", "28", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"29", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"28", "29", "element"},
						&Append{"elements", "28", "elements"},
						&Assign{"30", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:21:27"}, ""},
						&Assign{"element", nil, "30"},
						&Len{"delimiter", "31"},
						&Assign{"32", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:22:39"}, ""},
						&Subtract{"31", "32", "33"},
						&Add{"i", "33", "i"},
						&Jump{51},
						&StringIndex{"s", "i", "34"},
						&CastString{"34", "35"},
						&Concat{"element", "35", "element"},
						&Assign{"36", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"i", "36", "i"},
						&Jump{26},
						&Assign{"37", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&ArrayAlloc{"37", "38", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"39", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, ""}, ""},
						&ArraySet{"38", "39", "element"},
						&Append{"elements", "38", "elements"},
						&Return{Registers{"elements"}},
					},
					Registers: 39,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					Name:       "Split",
					UniqueName: "19",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Substr",
					UniqueName: "20",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ToLower",
					UniqueName: "1",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ToUpper",
					UniqueName: "2",
				},
				"Trim": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "21", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "cutset"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "22", nil, nil, ""}, ""},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Trim",
					UniqueName: "23",
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
						&JumpUnless{"5", 17},
						&StringIndex{"s", "offset", "6"},
						&CastString{"6", "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "6", nil, nil, ""}, ""},
						&Call{"*8", Registers{"cutset", "7"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/trim.ok:5:47"}, ""},
						&EqualNumber{"9", "10", "11"},
						&JumpUnless{"11", 14},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "26", nil, nil, ""}, ""},
						&Call{"*12", Registers{"s", "offset"}, Registers{"13"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, ""}, ""},
						&Add{"offset", "14", "offset"},
						&Jump{2},
						&Return{Registers{"s"}},
					},
					Registers: 14,
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimLeft",
					UniqueName: "21",
				},
				"TrimPrefix": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						}, "4", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s", "prefix"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"4", 6},
						&Len{"prefix", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "26", nil, nil, ""}, ""},
						&Call{"*6", Registers{"s", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
						&Return{Registers{"s"}},
					},
					Registers: 7,
					Variables: map[string]*types.Type{
						"prefix": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimPrefix",
					UniqueName: "24",
				},
				"TrimRight": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "21", nil, nil, ""}, ""},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*7", Registers{"6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"8"}},
					},
					Registers: 8,
					Variables: map[string]*types.Type{
						"cutset": &types.Type{
							Kind: 7,
						},
						"s": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimRight",
					UniqueName: "22",
				},
				"TrimSuffix": &CompiledFunc{
					Arguments: []string{"s", "suffix"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*3", Registers{"s"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*5", Registers{"suffix"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "24", nil, nil, ""}, ""},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "18", nil, nil, ""}, ""},
						&Call{"*9", Registers{"8"}, Registers{"10"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"10"}},
					},
					Registers: 10,
					Variables: map[string]*types.Type{
						"s": &types.Type{
							Kind: 7,
						},
						"suffix": &types.Type{
							Kind: 7,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TrimSuffix",
					UniqueName: "25",
				},
				"createPad": &CompiledFunc{
					Arguments: []string{"pad", "toLen"},
					Instructions: []Instruction{
						&Len{"pad", "3"},
						&Divide{"toLen", "3", "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "16", nil, nil, ""}, ""},
						&Call{"*5", Registers{"pad", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/pad.ok:28:50"}, ""},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "20", nil, nil, ""}, ""},
						&Call{"*8", Registers{"6", "7", "toLen"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Variables: map[string]*types.Type{
						"pad": &types.Type{
							Kind: 7,
						},
						"toLen": &types.Type{
							Kind: 6,
						},
					},
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "createPad",
					UniqueName: "15",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "max",
					UniqueName: "9",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "min",
					UniqueName: "8",
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
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "substrFrom",
					UniqueName: "26",
				},
			},
			Constants: nil,
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"1"},
					&Assign{"1", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "10", nil, nil, ""}, ""},
					&ParentScope{"2"},
					&Assign{"10", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "11", nil, nil, ""}, ""},
					&ParentScope{"3"},
					&Assign{"11", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "12", nil, nil, ""}, ""},
					&ParentScope{"4"},
					&Assign{"12", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "13", nil, nil, ""}, ""},
					&ParentScope{"5"},
					&Assign{"13", nil, "5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "14", nil, nil, ""}, ""},
					&ParentScope{"6"},
					&Assign{"14", nil, "6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "15", nil, nil, ""}, ""},
					&ParentScope{"7"},
					&Assign{"15", nil, "7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "16", nil, nil, ""}, ""},
					&ParentScope{"8"},
					&Assign{"16", nil, "8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "17", nil, nil, ""}, ""},
					&ParentScope{"9"},
					&Assign{"17", nil, "9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "18", nil, nil, ""}, ""},
					&ParentScope{"10"},
					&Assign{"18", nil, "10"},
					&Assign{"11", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					}, "19", nil, nil, ""}, ""},
					&ParentScope{"11"},
					&Assign{"19", nil, "11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"12"},
					&Assign{"2", nil, "12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "20", nil, nil, ""}, ""},
					&ParentScope{"13"},
					&Assign{"20", nil, "13"},
					&Assign{"14", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "21", nil, nil, ""}, ""},
					&ParentScope{"14"},
					&Assign{"21", nil, "14"},
					&Assign{"15", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "22", nil, nil, ""}, ""},
					&ParentScope{"15"},
					&Assign{"22", nil, "15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "23", nil, nil, ""}, ""},
					&ParentScope{"16"},
					&Assign{"23", nil, "16"},
					&Assign{"17", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "24", nil, nil, ""}, ""},
					&ParentScope{"17"},
					&Assign{"24", nil, "17"},
					&Assign{"18", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "25", nil, nil, ""}, ""},
					&ParentScope{"18"},
					&Assign{"25", nil, "18"},
					&Assign{"19", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "26", nil, nil, ""}, ""},
					&ParentScope{"19"},
					&Assign{"26", nil, "19"},
					&Assign{"20", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"20"},
					&Assign{"3", nil, "20"},
					&Assign{"21", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"21"},
					&Assign{"4", nil, "21"},
					&Assign{"22", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"22"},
					&Assign{"5", nil, "22"},
					&Assign{"23", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"23"},
					&Assign{"6", nil, "23"},
					&Assign{"24", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"24"},
					&Assign{"7", nil, "24"},
					&Assign{"25", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"25"},
					&Assign{"8", nil, "25"},
					&Assign{"26", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"26"},
					&Assign{"9", nil, "26"},
					&Assign{"27", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3", nil, nil, ""}, ""},
					&ParentScope{"27"},
					&Assign{"Contains", nil, "27"},
					&Assign{"28", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "4", nil, nil, ""}, ""},
					&ParentScope{"28"},
					&Assign{"HasPrefix", nil, "28"},
					&Assign{"29", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "5", nil, nil, ""}, ""},
					&ParentScope{"29"},
					&Assign{"HasSuffix", nil, "29"},
					&Assign{"30", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "6", nil, nil, ""}, ""},
					&ParentScope{"30"},
					&Assign{"Index", nil, "30"},
					&Assign{"31", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "7", nil, nil, ""}, ""},
					&ParentScope{"31"},
					&Assign{"IndexAfter", nil, "31"},
					&Assign{"32", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "12", nil, nil, ""}, ""},
					&ParentScope{"32"},
					&Assign{"Join", nil, "32"},
					&Assign{"33", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "10", nil, nil, ""}, ""},
					&ParentScope{"33"},
					&Assign{"LastIndex", nil, "33"},
					&Assign{"34", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "11", nil, nil, ""}, ""},
					&ParentScope{"34"},
					&Assign{"LastIndexBefore", nil, "34"},
					&Assign{"35", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "13", nil, nil, ""}, ""},
					&ParentScope{"35"},
					&Assign{"PadLeft", nil, "35"},
					&Assign{"36", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "14", nil, nil, ""}, ""},
					&ParentScope{"36"},
					&Assign{"PadRight", nil, "36"},
					&Assign{"37", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "16", nil, nil, ""}, ""},
					&ParentScope{"37"},
					&Assign{"Repeat", nil, "37"},
					&Assign{"38", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "17", nil, nil, ""}, ""},
					&ParentScope{"38"},
					&Assign{"ReplaceAll", nil, "38"},
					&Assign{"39", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "18", nil, nil, ""}, ""},
					&ParentScope{"39"},
					&Assign{"Reverse", nil, "39"},
					&Assign{"40", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					}, "19", nil, nil, ""}, ""},
					&ParentScope{"40"},
					&Assign{"Split", nil, "40"},
					&Assign{"41", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "20", nil, nil, ""}, ""},
					&ParentScope{"41"},
					&Assign{"Substr", nil, "41"},
					&Assign{"42", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "1", nil, nil, ""}, ""},
					&ParentScope{"42"},
					&Assign{"ToLower", nil, "42"},
					&Assign{"43", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "2", nil, nil, ""}, ""},
					&ParentScope{"43"},
					&Assign{"ToUpper", nil, "43"},
					&Assign{"44", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "23", nil, nil, ""}, ""},
					&ParentScope{"44"},
					&Assign{"Trim", nil, "44"},
					&Assign{"45", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "21", nil, nil, ""}, ""},
					&ParentScope{"45"},
					&Assign{"TrimLeft", nil, "45"},
					&Assign{"46", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "24", nil, nil, ""}, ""},
					&ParentScope{"46"},
					&Assign{"TrimPrefix", nil, "46"},
					&Assign{"47", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "22", nil, nil, ""}, ""},
					&ParentScope{"47"},
					&Assign{"TrimRight", nil, "47"},
					&Assign{"48", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "25", nil, nil, ""}, ""},
					&ParentScope{"48"},
					&Assign{"TrimSuffix", nil, "48"},
					&Assign{"49", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "15", nil, nil, ""}, ""},
					&ParentScope{"49"},
					&Assign{"createPad", nil, "49"},
					&Assign{"50", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "9", nil, nil, ""}, ""},
					&ParentScope{"50"},
					&Assign{"max", nil, "50"},
					&Assign{"51", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "8", nil, nil, ""}, ""},
					&ParentScope{"51"},
					&Assign{"min", nil, "51"},
					&Assign{"52", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "26", nil, nil, ""}, ""},
					&ParentScope{"52"},
					&Assign{"substrFrom", nil, "52"},
					&Return{Registers{"0"}},
				},
				Registers: 52,
				Variables: map[string]*types.Type{
					"1": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"10": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"11": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"12": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"13": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"14": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"15": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"16": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"17": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"18": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"19": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					"2": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"20": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"21": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"22": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"23": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"24": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"25": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"26": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"3": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"4": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"5": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"6": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"7": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"8": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"9": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Contains": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"HasPrefix": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"HasSuffix": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					"Index": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"IndexAfter": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"Join": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"LastIndex": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"LastIndexBefore": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"PadLeft": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"PadRight": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Repeat": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"ReplaceAll": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Reverse": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Split": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Kind: 7,
								},
							},
						},
					},
					"Substr": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"ToLower": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"ToUpper": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"Trim": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"TrimLeft": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"TrimPrefix": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"TrimRight": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"TrimSuffix": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"createPad": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					"max": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"min": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					"substrFrom": &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
				},
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__strings",
							Properties: map[string]*types.Type{
								"Contains": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"HasPrefix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"HasSuffix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"Index": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"IndexAfter": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Join": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Kind: 7,
											},
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"LastIndex": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"LastIndexBefore": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"PadLeft": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"PadRight": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Repeat": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"ReplaceAll": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Reverse": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Split": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Kind: 7,
											},
										},
									},
								},
								"Substr": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"ToLower": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"ToUpper": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Trim": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"TrimLeft": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"TrimPrefix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"TrimRight": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"TrimSuffix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						},
					},
				},
				Name: "..__strings",
			},
		},
	}
}
