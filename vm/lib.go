package vm

import "github.com/elliotchance/ok/ast"
import "github.com/elliotchance/ok/types"

func init() {
	Packages = map[string]*File{
		"error": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1657625601": &CompiledFunc{
					Arguments: []string{"Error"},
					Instructions: []Instruction{
						&Return{Registers{"0"}},
					},
					Registers: 1,
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
					UniqueName: "1657625601",
					Pos:        "lib/error/error.ok:2:1",
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
					}, "1657625601", nil, nil, "", nil, nil}, ""},
					&ParentScope{"1"},
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
					}, "1657625601", nil, nil, "", nil, nil}, ""},
					&ParentScope{"2"},
					&Assign{"1657625601", nil, "1"},
					&Assign{"Error", nil, "2"},
					&Return{Registers{"0"}},
				},
				Registers: 2,
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
				Pos:  "lib/error/error.ok:2:1",
			},
		},
		"log": &File{
			Imports: map[string]*types.Type{
				"runtime": &types.Type{
					Kind: 1,
					Name: "..__runtime",
					Properties: map[string]*types.Type{
						"Env": &types.Type{
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
						"Exit": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						},
						"LookupEnv": &types.Type{
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
								&types.Type{
									Kind: 3,
								},
							},
						},
						"SetEnv": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 7,
								},
							},
						},
						"Stack": &types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "StackTrace",
									Properties: map[string]*types.Type{
										"Elements": &types.Type{
											Kind: 8,
											Element: &types.Type{
												Name: "StackElement",
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"StackElement": &types.Type{
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
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "StackElement",
									Properties: map[string]*types.Type{
										"File": &types.Type{
											Kind: 7,
										},
										"FunctionName": &types.Type{
											Kind: 7,
										},
										"LineNumber": &types.Type{
											Kind: 6,
										},
										"LineOffset": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"StackTrace": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Name: "StackElement",
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "StackTrace",
									Properties: map[string]*types.Type{
										"Elements": &types.Type{
											Kind: 8,
											Element: &types.Type{
												Name: "StackElement",
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"UnsetEnv": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						},
					},
				},
				"strings": &types.Type{
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
				"time": &types.Type{
					Kind: 1,
					Name: "..__time",
					Properties: map[string]*types.Type{
						"Add": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"After": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"April": &types.Type{
							Kind: 6,
						},
						"August": &types.Type{
							Kind: 6,
						},
						"Before": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"December": &types.Type{
							Kind: 6,
						},
						"Duration": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Equal": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"February": &types.Type{
							Kind: 6,
						},
						"FromUnix": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"Hour": &types.Type{
							Kind: 6,
						},
						"January": &types.Type{
							Kind: 6,
						},
						"July": &types.Type{
							Kind: 6,
						},
						"June": &types.Type{
							Kind: 6,
						},
						"March": &types.Type{
							Kind: 6,
						},
						"May": &types.Type{
							Kind: 6,
						},
						"Microsecond": &types.Type{
							Kind: 6,
						},
						"Millisecond": &types.Type{
							Kind: 6,
						},
						"Minute": &types.Type{
							Kind: 6,
						},
						"Nanosecond": &types.Type{
							Kind: 6,
						},
						"November": &types.Type{
							Kind: 6,
						},
						"Now": &types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"October": &types.Type{
							Kind: 6,
						},
						"Second": &types.Type{
							Kind: 6,
						},
						"September": &types.Type{
							Kind: 6,
						},
						"Sleep": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Sub": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Time": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
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
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"Unix": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						},
					},
				},
			},
			Funcs: map[string]*CompiledFunc{
				"3624248081": &CompiledFunc{
					Arguments: []string{"Log"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3624248086", nil, nil, "", nil, nil}, ""},
						&ParentScope{"2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3624248085", nil, nil, "", nil, nil}, ""},
						&ParentScope{"3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3624248084", nil, nil, "", nil, nil}, ""},
						&ParentScope{"4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3624248083", nil, nil, "", nil, nil}, ""},
						&ParentScope{"5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3624248082", nil, nil, "", nil, nil}, ""},
						&ParentScope{"6"},
						&Assign{"Fatal", nil, "2"},
						&Assign{"Error", nil, "3"},
						&Assign{"Warn", nil, "4"},
						&Assign{"Info", nil, "5"},
						&Assign{"Debug", nil, "6"},
						&Return{Registers{"0"}},
					},
					Registers: 6,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 10,
								Arguments: []*types.Type{
									&types.Type{
										Kind: 7,
									},
									&types.Type{
										Kind: 7,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Logger",
								Properties: map[string]*types.Type{
									"Debug": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
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
									},
									"Fatal": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Info": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Log": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Warn": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					},
					Name:       "Logger",
					UniqueName: "3624248081",
					Pos:        "lib/log/log.ok:19:1",
				},
				"3624248082": &CompiledFunc{
					Arguments: []string{"message"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "DEBUG", nil, nil, "", nil, nil}, ""},
						&Call{"*^Log", Registers{"2", "message"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Debug",
					UniqueName: "3624248082",
					Pos:        "lib/log/log.ok:23:5",
				},
				"3624248083": &CompiledFunc{
					Arguments: []string{"message"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "INFO", nil, nil, "", nil, nil}, ""},
						&Call{"*^Log", Registers{"2", "message"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Info",
					UniqueName: "3624248083",
					Pos:        "lib/log/log.ok:29:5",
				},
				"3624248084": &CompiledFunc{
					Arguments: []string{"message"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "WARN", nil, nil, "", nil, nil}, ""},
						&Call{"*^Log", Registers{"2", "message"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Warn",
					UniqueName: "3624248084",
					Pos:        "lib/log/log.ok:36:5",
				},
				"3624248085": &CompiledFunc{
					Arguments: []string{"message"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "ERROR", nil, nil, "", nil, nil}, ""},
						&Call{"*^Log", Registers{"2", "message"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Error",
					UniqueName: "3624248085",
					Pos:        "lib/log/log.ok:42:5",
				},
				"3624248086": &CompiledFunc{
					Arguments: []string{"message"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "FATAL", nil, nil, "", nil, nil}, ""},
						&Call{"*^Log", Registers{"2", "message"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Fatal",
					UniqueName: "3624248086",
					Pos:        "lib/log/log.ok:49:5",
				},
				"3624248087": &CompiledFunc{
					Arguments: []string{"level", "message"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/log/log.ok:58:36", nil, nil}, ""},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "5", nil, nil, "lib/log/log.ok:58:41", nil, nil}, ""},
						&LoadPackage{"5", "strings"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "PadLeft", nil, nil, "", nil, nil}, ""},
						&MapGet{"5", "6", "7"},
						&Call{"*7", Registers{"level", "3", "4"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"level", nil, "8"},
						&LoadPackage{"10", "time"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Now", nil, nil, "", nil, nil}, ""},
						&MapGet{"10", "11", "12"},
						&Call{"*12", nil, Registers{"13"}, &types.Type{
							Kind: 1,
							Name: "Time",
							Properties: map[string]*types.Type{
								"Day": &types.Type{
									Kind: 6,
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Month": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Year": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 7,
						}, "String", nil, nil, "", nil, nil}, ""},
						&MapGet{"13", "14", "15"},
						&Call{"*15", nil, Registers{"16"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/log/log.ok:59:33", nil, nil}, ""},
						&Assign{"18", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/log/log.ok:59:41", nil, nil}, ""},
						&Interpolate{"9", Registers{"16", "17", "level", "18", "message"}},
						&Print{Registers{"9"}},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 7,
						}, "FATAL", nil, nil, "", nil, nil}, ""},
						&Equal{"level", "19", "20"},
						&JumpUnless{"20", 25},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/log/log.ok:62:22", nil, nil}, ""},
						&LoadPackage{"22", "runtime"},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Exit", nil, nil, "", nil, nil}, ""},
						&MapGet{"22", "23", "24"},
						&Call{"*24", Registers{"21"}, nil, &types.Type{
							Kind: 2,
						}},
					},
					Registers: 24,
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
					},
					Name:       "Log",
					UniqueName: "3624248087",
					Pos:        "lib/log/log.ok:56:1",
				},
			},
			Constants: map[string]*ast.Literal{
				"LevelDebug": &ast.Literal{&types.Type{
					Kind: 7,
				}, "DEBUG", nil, nil, "lib/log/log.ok:7:14", nil, nil},
				"LevelError": &ast.Literal{&types.Type{
					Kind: 7,
				}, "ERROR", nil, nil, "lib/log/log.ok:10:14", nil, nil},
				"LevelFatal": &ast.Literal{&types.Type{
					Kind: 7,
				}, "FATAL", nil, nil, "lib/log/log.ok:11:14", nil, nil},
				"LevelInfo": &ast.Literal{&types.Type{
					Kind: 7,
				}, "INFO", nil, nil, "lib/log/log.ok:8:13", nil, nil},
				"LevelWarn": &ast.Literal{&types.Type{
					Kind: 7,
				}, "WARN", nil, nil, "lib/log/log.ok:9:13", nil, nil},
			},
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 10,
								Arguments: []*types.Type{
									&types.Type{
										Kind: 7,
									},
									&types.Type{
										Kind: 7,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Logger",
								Properties: map[string]*types.Type{
									"Debug": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
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
									},
									"Fatal": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Info": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Log": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Warn": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3624248081", nil, nil, "", nil, nil}, ""},
					&ParentScope{"6"},
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
					}, "3624248087", nil, nil, "", nil, nil}, ""},
					&ParentScope{"7"},
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
					}, "3624248087", nil, nil, "", nil, nil}, ""},
					&ParentScope{"8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 10,
								Arguments: []*types.Type{
									&types.Type{
										Kind: 7,
									},
									&types.Type{
										Kind: 7,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Logger",
								Properties: map[string]*types.Type{
									"Debug": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
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
									},
									"Fatal": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Info": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Log": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Warn": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3624248081", nil, nil, "", nil, nil}, ""},
					&ParentScope{"9"},
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 7,
					}, "DEBUG", nil, nil, "lib/log/log.ok:7:14", nil, nil}, ""},
					&Assign{"LevelDebug", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 7,
					}, "ERROR", nil, nil, "lib/log/log.ok:10:14", nil, nil}, ""},
					&Assign{"LevelError", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 7,
					}, "FATAL", nil, nil, "lib/log/log.ok:11:14", nil, nil}, ""},
					&Assign{"LevelFatal", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 7,
					}, "INFO", nil, nil, "lib/log/log.ok:8:13", nil, nil}, ""},
					&Assign{"LevelInfo", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 7,
					}, "WARN", nil, nil, "lib/log/log.ok:9:13", nil, nil}, ""},
					&Assign{"LevelWarn", nil, "5"},
					&Assign{"3624248081", nil, "6"},
					&Assign{"3624248087", nil, "7"},
					&Assign{"Log", nil, "8"},
					&Assign{"Logger", nil, "9"},
					&Return{Registers{"0"}},
				},
				Registers: 9,
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__log",
							Properties: map[string]*types.Type{
								"LevelDebug": &types.Type{
									Kind: 7,
								},
								"LevelError": &types.Type{
									Kind: 7,
								},
								"LevelFatal": &types.Type{
									Kind: 7,
								},
								"LevelInfo": &types.Type{
									Kind: 7,
								},
								"LevelWarn": &types.Type{
									Kind: 7,
								},
								"Log": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Logger": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 10,
											Arguments: []*types.Type{
												&types.Type{
													Kind: 7,
												},
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Logger",
											Properties: map[string]*types.Type{
												"Debug": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
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
												},
												"Fatal": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Info": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Log": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
														},
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Warn": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
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
					},
				},
				Name: "..__log",
				Pos:  "lib/log/log.ok:1:1",
			},
		},
		"math": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1725484959": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/abs.ok:3:12", nil, nil}, ""},
						&LessThanNumber{"x", "2", "3"},
						&JumpUnless{"3", 5},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&Subtract{"4", "x", "5"},
						&Return{Registers{"5"}},
						&Return{Registers{"x"}},
					},
					Registers: 5,
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
					UniqueName: "1725484959",
					Pos:        "lib/math/abs.ok:2:1",
				},
				"1725484960": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Log{"x", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
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
					UniqueName: "1725484960",
					Pos:        "lib/math/log.ok:2:1",
				},
				"1725484961": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Log{"x", "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "10", nil, nil, "lib/math/log.ok:8:29", nil, nil}, ""},
						&Log{"3", "4"},
						&Divide{"2", "4", "5"},
						&Return{Registers{"5"}},
					},
					Registers: 5,
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
					UniqueName: "1725484961",
					Pos:        "lib/math/log.ok:7:1",
				},
				"1725484962": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "2.71828182845904523536028747135266249775724709369995", nil, nil, "lib/math/powers.ok:4:9", nil, nil}, ""},
						&Assign{"e", nil, "2"},
						&Power{"e", "x", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
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
					UniqueName: "1725484962",
					Pos:        "lib/math/powers.ok:2:1",
				},
				"1725484963": &CompiledFunc{
					Arguments: []string{"base", "power"},
					Instructions: []Instruction{
						&Power{"base", "power", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
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
					UniqueName: "1725484963",
					Pos:        "lib/math/powers.ok:10:1",
				},
				"1725484964": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.5", nil, nil, "lib/math/powers.ok:16:21", nil, nil}, ""},
						&Power{"x", "2", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
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
					UniqueName: "1725484964",
					Pos:        "lib/math/powers.ok:15:1",
				},
				"1725484965": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/powers.ok:21:21", nil, nil}, ""},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3", nil, nil, "lib/math/powers.ok:21:23", nil, nil}, ""},
						&Divide{"2", "3", "4"},
						&Power{"x", "4", "5"},
						&Return{Registers{"5"}},
					},
					Registers: 5,
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
					UniqueName: "1725484965",
					Pos:        "lib/math/powers.ok:20:1",
				},
				"1725484966": &CompiledFunc{
					Instructions: []Instruction{
						&Rand{"1"},
						&Return{Registers{"1"}},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Rand",
					UniqueName: "1725484966",
					Pos:        "lib/math/rand.ok:2:1",
				},
				"1725484967": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:3:16", nil, nil}, ""},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:4:16", nil, nil}, ""},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:8:12", nil, nil}, ""},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 11},
						&Subtract{"x", "frac", "8"},
						&Return{Registers{"8"}},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:12:17", nil, nil}, ""},
						&Subtract{"9", "frac", "10"},
						&Add{"x", "10", "11"},
						&Return{Registers{"11"}},
					},
					Registers: 11,
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
					UniqueName: "1725484967",
					Pos:        "lib/math/rounding.ok:2:1",
				},
				"1725484968": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:17:16", nil, nil}, ""},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:18:16", nil, nil}, ""},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/math/rounding.ok:22:12", nil, nil}, ""},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 13},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:23:27", nil, nil}, ""},
						&Add{"frac", "8", "9"},
						&Subtract{"x", "9", "10"},
						&Return{Registers{"10"}},
						&Subtract{"x", "frac", "11"},
						&Return{Registers{"11"}},
					},
					Registers: 11,
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
					UniqueName: "1725484968",
					Pos:        "lib/math/rounding.ok:16:1",
				},
				"1725484969": &CompiledFunc{
					Arguments: []string{"x", "prec"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "10", nil, nil, "lib/math/rounding.ok:32:15", nil, nil}, ""},
						&Power{"3", "prec", "4"},
						&Assign{"p", nil, "4"},
						&Multiply{"x", "p", "5"},
						&Assign{"y", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:35:16", nil, nil}, ""},
						&Remainder{"y", "6", "7"},
						&Assign{"diff", nil, "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.5", nil, nil, "lib/math/rounding.ok:36:16", nil, nil}, ""},
						&GreaterThanEqualNumber{"diff", "8", "9"},
						&JumpUnless{"9", 15},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/math/rounding.ok:37:22", nil, nil}, ""},
						&Subtract{"10", "diff", "11"},
						&Add{"y", "11", "12"},
						&Divide{"12", "p", "13"},
						&Return{Registers{"13"}},
						&Subtract{"y", "diff", "14"},
						&Divide{"14", "p", "15"},
						&Return{Registers{"15"}},
					},
					Registers: 15,
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
					UniqueName: "1725484969",
					Pos:        "lib/math/rounding.ok:31:1",
				},
			},
			Constants: map[string]*ast.Literal{
				"E": &ast.Literal{&types.Type{
					Kind: 6,
				}, "2.71828182845904523536028747135266249775724709369995957496696763", nil, nil, "lib/math/constants.ok:1:7", nil, nil},
				"Ln10": &ast.Literal{&types.Type{
					Kind: 6,
				}, "2.30258509299404568401799145468436420760110148862877297603332790", nil, nil, "lib/math/constants.ok:11:8", nil, nil},
				"Ln2": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.693147180559945309417232121458176568075500134360255254120680009", nil, nil, "lib/math/constants.ok:10:8", nil, nil},
				"Phi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.61803398874989484820458683436563811772030917980576286213544862", nil, nil, "lib/math/constants.ok:3:7", nil, nil},
				"Pi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "3.14159265358979323846264338327950288419716939937510582097494459", nil, nil, "lib/math/constants.ok:2:7", nil, nil},
				"Sqrt2": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.41421356237309504880168872420969807856967187537694807317667974", nil, nil, "lib/math/constants.ok:5:11", nil, nil},
				"SqrtE": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.64872127070012814684865078781416357165377610071014801157507931", nil, nil, "lib/math/constants.ok:6:11", nil, nil},
				"SqrtPhi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.27201964951406896425242246173749149171560804184009624861664038", nil, nil, "lib/math/constants.ok:8:11", nil, nil},
				"SqrtPi": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1.77245385090551602729816748334114518279754945612238712821380779", nil, nil, "lib/math/constants.ok:7:11", nil, nil},
			},
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
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
					}, "1725484959", nil, nil, "", nil, nil}, ""},
					&ParentScope{"10"},
					&Assign{"11", &ast.Literal{&types.Type{
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
					}, "1725484960", nil, nil, "", nil, nil}, ""},
					&ParentScope{"11"},
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
					}, "1725484961", nil, nil, "", nil, nil}, ""},
					&ParentScope{"12"},
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
					}, "1725484962", nil, nil, "", nil, nil}, ""},
					&ParentScope{"13"},
					&Assign{"14", &ast.Literal{&types.Type{
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
					}, "1725484963", nil, nil, "", nil, nil}, ""},
					&ParentScope{"14"},
					&Assign{"15", &ast.Literal{&types.Type{
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
					}, "1725484964", nil, nil, "", nil, nil}, ""},
					&ParentScope{"15"},
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
					}, "1725484965", nil, nil, "", nil, nil}, ""},
					&ParentScope{"16"},
					&Assign{"17", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "1725484966", nil, nil, "", nil, nil}, ""},
					&ParentScope{"17"},
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
					}, "1725484967", nil, nil, "", nil, nil}, ""},
					&ParentScope{"18"},
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
					}, "1725484968", nil, nil, "", nil, nil}, ""},
					&ParentScope{"19"},
					&Assign{"20", &ast.Literal{&types.Type{
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
					}, "1725484969", nil, nil, "", nil, nil}, ""},
					&ParentScope{"20"},
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
					}, "1725484959", nil, nil, "", nil, nil}, ""},
					&ParentScope{"21"},
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
					}, "1725484965", nil, nil, "", nil, nil}, ""},
					&ParentScope{"22"},
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
					}, "1725484967", nil, nil, "", nil, nil}, ""},
					&ParentScope{"23"},
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
					}, "1725484962", nil, nil, "", nil, nil}, ""},
					&ParentScope{"24"},
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
					}, "1725484968", nil, nil, "", nil, nil}, ""},
					&ParentScope{"25"},
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
					}, "1725484961", nil, nil, "", nil, nil}, ""},
					&ParentScope{"26"},
					&Assign{"27", &ast.Literal{&types.Type{
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
					}, "1725484960", nil, nil, "", nil, nil}, ""},
					&ParentScope{"27"},
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
					}, "1725484963", nil, nil, "", nil, nil}, ""},
					&ParentScope{"28"},
					&Assign{"29", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "1725484966", nil, nil, "", nil, nil}, ""},
					&ParentScope{"29"},
					&Assign{"30", &ast.Literal{&types.Type{
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
					}, "1725484969", nil, nil, "", nil, nil}, ""},
					&ParentScope{"30"},
					&Assign{"31", &ast.Literal{&types.Type{
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
					}, "1725484964", nil, nil, "", nil, nil}, ""},
					&ParentScope{"31"},
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 6,
					}, "2.71828182845904523536028747135266249775724709369995957496696763", nil, nil, "lib/math/constants.ok:1:7", nil, nil}, ""},
					&Assign{"E", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 6,
					}, "2.30258509299404568401799145468436420760110148862877297603332790", nil, nil, "lib/math/constants.ok:11:8", nil, nil}, ""},
					&Assign{"Ln10", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 6,
					}, "0.693147180559945309417232121458176568075500134360255254120680009", nil, nil, "lib/math/constants.ok:10:8", nil, nil}, ""},
					&Assign{"Ln2", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.61803398874989484820458683436563811772030917980576286213544862", nil, nil, "lib/math/constants.ok:3:7", nil, nil}, ""},
					&Assign{"Phi", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 6,
					}, "3.14159265358979323846264338327950288419716939937510582097494459", nil, nil, "lib/math/constants.ok:2:7", nil, nil}, ""},
					&Assign{"Pi", nil, "5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.41421356237309504880168872420969807856967187537694807317667974", nil, nil, "lib/math/constants.ok:5:11", nil, nil}, ""},
					&Assign{"Sqrt2", nil, "6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.64872127070012814684865078781416357165377610071014801157507931", nil, nil, "lib/math/constants.ok:6:11", nil, nil}, ""},
					&Assign{"SqrtE", nil, "7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.27201964951406896425242246173749149171560804184009624861664038", nil, nil, "lib/math/constants.ok:8:11", nil, nil}, ""},
					&Assign{"SqrtPhi", nil, "8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1.77245385090551602729816748334114518279754945612238712821380779", nil, nil, "lib/math/constants.ok:7:11", nil, nil}, ""},
					&Assign{"SqrtPi", nil, "9"},
					&Assign{"1725484959", nil, "10"},
					&Assign{"1725484960", nil, "11"},
					&Assign{"1725484961", nil, "12"},
					&Assign{"1725484962", nil, "13"},
					&Assign{"1725484963", nil, "14"},
					&Assign{"1725484964", nil, "15"},
					&Assign{"1725484965", nil, "16"},
					&Assign{"1725484966", nil, "17"},
					&Assign{"1725484967", nil, "18"},
					&Assign{"1725484968", nil, "19"},
					&Assign{"1725484969", nil, "20"},
					&Assign{"Abs", nil, "21"},
					&Assign{"Cbrt", nil, "22"},
					&Assign{"Ceil", nil, "23"},
					&Assign{"Exp", nil, "24"},
					&Assign{"Floor", nil, "25"},
					&Assign{"Log10", nil, "26"},
					&Assign{"LogE", nil, "27"},
					&Assign{"Pow", nil, "28"},
					&Assign{"Rand", nil, "29"},
					&Assign{"Round", nil, "30"},
					&Assign{"Sqrt", nil, "31"},
					&Return{Registers{"0"}},
				},
				Registers: 31,
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
								"Rand": &types.Type{
									Kind: 10,
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
				Pos:  "lib/math/rounding.ok:2:1",
			},
		},
		"os": &File{
			Imports: map[string]*types.Type{
				"time": &types.Type{
					Kind: 1,
					Name: "..__time",
					Properties: map[string]*types.Type{
						"Add": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"After": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"April": &types.Type{
							Kind: 6,
						},
						"August": &types.Type{
							Kind: 6,
						},
						"Before": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"December": &types.Type{
							Kind: 6,
						},
						"Duration": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Equal": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 3,
								},
							},
						},
						"February": &types.Type{
							Kind: 6,
						},
						"FromUnix": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"Hour": &types.Type{
							Kind: 6,
						},
						"January": &types.Type{
							Kind: 6,
						},
						"July": &types.Type{
							Kind: 6,
						},
						"June": &types.Type{
							Kind: 6,
						},
						"March": &types.Type{
							Kind: 6,
						},
						"May": &types.Type{
							Kind: 6,
						},
						"Microsecond": &types.Type{
							Kind: 6,
						},
						"Millisecond": &types.Type{
							Kind: 6,
						},
						"Minute": &types.Type{
							Kind: 6,
						},
						"Nanosecond": &types.Type{
							Kind: 6,
						},
						"November": &types.Type{
							Kind: 6,
						},
						"Now": &types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"October": &types.Type{
							Kind: 6,
						},
						"Second": &types.Type{
							Kind: 6,
						},
						"September": &types.Type{
							Kind: 6,
						},
						"Sleep": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Sub": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
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
						"Time": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
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
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
						"Unix": &types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						},
					},
				},
			},
			Funcs: map[string]*CompiledFunc{
				"3819329014": &CompiledFunc{
					Arguments: []string{"fd"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3819329021", nil, nil, "", nil, nil}, ""},
						&ParentScope{"2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3819329020", nil, nil, "", nil, nil}, ""},
						&ParentScope{"3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 5,
								},
							},
						}, "3819329019", nil, nil, "", nil, nil}, ""},
						&ParentScope{"4"},
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
						}, "3819329018", nil, nil, "", nil, nil}, ""},
						&ParentScope{"5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
						}, "3819329017", nil, nil, "", nil, nil}, ""},
						&ParentScope{"6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 5,
								},
							},
						}, "3819329016", nil, nil, "", nil, nil}, ""},
						&ParentScope{"7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3819329015", nil, nil, "", nil, nil}, ""},
						&ParentScope{"8"},
						&Assign{"ReadLine", nil, "2"},
						&Assign{"ReadString", nil, "3"},
						&Assign{"ReadData", nil, "4"},
						&Assign{"Seek", nil, "5"},
						&Assign{"Close", nil, "6"},
						&Assign{"WriteData", nil, "7"},
						&Assign{"WriteString", nil, "8"},
						&Return{Registers{"0"}},
					},
					Registers: 8,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 5,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					},
					Name:       "File",
					UniqueName: "3819329014",
					Pos:        "lib/os/file.ok:2:1",
				},
				"3819329015": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&CastData{"s", "2"},
						&Write{"2", "^fd"},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "WriteString",
					UniqueName: "3819329015",
					Pos:        "lib/os/file.ok:4:5",
				},
				"3819329016": &CompiledFunc{
					Arguments: []string{"d"},
					Instructions: []Instruction{
						&Write{"d", "^fd"},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 5,
							},
						},
					},
					Name:       "WriteData",
					UniqueName: "3819329016",
					Pos:        "lib/os/file.ok:9:5",
				},
				"3819329017": &CompiledFunc{
					Instructions: []Instruction{
						&Close{"^fd"},
					},
					Type: &types.Type{
						Kind: 10,
					},
					Name:       "Close",
					UniqueName: "3819329017",
					Pos:        "lib/os/file.ok:15:5",
				},
				"3819329018": &CompiledFunc{
					Arguments: []string{"offset", "whence"},
					Instructions: []Instruction{
						&Seek{"^fd", "offset", "offset", "3"},
					},
					Registers: 3,
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
					Name:       "Seek",
					UniqueName: "3819329018",
					Pos:        "lib/os/file.ok:26:5",
				},
				"3819329019": &CompiledFunc{
					Arguments: []string{"bytes"},
					Instructions: []Instruction{
						&ReadData{"^fd", "bytes", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 5,
							},
						},
					},
					Name:       "ReadData",
					UniqueName: "3819329019",
					Pos:        "lib/os/file.ok:38:5",
				},
				"3819329020": &CompiledFunc{
					Arguments: []string{"chars"},
					Instructions: []Instruction{
						&ReadString{"^fd", "chars", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
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
					Name:       "ReadString",
					UniqueName: "3819329020",
					Pos:        "lib/os/file.ok:45:5",
				},
				"3819329021": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/os/file.ok:51:16", nil, nil}, ""},
						&Assign{"line", nil, "1"},
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "", nil, nil}, ""},
						&JumpUnless{"2", 17},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/os/file.ok:54:40", nil, nil}, ""},
						&ReadString{"^fd", "3", "4"},
						&Assign{"chars", nil, "4"},
						&Len{"chars", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/os/file.ok:55:30", nil, nil}, ""},
						&EqualNumber{"5", "6", "7"},
						&JumpUnless{"7", 11},
						&Jump{17},
						&Concat{"line", "chars", "line"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 7,
						}, "\n", nil, nil, "lib/os/file.ok:61:25", nil, nil}, ""},
						&Equal{"chars", "8", "9"},
						&JumpUnless{"9", 16},
						&Jump{17},
						&Jump{1},
						&Return{Registers{"line"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "ReadLine",
					UniqueName: "3819329021",
					Pos:        "lib/os/file.ok:50:5",
				},
				"3819329022": &CompiledFunc{
					Arguments: []string{"path"},
					Instructions: []Instruction{
						&Open{"path", "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 5,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "File",
									Properties: map[string]*types.Type{
										"Close": &types.Type{
											Kind: 10,
										},
										"ReadData": &types.Type{
											Kind: 10,
											Arguments: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
											Returns: []*types.Type{
												&types.Type{
													Kind: 5,
												},
											},
										},
										"ReadLine": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"ReadString": &types.Type{
											Kind: 10,
											Arguments: []*types.Type{
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
										"Seek": &types.Type{
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
										"WriteData": &types.Type{
											Kind: 10,
											Arguments: []*types.Type{
												&types.Type{
													Kind: 5,
												},
											},
										},
										"WriteString": &types.Type{
											Kind: 10,
											Arguments: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3819329014", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"2"}, Registers{"4"}, &types.Type{
							Kind: 1,
							Name: "File",
							Properties: map[string]*types.Type{
								"Close": &types.Type{
									Kind: 10,
								},
								"ReadData": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 5,
										},
									},
								},
								"ReadLine": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"ReadString": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
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
								"Seek": &types.Type{
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
								"WriteData": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 5,
										},
									},
								},
								"WriteString": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Return{Registers{"4"}},
					},
					Registers: 4,
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
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					},
					Name:       "Open",
					UniqueName: "3819329022",
					Pos:        "lib/os/file.ok:74:1",
				},
				"3819329023": &CompiledFunc{
					Arguments: []string{"path"},
					Instructions: []Instruction{
						&Remove{"path"},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "Remove",
					UniqueName: "3819329023",
					Pos:        "lib/os/filesystem.ok:2:1",
				},
				"3819329024": &CompiledFunc{
					Arguments: []string{"old", "new"},
					Instructions: []Instruction{
						&Rename{"old", "new"},
					},
					Registers: 2,
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
					},
					Name:       "Rename",
					UniqueName: "3819329024",
					Pos:        "lib/os/filesystem.ok:7:1",
				},
				"3819329025": &CompiledFunc{
					Arguments: []string{"path"},
					Instructions: []Instruction{
						&Mkdir{"path"},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "CreateDirectory",
					UniqueName: "3819329025",
					Pos:        "lib/os/filesystem.ok:13:1",
				},
				"3819329026": &CompiledFunc{
					Arguments: []string{"Name", "Size", "Mode", "ModifiedTime", "IsDir"},
					Instructions: []Instruction{
						&Return{Registers{"0"}},
					},
					Registers: 5,
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
								Kind: 7,
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 3,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Kind: 1,
										Name: "Time",
										Properties: map[string]*types.Type{
											"Day": &types.Type{
												Kind: 6,
											},
											"Hour": &types.Type{
												Kind: 6,
											},
											"Minute": &types.Type{
												Kind: 6,
											},
											"Month": &types.Type{
												Kind: 6,
											},
											"Second": &types.Type{
												Kind: 6,
											},
											"String": &types.Type{
												Kind: 10,
												Returns: []*types.Type{
													&types.Type{
														Kind: 7,
													},
												},
											},
											"Year": &types.Type{
												Kind: 6,
											},
										},
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "FileInfo",
					UniqueName: "3819329026",
					Pos:        "lib/os/info.ok:3:1",
				},
				"3819329027": &CompiledFunc{
					Arguments: []string{"path"},
					Instructions: []Instruction{
						&Info{"path", "2", "3", "4", []Register{
							"5",
							"6",
							"7",
							"8",
							"9",
							"10",
						}, "11"},
						&Assign{"name", nil, "2"},
						&Assign{"size", nil, "3"},
						&Assign{"mode", nil, "4"},
						&Assign{"modTimeYear", nil, "5"},
						&Assign{"modTimeMonth", nil, "6"},
						&Assign{"modTimeDay", nil, "7"},
						&Assign{"modTimeHour", nil, "8"},
						&Assign{"modTimeMinute", nil, "9"},
						&Assign{"modTimeSecond", nil, "10"},
						&Assign{"isDir", nil, "11"},
						&LoadPackage{"12", "time"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Time", nil, nil, "", nil, nil}, ""},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"modTimeYear", "modTimeMonth", "modTimeDay", "modTimeHour", "modTimeMinute", "modTimeSecond"}, Registers{"15"}, &types.Type{
							Kind: 1,
							Name: "Time",
							Properties: map[string]*types.Type{
								"Day": &types.Type{
									Kind: 6,
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Month": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Year": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Assign{"modTime", nil, "15"},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 7,
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 3,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "FileInfo",
									Properties: map[string]*types.Type{
										"IsDir": &types.Type{
											Kind: 3,
										},
										"Mode": &types.Type{
											Kind: 7,
										},
										"ModifiedTime": &types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										"Name": &types.Type{
											Kind: 7,
										},
										"Size": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						}, "3819329026", nil, nil, "", nil, nil}, ""},
						&Call{"*16", Registers{"name", "size", "mode", "modTime", "isDir"}, Registers{"17"}, &types.Type{
							Kind: 1,
							Name: "FileInfo",
							Properties: map[string]*types.Type{
								"IsDir": &types.Type{
									Kind: 3,
								},
								"Mode": &types.Type{
									Kind: 7,
								},
								"ModifiedTime": &types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								"Name": &types.Type{
									Kind: 7,
								},
								"Size": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Return{Registers{"17"}},
					},
					Registers: 17,
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
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Name: "time.Time",
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "Info",
					UniqueName: "3819329027",
					Pos:        "lib/os/info.ok:13:1",
				},
				"3819329028": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 7,
						}, "/tmp/ok.", nil, nil, "lib/os/temp.ok:5:12", nil, nil}, ""},
						&Rand{"2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1000000000", nil, nil, "lib/os/temp.ok:5:44", nil, nil}, ""},
						&Multiply{"2", "3", "4"},
						&CastString{"4", "5"},
						&Concat{"1", "5", "6"},
						&Return{Registers{"6"}},
					},
					Registers: 6,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "TempPath",
					UniqueName: "3819329028",
					Pos:        "lib/os/temp.ok:2:1",
				},
			},
			Constants: nil,
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 5,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3819329014", nil, nil, "", nil, nil}, ""},
					&ParentScope{"1"},
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
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3819329022", nil, nil, "", nil, nil}, ""},
					&ParentScope{"2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329023", nil, nil, "", nil, nil}, ""},
					&ParentScope{"3"},
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
					}, "3819329024", nil, nil, "", nil, nil}, ""},
					&ParentScope{"4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329025", nil, nil, "", nil, nil}, ""},
					&ParentScope{"5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 3,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Kind: 1,
										Name: "Time",
										Properties: map[string]*types.Type{
											"Day": &types.Type{
												Kind: 6,
											},
											"Hour": &types.Type{
												Kind: 6,
											},
											"Minute": &types.Type{
												Kind: 6,
											},
											"Month": &types.Type{
												Kind: 6,
											},
											"Second": &types.Type{
												Kind: 6,
											},
											"String": &types.Type{
												Kind: 10,
												Returns: []*types.Type{
													&types.Type{
														Kind: 7,
													},
												},
											},
											"Year": &types.Type{
												Kind: 6,
											},
										},
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3819329026", nil, nil, "", nil, nil}, ""},
					&ParentScope{"6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Name: "time.Time",
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3819329027", nil, nil, "", nil, nil}, ""},
					&ParentScope{"7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329028", nil, nil, "", nil, nil}, ""},
					&ParentScope{"8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329025", nil, nil, "", nil, nil}, ""},
					&ParentScope{"9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 5,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3819329014", nil, nil, "", nil, nil}, ""},
					&ParentScope{"10"},
					&Assign{"11", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 3,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Kind: 1,
										Name: "Time",
										Properties: map[string]*types.Type{
											"Day": &types.Type{
												Kind: 6,
											},
											"Hour": &types.Type{
												Kind: 6,
											},
											"Minute": &types.Type{
												Kind: 6,
											},
											"Month": &types.Type{
												Kind: 6,
											},
											"Second": &types.Type{
												Kind: 6,
											},
											"String": &types.Type{
												Kind: 10,
												Returns: []*types.Type{
													&types.Type{
														Kind: 7,
													},
												},
											},
											"Year": &types.Type{
												Kind: 6,
											},
										},
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3819329026", nil, nil, "", nil, nil}, ""},
					&ParentScope{"11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "FileInfo",
								Properties: map[string]*types.Type{
									"IsDir": &types.Type{
										Kind: 3,
									},
									"Mode": &types.Type{
										Kind: 7,
									},
									"ModifiedTime": &types.Type{
										Name: "time.Time",
									},
									"Name": &types.Type{
										Kind: 7,
									},
									"Size": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3819329027", nil, nil, "", nil, nil}, ""},
					&ParentScope{"12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "File",
								Properties: map[string]*types.Type{
									"Close": &types.Type{
										Kind: 10,
									},
									"ReadData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
										Returns: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"ReadLine": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"ReadString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
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
									"Seek": &types.Type{
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
									"WriteData": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 5,
											},
										},
									},
									"WriteString": &types.Type{
										Kind: 10,
										Arguments: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3819329022", nil, nil, "", nil, nil}, ""},
					&ParentScope{"13"},
					&Assign{"14", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329023", nil, nil, "", nil, nil}, ""},
					&ParentScope{"14"},
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
					}, "3819329024", nil, nil, "", nil, nil}, ""},
					&ParentScope{"15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3819329028", nil, nil, "", nil, nil}, ""},
					&ParentScope{"16"},
					&Assign{"3819329014", nil, "1"},
					&Assign{"3819329022", nil, "2"},
					&Assign{"3819329023", nil, "3"},
					&Assign{"3819329024", nil, "4"},
					&Assign{"3819329025", nil, "5"},
					&Assign{"3819329026", nil, "6"},
					&Assign{"3819329027", nil, "7"},
					&Assign{"3819329028", nil, "8"},
					&Assign{"CreateDirectory", nil, "9"},
					&Assign{"File", nil, "10"},
					&Assign{"FileInfo", nil, "11"},
					&Assign{"Info", nil, "12"},
					&Assign{"Open", nil, "13"},
					&Assign{"Remove", nil, "14"},
					&Assign{"Rename", nil, "15"},
					&Assign{"TempPath", nil, "16"},
					&Return{Registers{"0"}},
				},
				Registers: 16,
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__os",
							Properties: map[string]*types.Type{
								"CreateDirectory": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"File": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 5,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "File",
											Properties: map[string]*types.Type{
												"Close": &types.Type{
													Kind: 10,
												},
												"ReadData": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
													Returns: []*types.Type{
														&types.Type{
															Kind: 5,
														},
													},
												},
												"ReadLine": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"ReadString": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
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
												"Seek": &types.Type{
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
												"WriteData": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 5,
														},
													},
												},
												"WriteString": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
											},
										},
									},
								},
								"FileInfo": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 3,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "FileInfo",
											Properties: map[string]*types.Type{
												"IsDir": &types.Type{
													Kind: 3,
												},
												"Mode": &types.Type{
													Kind: 7,
												},
												"ModifiedTime": &types.Type{
													Kind: 1,
													Name: "Time",
													Properties: map[string]*types.Type{
														"Day": &types.Type{
															Kind: 6,
														},
														"Hour": &types.Type{
															Kind: 6,
														},
														"Minute": &types.Type{
															Kind: 6,
														},
														"Month": &types.Type{
															Kind: 6,
														},
														"Second": &types.Type{
															Kind: 6,
														},
														"String": &types.Type{
															Kind: 10,
															Returns: []*types.Type{
																&types.Type{
																	Kind: 7,
																},
															},
														},
														"Year": &types.Type{
															Kind: 6,
														},
													},
												},
												"Name": &types.Type{
													Kind: 7,
												},
												"Size": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"Info": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "FileInfo",
											Properties: map[string]*types.Type{
												"IsDir": &types.Type{
													Kind: 3,
												},
												"Mode": &types.Type{
													Kind: 7,
												},
												"ModifiedTime": &types.Type{
													Name: "time.Time",
												},
												"Name": &types.Type{
													Kind: 7,
												},
												"Size": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"Open": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "File",
											Properties: map[string]*types.Type{
												"Close": &types.Type{
													Kind: 10,
												},
												"ReadData": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
													Returns: []*types.Type{
														&types.Type{
															Kind: 5,
														},
													},
												},
												"ReadLine": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"ReadString": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
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
												"Seek": &types.Type{
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
												"WriteData": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 5,
														},
													},
												},
												"WriteString": &types.Type{
													Kind: 10,
													Arguments: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
											},
										},
									},
								},
								"Remove": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Rename": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
								},
								"TempPath": &types.Type{
									Kind: 10,
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
				Name: "..__os",
				Pos:  "lib/os/temp.ok:2:1",
			},
		},
		"reflect": &File{
			Imports: map[string]*types.Type{
				"strings": &types.Type{
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
			Funcs: map[string]*CompiledFunc{
				"1100080403": &CompiledFunc{
					Arguments: []string{"fn", "args"},
					Instructions: []Instruction{
						&DynamicCall{"fn", "args", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
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
					UniqueName: "1100080403",
					Pos:        "lib/reflect/call.ok:16:1",
				},
				"1100080404": &CompiledFunc{
					Arguments: []string{"obj", "prop"},
					Instructions: []Instruction{
						&Get{"obj", "prop", "3"},
						&Return{Registers{"3"}},
					},
					Registers: 3,
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
					UniqueName: "1100080404",
					Pos:        "lib/reflect/get.ok:15:1",
				},
				"1100080405": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Interface{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
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
					UniqueName: "1100080405",
					Pos:        "lib/reflect/interface.ok:10:1",
				},
				"1100080406": &CompiledFunc{
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
						}, "1100080410", nil, nil, "", nil, nil}, ""},
						&Call{"*2", Registers{"value"}, Registers{"3"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"type", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "[]", nil, nil, "lib/reflect/kind.ok:9:38", nil, nil}, ""},
						&LoadPackage{"5", "strings"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "HasPrefix", nil, nil, "", nil, nil}, ""},
						&MapGet{"5", "6", "7"},
						&Call{"*7", Registers{"type", "4"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"8", 11},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 7,
						}, "array", nil, nil, "lib/reflect/kind.ok:10:20", nil, nil}, ""},
						&Return{Registers{"9"}},
						&Jump{30},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 7,
						}, "{}", nil, nil, "lib/reflect/kind.ok:13:39", nil, nil}, ""},
						&Interpolate{"10", Registers{"11"}},
						&LoadPackage{"12", "strings"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "HasPrefix", nil, nil, "", nil, nil}, ""},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"type", "10"}, Registers{"15"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"15", 21},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 7,
						}, "map", nil, nil, "lib/reflect/kind.ok:14:20", nil, nil}, ""},
						&Return{Registers{"16"}},
						&Jump{30},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func(", nil, nil, "lib/reflect/kind.ok:17:38", nil, nil}, ""},
						&LoadPackage{"18", "strings"},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 7,
						}, "HasPrefix", nil, nil, "", nil, nil}, ""},
						&MapGet{"18", "19", "20"},
						&Call{"*20", Registers{"type", "17"}, Registers{"21"}, &types.Type{
							Kind: 2,
						}},
						&JumpUnless{"21", 30},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 7,
						}, "func", nil, nil, "lib/reflect/kind.ok:18:20", nil, nil}, ""},
						&Return{Registers{"22"}},
						&Jump{30},
						&Return{Registers{"type"}},
					},
					Registers: 22,
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
					UniqueName: "1100080406",
					Pos:        "lib/reflect/kind.ok:5:1",
				},
				"1100080407": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Len{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
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
					UniqueName: "1100080407",
					Pos:        "lib/reflect/len.ok:3:1",
				},
				"1100080408": &CompiledFunc{
					Arguments: []string{"obj"},
					Instructions: []Instruction{
						&Props{"obj", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
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
					UniqueName: "1100080408",
					Pos:        "lib/reflect/props.ok:4:1",
				},
				"1100080409": &CompiledFunc{
					Arguments: []string{"obj", "prop", "value"},
					Instructions: []Instruction{
						&Set{"obj", "prop", "value", "4"},
						&Return{Registers{"4"}},
					},
					Registers: 4,
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
					UniqueName: "1100080409",
					Pos:        "lib/reflect/set.ok:16:1",
				},
				"1100080410": &CompiledFunc{
					Arguments: []string{"value"},
					Instructions: []Instruction{
						&Type{"value", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
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
					UniqueName: "1100080410",
					Pos:        "lib/reflect/type.ok:8:1",
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
					}, "1100080403", nil, nil, "", nil, nil}, ""},
					&ParentScope{"1"},
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
					}, "1100080404", nil, nil, "", nil, nil}, ""},
					&ParentScope{"2"},
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
					}, "1100080405", nil, nil, "", nil, nil}, ""},
					&ParentScope{"3"},
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
					}, "1100080406", nil, nil, "", nil, nil}, ""},
					&ParentScope{"4"},
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
					}, "1100080407", nil, nil, "", nil, nil}, ""},
					&ParentScope{"5"},
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
					}, "1100080408", nil, nil, "", nil, nil}, ""},
					&ParentScope{"6"},
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
					}, "1100080409", nil, nil, "", nil, nil}, ""},
					&ParentScope{"7"},
					&Assign{"8", &ast.Literal{&types.Type{
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
					}, "1100080410", nil, nil, "", nil, nil}, ""},
					&ParentScope{"8"},
					&Assign{"9", &ast.Literal{&types.Type{
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
					}, "1100080403", nil, nil, "", nil, nil}, ""},
					&ParentScope{"9"},
					&Assign{"10", &ast.Literal{&types.Type{
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
					}, "1100080404", nil, nil, "", nil, nil}, ""},
					&ParentScope{"10"},
					&Assign{"11", &ast.Literal{&types.Type{
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
					}, "1100080405", nil, nil, "", nil, nil}, ""},
					&ParentScope{"11"},
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
					}, "1100080406", nil, nil, "", nil, nil}, ""},
					&ParentScope{"12"},
					&Assign{"13", &ast.Literal{&types.Type{
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
					}, "1100080407", nil, nil, "", nil, nil}, ""},
					&ParentScope{"13"},
					&Assign{"14", &ast.Literal{&types.Type{
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
					}, "1100080408", nil, nil, "", nil, nil}, ""},
					&ParentScope{"14"},
					&Assign{"15", &ast.Literal{&types.Type{
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
					}, "1100080409", nil, nil, "", nil, nil}, ""},
					&ParentScope{"15"},
					&Assign{"16", &ast.Literal{&types.Type{
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
					}, "1100080410", nil, nil, "", nil, nil}, ""},
					&ParentScope{"16"},
					&Assign{"1100080403", nil, "1"},
					&Assign{"1100080404", nil, "2"},
					&Assign{"1100080405", nil, "3"},
					&Assign{"1100080406", nil, "4"},
					&Assign{"1100080407", nil, "5"},
					&Assign{"1100080408", nil, "6"},
					&Assign{"1100080409", nil, "7"},
					&Assign{"1100080410", nil, "8"},
					&Assign{"Call", nil, "9"},
					&Assign{"Get", nil, "10"},
					&Assign{"Interface", nil, "11"},
					&Assign{"Kind", nil, "12"},
					&Assign{"Len", nil, "13"},
					&Assign{"Properties", nil, "14"},
					&Assign{"Set", nil, "15"},
					&Assign{"Type", nil, "16"},
					&Return{Registers{"0"}},
				},
				Registers: 16,
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
				Pos:  "lib/reflect/type.ok:8:1",
			},
		},
		"runtime": &File{
			Imports: map[string]*types.Type{
				"strings": &types.Type{
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
			Funcs: map[string]*CompiledFunc{
				"3243469102": &CompiledFunc{
					Arguments: []string{"name"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
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
								&types.Type{
									Kind: 3,
								},
							},
						}, "3243469103", nil, nil, "", nil, nil}, ""},
						&Call{"*2", Registers{"name"}, Registers{"3", "4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"value", nil, "3"},
						&Assign{"_", nil, "4"},
						&Return{Registers{"value"}},
					},
					Registers: 4,
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
					Name:       "Env",
					UniqueName: "3243469102",
					Pos:        "lib/runtime/env.ok:4:1",
				},
				"3243469103": &CompiledFunc{
					Arguments: []string{"name"},
					Instructions: []Instruction{
						&EnvGet{"name", "2", "3"},
						&Return{Registers{"2", "3"}},
					},
					Registers: 3,
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
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "LookupEnv",
					UniqueName: "3243469103",
					Pos:        "lib/runtime/env.ok:13:1",
				},
				"3243469104": &CompiledFunc{
					Arguments: []string{"name", "value"},
					Instructions: []Instruction{
						&EnvSet{"name", "value"},
						&Return{nil},
					},
					Registers: 2,
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
					},
					Name:       "SetEnv",
					UniqueName: "3243469104",
					Pos:        "lib/runtime/env.ok:20:1",
				},
				"3243469105": &CompiledFunc{
					Arguments: []string{"name"},
					Instructions: []Instruction{
						&EnvUnset{"name"},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "UnsetEnv",
					UniqueName: "3243469105",
					Pos:        "lib/runtime/env.ok:26:1",
				},
				"3243469106": &CompiledFunc{
					Arguments: []string{"status"},
					Instructions: []Instruction{
						&Exit{"status"},
					},
					Registers: 1,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Exit",
					UniqueName: "3243469106",
					Pos:        "lib/runtime/exit.ok:3:1",
				},
				"3243469107": &CompiledFunc{
					Arguments: []string{"File", "LineNumber", "LineOffset", "FunctionName"},
					Instructions: []Instruction{
						&Return{Registers{"0"}},
					},
					Registers: 4,
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
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackElement",
								Properties: map[string]*types.Type{
									"File": &types.Type{
										Kind: 7,
									},
									"FunctionName": &types.Type{
										Kind: 7,
									},
									"LineNumber": &types.Type{
										Kind: 6,
									},
									"LineOffset": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "StackElement",
					UniqueName: "3243469107",
					Pos:        "lib/runtime/stack.ok:3:1",
				},
				"3243469108": &CompiledFunc{
					Instructions: []Instruction{
						&Stack{"1"},
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "\n", nil, nil, "lib/runtime/stack.ok:12:41", nil, nil}, ""},
						&LoadPackage{"3", "strings"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Split", nil, nil, "", nil, nil}, ""},
						&MapGet{"3", "4", "5"},
						&Call{"*5", Registers{"1", "2"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"elements", nil, "6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"7", "8", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Name: "StackElement",
							},
						}},
						&Assign{"stack", nil, "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&NextArray{"elements", "9", "", "element", "10"},
						&JumpUnless{"10", 44},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 7,
						}, "|", nil, nil, "lib/runtime/stack.ok:15:40", nil, nil}, ""},
						&LoadPackage{"12", "strings"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Split", nil, nil, "", nil, nil}, ""},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"element", "11"}, Registers{"15"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"parts", nil, "15"},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/runtime/stack.ok:16:45", nil, nil}, ""},
						&ArrayGet{"parts", "16", "17"},
						&Assign{"18", &ast.Literal{&types.Type{
							Kind: 7,
						}, ":", nil, nil, "lib/runtime/stack.ok:16:49", nil, nil}, ""},
						&LoadPackage{"19", "strings"},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Split", nil, nil, "", nil, nil}, ""},
						&MapGet{"19", "20", "21"},
						&Call{"*21", Registers{"17", "18"}, Registers{"22"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"locationParts", nil, "22"},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"23", "24", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 1,
								Name: "StackElement",
								Properties: map[string]*types.Type{
									"File": &types.Type{
										Kind: 7,
									},
									"FunctionName": &types.Type{
										Kind: 7,
									},
									"LineNumber": &types.Type{
										Kind: 6,
									},
									"LineOffset": &types.Type{
										Kind: 6,
									},
								},
							},
						}},
						&Assign{"25", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&Assign{"26", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/runtime/stack.ok:18:27", nil, nil}, ""},
						&ArrayGet{"locationParts", "26", "27"},
						&Assign{"28", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/runtime/stack.ok:19:34", nil, nil}, ""},
						&ArrayGet{"locationParts", "28", "29"},
						&CastNumber{"29", "30"},
						&Assign{"31", &ast.Literal{&types.Type{
							Kind: 6,
						}, "2", nil, nil, "lib/runtime/stack.ok:20:34", nil, nil}, ""},
						&ArrayGet{"locationParts", "31", "32"},
						&CastNumber{"32", "33"},
						&Assign{"34", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/runtime/stack.ok:21:19", nil, nil}, ""},
						&ArrayGet{"parts", "34", "35"},
						&Assign{"36", &ast.Literal{&types.Type{
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
								&types.Type{
									Kind: 7,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "StackElement",
									Properties: map[string]*types.Type{
										"File": &types.Type{
											Kind: 7,
										},
										"FunctionName": &types.Type{
											Kind: 7,
										},
										"LineNumber": &types.Type{
											Kind: 6,
										},
										"LineOffset": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						}, "3243469107", nil, nil, "", nil, nil}, ""},
						&Call{"*36", Registers{"27", "30", "33", "35"}, Registers{"37"}, &types.Type{
							Kind: 1,
							Name: "StackElement",
							Properties: map[string]*types.Type{
								"File": &types.Type{
									Kind: 7,
								},
								"FunctionName": &types.Type{
									Kind: 7,
								},
								"LineNumber": &types.Type{
									Kind: 6,
								},
								"LineOffset": &types.Type{
									Kind: 6,
								},
							},
						}},
						&ArraySet{"24", "25", "37"},
						&Append{"stack", "24", "stack"},
						&Jump{10},
						&Assign{"38", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 8,
									Element: &types.Type{
										Name: "StackElement",
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "StackTrace",
									Properties: map[string]*types.Type{
										"Elements": &types.Type{
											Kind: 8,
											Element: &types.Type{
												Name: "StackElement",
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3243469109", nil, nil, "", nil, nil}, ""},
						&Call{"*38", Registers{"stack"}, Registers{"39"}, &types.Type{
							Kind: 1,
							Name: "StackTrace",
							Properties: map[string]*types.Type{
								"Elements": &types.Type{
									Kind: 8,
									Element: &types.Type{
										Name: "StackElement",
									},
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Return{Registers{"39"}},
					},
					Registers: 39,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
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
					Name:       "Stack",
					UniqueName: "3243469108",
					Pos:        "lib/runtime/stack.ok:11:1",
				},
				"3243469109": &CompiledFunc{
					Arguments: []string{"Elements"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3243469110", nil, nil, "", nil, nil}, ""},
						&ParentScope{"2"},
						&Assign{"String", nil, "2"},
						&Return{Registers{"0"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Name: "StackElement",
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
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
					Name:       "StackTrace",
					UniqueName: "3243469109",
					Pos:        "lib/runtime/stack.ok:27:1",
				},
				"3243469110": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/runtime/stack.ok:29:13", nil, nil}, ""},
						&Assign{"s", nil, "1"},
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/runtime/stack.ok:30:17", nil, nil}, ""},
						&Assign{"i", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&NextArray{"^Elements", "3", "", "element", "4"},
						&JumpUnless{"4", 21},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/runtime/stack.ok:31:22", nil, nil}, ""},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 7,
						}, "FunctionName", nil, nil, "", nil, nil}, ""},
						&MapGet{"element", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/runtime/stack.ok:31:45", nil, nil}, ""},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 7,
						}, "File", nil, nil, "", nil, nil}, ""},
						&MapGet{"element", "10", "11"},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 7,
						}, ":", nil, nil, "lib/runtime/stack.ok:31:60", nil, nil}, ""},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 7,
						}, "LineNumber", nil, nil, "", nil, nil}, ""},
						&MapGet{"element", "13", "14"},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 7,
						}, "\n", nil, nil, "lib/runtime/stack.ok:31:81", nil, nil}, ""},
						&Interpolate{"5", Registers{"i", "6", "8", "9", "11", "12", "14", "15"}},
						&Concat{"s", "5", "s"},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "16", "i"},
						&Jump{4},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 7,
						}, "\n", nil, nil, "lib/runtime/stack.ok:35:37", nil, nil}, ""},
						&LoadPackage{"18", "strings"},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 7,
						}, "TrimRight", nil, nil, "", nil, nil}, ""},
						&MapGet{"18", "19", "20"},
						&Call{"*20", Registers{"s", "17"}, Registers{"21"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"21"}},
					},
					Registers: 21,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "String",
					UniqueName: "3243469110",
					Pos:        "lib/runtime/stack.ok:28:5",
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
					}, "3243469102", nil, nil, "", nil, nil}, ""},
					&ParentScope{"1"},
					&Assign{"2", &ast.Literal{&types.Type{
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
							&types.Type{
								Kind: 3,
							},
						},
					}, "3243469103", nil, nil, "", nil, nil}, ""},
					&ParentScope{"2"},
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
					}, "3243469104", nil, nil, "", nil, nil}, ""},
					&ParentScope{"3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3243469105", nil, nil, "", nil, nil}, ""},
					&ParentScope{"4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3243469106", nil, nil, "", nil, nil}, ""},
					&ParentScope{"5"},
					&Assign{"6", &ast.Literal{&types.Type{
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
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackElement",
								Properties: map[string]*types.Type{
									"File": &types.Type{
										Kind: 7,
									},
									"FunctionName": &types.Type{
										Kind: 7,
									},
									"LineNumber": &types.Type{
										Kind: 6,
									},
									"LineOffset": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3243469107", nil, nil, "", nil, nil}, ""},
					&ParentScope{"6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3243469108", nil, nil, "", nil, nil}, ""},
					&ParentScope{"7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Name: "StackElement",
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3243469109", nil, nil, "", nil, nil}, ""},
					&ParentScope{"8"},
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
					}, "3243469102", nil, nil, "", nil, nil}, ""},
					&ParentScope{"9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3243469106", nil, nil, "", nil, nil}, ""},
					&ParentScope{"10"},
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
							&types.Type{
								Kind: 3,
							},
						},
					}, "3243469103", nil, nil, "", nil, nil}, ""},
					&ParentScope{"11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
							&types.Type{
								Kind: 7,
							},
						},
					}, "3243469104", nil, nil, "", nil, nil}, ""},
					&ParentScope{"12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3243469108", nil, nil, "", nil, nil}, ""},
					&ParentScope{"13"},
					&Assign{"14", &ast.Literal{&types.Type{
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
							&types.Type{
								Kind: 7,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackElement",
								Properties: map[string]*types.Type{
									"File": &types.Type{
										Kind: 7,
									},
									"FunctionName": &types.Type{
										Kind: 7,
									},
									"LineNumber": &types.Type{
										Kind: 6,
									},
									"LineOffset": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3243469107", nil, nil, "", nil, nil}, ""},
					&ParentScope{"14"},
					&Assign{"15", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 8,
								Element: &types.Type{
									Name: "StackElement",
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "StackTrace",
								Properties: map[string]*types.Type{
									"Elements": &types.Type{
										Kind: 8,
										Element: &types.Type{
											Name: "StackElement",
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3243469109", nil, nil, "", nil, nil}, ""},
					&ParentScope{"15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3243469105", nil, nil, "", nil, nil}, ""},
					&ParentScope{"16"},
					&Assign{"3243469102", nil, "1"},
					&Assign{"3243469103", nil, "2"},
					&Assign{"3243469104", nil, "3"},
					&Assign{"3243469105", nil, "4"},
					&Assign{"3243469106", nil, "5"},
					&Assign{"3243469107", nil, "6"},
					&Assign{"3243469108", nil, "7"},
					&Assign{"3243469109", nil, "8"},
					&Assign{"Env", nil, "9"},
					&Assign{"Exit", nil, "10"},
					&Assign{"LookupEnv", nil, "11"},
					&Assign{"SetEnv", nil, "12"},
					&Assign{"Stack", nil, "13"},
					&Assign{"StackElement", nil, "14"},
					&Assign{"StackTrace", nil, "15"},
					&Assign{"UnsetEnv", nil, "16"},
					&Return{Registers{"0"}},
				},
				Registers: 16,
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__runtime",
							Properties: map[string]*types.Type{
								"Env": &types.Type{
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
								"Exit": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"LookupEnv": &types.Type{
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
										&types.Type{
											Kind: 3,
										},
									},
								},
								"SetEnv": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Stack": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "StackTrace",
											Properties: map[string]*types.Type{
												"Elements": &types.Type{
													Kind: 8,
													Element: &types.Type{
														Name: "StackElement",
													},
												},
												"String": &types.Type{
													Kind: 10,
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
								"StackElement": &types.Type{
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
										&types.Type{
											Kind: 7,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "StackElement",
											Properties: map[string]*types.Type{
												"File": &types.Type{
													Kind: 7,
												},
												"FunctionName": &types.Type{
													Kind: 7,
												},
												"LineNumber": &types.Type{
													Kind: 6,
												},
												"LineOffset": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"StackTrace": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 8,
											Element: &types.Type{
												Name: "StackElement",
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "StackTrace",
											Properties: map[string]*types.Type{
												"Elements": &types.Type{
													Kind: 8,
													Element: &types.Type{
														Name: "StackElement",
													},
												},
												"String": &types.Type{
													Kind: 10,
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
								"UnsetEnv": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						},
					},
				},
				Name: "..__runtime",
				Pos:  "lib/runtime/stack.ok:1:1",
			},
		},
		"strings": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"601598823": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/case.ok:5:14", nil, nil}, ""},
						&Assign{"result", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 4,
						}, "A", nil, nil, "lib/strings/case.ok:8:24", nil, nil}, ""},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 4,
						}, "Z", nil, nil, "lib/strings/case.ok:8:44", nil, nil}, ""},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "32", nil, nil, "lib/strings/case.ok:9:40", nil, nil}, ""},
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
					UniqueName: "601598823",
					Pos:        "lib/strings/case.ok:4:1",
				},
				"601598824": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/case.ok:22:14", nil, nil}, ""},
						&Assign{"result", nil, "2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 4,
						}, "a", nil, nil, "lib/strings/case.ok:25:24", nil, nil}, ""},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 4,
						}, "z", nil, nil, "lib/strings/case.ok:25:44", nil, nil}, ""},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "32", nil, nil, "lib/strings/case.ok:26:40", nil, nil}, ""},
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
					UniqueName: "601598824",
					Pos:        "lib/strings/case.ok:21:1",
				},
				"601598825": &CompiledFunc{
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
						}, "601598828", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"s", "substr"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/contains.ok:3:32", nil, nil}, ""},
						&NotEqualNumber{"4", "5", "6"},
						&Return{Registers{"6"}},
					},
					Registers: 6,
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
					UniqueName: "601598825",
					Pos:        "lib/strings/contains.ok:2:1",
				},
				"601598826": &CompiledFunc{
					Arguments: []string{"s", "prefix"},
					Instructions: []Instruction{
						&Len{"s", "3"},
						&Len{"prefix", "4"},
						&LessThanNumber{"3", "4", "5"},
						&JumpUnless{"5", 5},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:9:16", nil, nil}, ""},
						&Return{Registers{"6"}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/contains.ok:12:13", nil, nil}, ""},
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
						}, "false", nil, nil, "lib/strings/contains.ok:14:20", nil, nil}, ""},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "14", "i"},
						&Jump{8},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/contains.ok:18:12", nil, nil}, ""},
						&Return{Registers{"15"}},
					},
					Registers: 15,
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
					UniqueName: "601598826",
					Pos:        "lib/strings/contains.ok:7:1",
				},
				"601598827": &CompiledFunc{
					Arguments: []string{"s", "suffix"},
					Instructions: []Instruction{
						&Len{"s", "3"},
						&Len{"suffix", "4"},
						&LessThanNumber{"3", "4", "5"},
						&JumpUnless{"5", 5},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:24:16", nil, nil}, ""},
						&Return{Registers{"6"}},
						&Len{"s", "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/contains.ok:27:18", nil, nil}, ""},
						&Subtract{"7", "8", "9"},
						&Assign{"j", nil, "9"},
						&Len{"suffix", "10"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/contains.ok:28:27", nil, nil}, ""},
						&Subtract{"10", "11", "12"},
						&Assign{"i", nil, "12"},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/contains.ok:28:35", nil, nil}, ""},
						&GreaterThanEqualNumber{"i", "13", "14"},
						&JumpUnless{"14", 27},
						&StringIndex{"s", "j", "15"},
						&StringIndex{"suffix", "i", "16"},
						&NotEqual{"15", "16", "17"},
						&JumpUnless{"17", 22},
						&Assign{"18", &ast.Literal{&types.Type{
							Kind: 3,
						}, "false", nil, nil, "lib/strings/contains.ok:30:20", nil, nil}, ""},
						&Return{Registers{"18"}},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Subtract{"j", "19", "j"},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Subtract{"i", "20", "i"},
						&Jump{14},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/contains.ok:36:12", nil, nil}, ""},
						&Return{Registers{"21"}},
					},
					Registers: 21,
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
					UniqueName: "601598827",
					Pos:        "lib/strings/contains.ok:22:1",
				},
				"601598828": &CompiledFunc{
					Arguments: []string{"s", "substr"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:3:34", nil, nil}, ""},
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
						}, "601598829", nil, nil, "", nil, nil}, ""},
						&Call{"*4", Registers{"s", "substr", "3"}, Registers{"5"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"5"}},
					},
					Registers: 5,
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
					UniqueName: "601598828",
					Pos:        "lib/strings/index.ok:2:1",
				},
				"601598829": &CompiledFunc{
					Arguments: []string{"s", "substr", "offset"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:18:26", nil, nil}, ""},
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
						}, "601598831", nil, nil, "", nil, nil}, ""},
						&Call{"*5", Registers{"offset", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"offset", nil, "6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:20:22", nil, nil}, ""},
						&Add{"offset", "7", "8"},
						&Assign{"i", nil, "8"},
						&Len{"s", "9"},
						&Len{"substr", "10"},
						&Subtract{"9", "10", "11"},
						&LessThanEqualNumber{"i", "11", "12"},
						&JumpUnless{"12", 34},
						&Assign{"13", &ast.Literal{&types.Type{
							Kind: 3,
						}, "true", nil, nil, "lib/strings/index.ok:21:17", nil, nil}, ""},
						&Assign{"found", nil, "13"},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/index.ok:23:17", nil, nil}, ""},
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
						}, "false", nil, nil, "lib/strings/index.ok:25:25", nil, nil}, ""},
						&Assign{"found", nil, "21"},
						&Jump{29},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"j", "22", "j"},
						&Jump{16},
						&JumpUnless{"found", 31},
						&Return{Registers{"i"}},
						&Assign{"23", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "23", "i"},
						&Jump{9},
						&Assign{"24", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:35:12", nil, nil}, ""},
						&Return{Registers{"24"}},
					},
					Registers: 24,
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
					UniqueName: "601598829",
					Pos:        "lib/strings/index.ok:17:1",
				},
				"601598830": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&LessThanNumber{"a", "b", "3"},
						&JumpUnless{"3", 2},
						&Return{Registers{"a"}},
						&Return{Registers{"b"}},
					},
					Registers: 3,
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
					UniqueName: "601598830",
					Pos:        "lib/strings/index.ok:39:1",
				},
				"601598831": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&GreaterThanNumber{"a", "b", "3"},
						&JumpUnless{"3", 2},
						&Return{Registers{"a"}},
						&Return{Registers{"b"}},
					},
					Registers: 3,
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
					UniqueName: "601598831",
					Pos:        "lib/strings/index.ok:48:1",
				},
				"601598832": &CompiledFunc{
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598828", nil, nil, "", nil, nil}, ""},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:59:17", nil, nil}, ""},
						&EqualNumber{"index", "9", "10"},
						&JumpUnless{"10", 11},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:60:16", nil, nil}, ""},
						&Return{Registers{"11"}},
						&Len{"s", "12"},
						&Len{"substr", "13"},
						&Add{"index", "13", "14"},
						&Subtract{"12", "14", "15"},
						&Return{Registers{"15"}},
					},
					Registers: 15,
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
					UniqueName: "601598832",
					Pos:        "lib/strings/index.ok:57:1",
				},
				"601598833": &CompiledFunc{
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
						}, "601598830", nil, nil, "", nil, nil}, ""},
						&Call{"*6", Registers{"offset", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/index.ok:79:45", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598829", nil, nil, "", nil, nil}, ""},
						&Call{"*15", Registers{"12", "14", "offset"}, Registers{"16"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"index", nil, "16"},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:82:17", nil, nil}, ""},
						&EqualNumber{"index", "17", "18"},
						&JumpUnless{"18", 19},
						&Assign{"19", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/index.ok:83:16", nil, nil}, ""},
						&Return{Registers{"19"}},
						&Len{"s", "20"},
						&Len{"substr", "21"},
						&Add{"index", "21", "22"},
						&Subtract{"20", "22", "23"},
						&Return{Registers{"23"}},
					},
					Registers: 23,
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
					UniqueName: "601598833",
					Pos:        "lib/strings/index.ok:76:1",
				},
				"601598834": &CompiledFunc{
					Arguments: []string{"strings", "glue"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/join.ok:5:14", nil, nil}, ""},
						&Assign{"result", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&NextArray{"strings", "4", "i", "s", "5"},
						&JumpUnless{"5", 10},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/join.ok:7:16", nil, nil}, ""},
						&GreaterThanNumber{"i", "6", "7"},
						&JumpUnless{"7", 8},
						&Concat{"result", "glue", "result"},
						&Concat{"result", "s", "result"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 7,
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
					UniqueName: "601598834",
					Pos:        "lib/strings/join.ok:4:1",
				},
				"601598835": &CompiledFunc{
					Arguments: []string{"s", "pad", "toLen"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&GreaterThanEqualNumber{"4", "toLen", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/pad.ok:10:34", nil, nil}, ""},
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
						}, "601598837", nil, nil, "", nil, nil}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"12", "s", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					UniqueName: "601598835",
					Pos:        "lib/strings/pad.ok:9:1",
				},
				"601598836": &CompiledFunc{
					Arguments: []string{"s", "pad", "toLen"},
					Instructions: []Instruction{
						&Len{"s", "4"},
						&GreaterThanEqualNumber{"4", "toLen", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/pad.ok:20:34", nil, nil}, ""},
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
						}, "601598837", nil, nil, "", nil, nil}, ""},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, &types.Type{
							Kind: 2,
						}},
						&Concat{"s", "12", "13"},
						&Return{Registers{"13"}},
					},
					Registers: 13,
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
					UniqueName: "601598836",
					Pos:        "lib/strings/pad.ok:19:1",
				},
				"601598837": &CompiledFunc{
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
						}, "601598838", nil, nil, "", nil, nil}, ""},
						&Call{"*5", Registers{"pad", "4"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/pad.ok:28:50", nil, nil}, ""},
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
						}, "601598842", nil, nil, "", nil, nil}, ""},
						&Call{"*8", Registers{"6", "7", "toLen"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"9"}},
					},
					Registers: 9,
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
					UniqueName: "601598837",
					Pos:        "lib/strings/pad.ok:27:1",
				},
				"601598838": &CompiledFunc{
					Arguments: []string{"str", "times"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/repeat.ok:4:14", nil, nil}, ""},
						&Assign{"result", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/repeat.ok:5:13", nil, nil}, ""},
						&Assign{"i", nil, "4"},
						&LessThanNumber{"i", "times", "5"},
						&JumpUnless{"5", 9},
						&Concat{"result", "str", "result"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "6", "i"},
						&Jump{3},
						&Return{Registers{"result"}},
					},
					Registers: 6,
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
					UniqueName: "601598838",
					Pos:        "lib/strings/repeat.ok:3:1",
				},
				"601598839": &CompiledFunc{
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
						}, "601598841", nil, nil, "", nil, nil}, ""},
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
						}, "601598834", nil, nil, "", nil, nil}, ""},
						&Call{"*6", Registers{"5", "replace"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
					},
					Registers: 7,
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
					UniqueName: "601598839",
					Pos:        "lib/strings/replace.ok:5:1",
				},
				"601598840": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/reverse.ok:3:14", nil, nil}, ""},
						&Assign{"result", nil, "2"},
						&Len{"s", "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/reverse.ok:4:22", nil, nil}, ""},
						&Subtract{"3", "4", "5"},
						&Assign{"i", nil, "5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/reverse.ok:4:30", nil, nil}, ""},
						&GreaterThanEqualNumber{"i", "6", "7"},
						&JumpUnless{"7", 14},
						&StringIndex{"s", "i", "8"},
						&CastString{"8", "9"},
						&Concat{"result", "9", "result"},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Subtract{"i", "10", "i"},
						&Jump{6},
						&Return{Registers{"result"}},
					},
					Registers: 10,
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
					UniqueName: "601598840",
					Pos:        "lib/strings/reverse.ok:2:1",
				},
				"601598841": &CompiledFunc{
					Arguments: []string{"s", "delimiter"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"3", "4", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"elements", nil, "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:10:21", nil, nil}, ""},
						&Equal{"delimiter", "5", "6"},
						&JumpUnless{"6", 21},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:13:17", nil, nil}, ""},
						&Assign{"i", nil, "7"},
						&Len{"s", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 20},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"10", "11", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&StringIndex{"s", "i", "13"},
						&CastString{"13", "14"},
						&ArraySet{"11", "12", "14"},
						&Append{"elements", "11", "elements"},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "15", "i"},
						&Jump{8},
						&Jump{59},
						&Assign{"16", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:17:19", nil, nil}, ""},
						&Assign{"element", nil, "16"},
						&Assign{"17", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:18:17", nil, nil}, ""},
						&Assign{"i", nil, "17"},
						&Len{"s", "18"},
						&LessThanNumber{"i", "18", "19"},
						&JumpUnless{"19", 54},
						&Assign{"20", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:19:45", nil, nil}, ""},
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
						}, "601598829", nil, nil, "", nil, nil}, ""},
						&Call{"*22", Registers{"s", "delimiter", "21"}, Registers{"23"}, &types.Type{
							Kind: 2,
						}},
						&Subtract{"23", "i", "24"},
						&Assign{"25", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/split.ok:19:55", nil, nil}, ""},
						&EqualNumber{"24", "25", "26"},
						&JumpUnless{"26", 48},
						&Assign{"27", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"27", "28", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"29", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&ArraySet{"28", "29", "element"},
						&Append{"elements", "28", "elements"},
						&Assign{"30", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/split.ok:21:27", nil, nil}, ""},
						&Assign{"element", nil, "30"},
						&Len{"delimiter", "31"},
						&Assign{"32", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/strings/split.ok:22:39", nil, nil}, ""},
						&Subtract{"31", "32", "33"},
						&Add{"i", "33", "i"},
						&Jump{51},
						&StringIndex{"s", "i", "34"},
						&CastString{"34", "35"},
						&Concat{"element", "35", "element"},
						&Assign{"36", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "36", "i"},
						&Jump{26},
						&Assign{"37", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&ArrayAlloc{"37", "38", &types.Type{
							Kind: 8,
							Element: &types.Type{
								Kind: 7,
							},
						}},
						&Assign{"39", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "", nil, nil}, ""},
						&ArraySet{"38", "39", "element"},
						&Append{"elements", "38", "elements"},
						&Return{Registers{"elements"}},
					},
					Registers: 39,
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
					UniqueName: "601598841",
					Pos:        "lib/strings/split.ok:7:1",
				},
				"601598842": &CompiledFunc{
					Arguments: []string{"s", "fromIndex", "toIndex"},
					Instructions: []Instruction{
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/substr.ok:4:9", nil, nil}, ""},
						&Assign{"r", nil, "4"},
						&Assign{"i", nil, "fromIndex"},
						&LessThanNumber{"i", "toIndex", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "i", "6"},
						&CastString{"6", "7"},
						&Concat{"r", "7", "r"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"i", "8", "i"},
						&Jump{2},
						&Return{Registers{"r"}},
					},
					Registers: 8,
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
					UniqueName: "601598842",
					Pos:        "lib/strings/substr.ok:3:1",
				},
				"601598843": &CompiledFunc{
					Arguments: []string{"s", "cutset"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/strings/trim.ok:4:18", nil, nil}, ""},
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
						}, "601598828", nil, nil, "", nil, nil}, ""},
						&Call{"*8", Registers{"cutset", "7"}, Registers{"9"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"10", &ast.Literal{&types.Type{
							Kind: 6,
						}, "-1", nil, nil, "lib/strings/trim.ok:5:47", nil, nil}, ""},
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
						}, "601598848", nil, nil, "", nil, nil}, ""},
						&Call{"*12", Registers{"s", "offset"}, Registers{"13"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"13"}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"offset", "14", "offset"},
						&Jump{2},
						&Return{Registers{"s"}},
					},
					Registers: 14,
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
					UniqueName: "601598843",
					Pos:        "lib/strings/trim.ok:3:1",
				},
				"601598844": &CompiledFunc{
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598843", nil, nil, "", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
						&Call{"*7", Registers{"6"}, Registers{"8"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"8"}},
					},
					Registers: 8,
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
					UniqueName: "601598844",
					Pos:        "lib/strings/trim.ok:15:1",
				},
				"601598845": &CompiledFunc{
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
						}, "601598843", nil, nil, "", nil, nil}, ""},
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
						}, "601598844", nil, nil, "", nil, nil}, ""},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"6"}},
					},
					Registers: 6,
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
					UniqueName: "601598845",
					Pos:        "lib/strings/trim.ok:21:1",
				},
				"601598846": &CompiledFunc{
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
						}, "601598826", nil, nil, "", nil, nil}, ""},
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
						}, "601598848", nil, nil, "", nil, nil}, ""},
						&Call{"*6", Registers{"s", "5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"7"}},
						&Return{Registers{"s"}},
					},
					Registers: 7,
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
					UniqueName: "601598846",
					Pos:        "lib/strings/trim.ok:32:1",
				},
				"601598847": &CompiledFunc{
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
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
						}, "601598846", nil, nil, "", nil, nil}, ""},
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
						}, "601598840", nil, nil, "", nil, nil}, ""},
						&Call{"*9", Registers{"8"}, Registers{"10"}, &types.Type{
							Kind: 2,
						}},
						&Return{Registers{"10"}},
					},
					Registers: 10,
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
					UniqueName: "601598847",
					Pos:        "lib/strings/trim.ok:47:1",
				},
				"601598848": &CompiledFunc{
					Arguments: []string{"s", "index"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/strings/trim.ok:53:14", nil, nil}, ""},
						&Assign{"result", nil, "3"},
						&Len{"s", "4"},
						&LessThanNumber{"index", "4", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "index", "6"},
						&CastString{"6", "7"},
						&Concat{"result", "7", "result"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "", nil, nil}, ""},
						&Add{"index", "8", "index"},
						&Jump{2},
						&Return{Registers{"result"}},
					},
					Registers: 8,
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
					UniqueName: "601598848",
					Pos:        "lib/strings/trim.ok:52:1",
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
					}, "601598823", nil, nil, "", nil, nil}, ""},
					&ParentScope{"1"},
					&Assign{"2", &ast.Literal{&types.Type{
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
					}, "601598824", nil, nil, "", nil, nil}, ""},
					&ParentScope{"2"},
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
					}, "601598825", nil, nil, "", nil, nil}, ""},
					&ParentScope{"3"},
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
								Kind: 3,
							},
						},
					}, "601598826", nil, nil, "", nil, nil}, ""},
					&ParentScope{"4"},
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
					}, "601598827", nil, nil, "", nil, nil}, ""},
					&ParentScope{"5"},
					&Assign{"6", &ast.Literal{&types.Type{
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
					}, "601598828", nil, nil, "", nil, nil}, ""},
					&ParentScope{"6"},
					&Assign{"7", &ast.Literal{&types.Type{
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
					}, "601598829", nil, nil, "", nil, nil}, ""},
					&ParentScope{"7"},
					&Assign{"8", &ast.Literal{&types.Type{
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
					}, "601598830", nil, nil, "", nil, nil}, ""},
					&ParentScope{"8"},
					&Assign{"9", &ast.Literal{&types.Type{
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
					}, "601598831", nil, nil, "", nil, nil}, ""},
					&ParentScope{"9"},
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
								Kind: 6,
							},
						},
					}, "601598832", nil, nil, "", nil, nil}, ""},
					&ParentScope{"10"},
					&Assign{"11", &ast.Literal{&types.Type{
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
					}, "601598833", nil, nil, "", nil, nil}, ""},
					&ParentScope{"11"},
					&Assign{"12", &ast.Literal{&types.Type{
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
					}, "601598834", nil, nil, "", nil, nil}, ""},
					&ParentScope{"12"},
					&Assign{"13", &ast.Literal{&types.Type{
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
					}, "601598835", nil, nil, "", nil, nil}, ""},
					&ParentScope{"13"},
					&Assign{"14", &ast.Literal{&types.Type{
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
					}, "601598836", nil, nil, "", nil, nil}, ""},
					&ParentScope{"14"},
					&Assign{"15", &ast.Literal{&types.Type{
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
					}, "601598837", nil, nil, "", nil, nil}, ""},
					&ParentScope{"15"},
					&Assign{"16", &ast.Literal{&types.Type{
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
					}, "601598838", nil, nil, "", nil, nil}, ""},
					&ParentScope{"16"},
					&Assign{"17", &ast.Literal{&types.Type{
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
					}, "601598839", nil, nil, "", nil, nil}, ""},
					&ParentScope{"17"},
					&Assign{"18", &ast.Literal{&types.Type{
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
					}, "601598840", nil, nil, "", nil, nil}, ""},
					&ParentScope{"18"},
					&Assign{"19", &ast.Literal{&types.Type{
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
					}, "601598841", nil, nil, "", nil, nil}, ""},
					&ParentScope{"19"},
					&Assign{"20", &ast.Literal{&types.Type{
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
					}, "601598842", nil, nil, "", nil, nil}, ""},
					&ParentScope{"20"},
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
								Kind: 7,
							},
						},
					}, "601598843", nil, nil, "", nil, nil}, ""},
					&ParentScope{"21"},
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
								Kind: 7,
							},
						},
					}, "601598844", nil, nil, "", nil, nil}, ""},
					&ParentScope{"22"},
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
								Kind: 7,
							},
						},
					}, "601598845", nil, nil, "", nil, nil}, ""},
					&ParentScope{"23"},
					&Assign{"24", &ast.Literal{&types.Type{
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
					}, "601598846", nil, nil, "", nil, nil}, ""},
					&ParentScope{"24"},
					&Assign{"25", &ast.Literal{&types.Type{
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
					}, "601598847", nil, nil, "", nil, nil}, ""},
					&ParentScope{"25"},
					&Assign{"26", &ast.Literal{&types.Type{
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
					}, "601598848", nil, nil, "", nil, nil}, ""},
					&ParentScope{"26"},
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
					}, "601598825", nil, nil, "", nil, nil}, ""},
					&ParentScope{"27"},
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
					}, "601598826", nil, nil, "", nil, nil}, ""},
					&ParentScope{"28"},
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
					}, "601598827", nil, nil, "", nil, nil}, ""},
					&ParentScope{"29"},
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
					}, "601598828", nil, nil, "", nil, nil}, ""},
					&ParentScope{"30"},
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
					}, "601598829", nil, nil, "", nil, nil}, ""},
					&ParentScope{"31"},
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
					}, "601598834", nil, nil, "", nil, nil}, ""},
					&ParentScope{"32"},
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
					}, "601598832", nil, nil, "", nil, nil}, ""},
					&ParentScope{"33"},
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
					}, "601598833", nil, nil, "", nil, nil}, ""},
					&ParentScope{"34"},
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
					}, "601598835", nil, nil, "", nil, nil}, ""},
					&ParentScope{"35"},
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
					}, "601598836", nil, nil, "", nil, nil}, ""},
					&ParentScope{"36"},
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
					}, "601598838", nil, nil, "", nil, nil}, ""},
					&ParentScope{"37"},
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
					}, "601598839", nil, nil, "", nil, nil}, ""},
					&ParentScope{"38"},
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
					}, "601598840", nil, nil, "", nil, nil}, ""},
					&ParentScope{"39"},
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
					}, "601598841", nil, nil, "", nil, nil}, ""},
					&ParentScope{"40"},
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
					}, "601598842", nil, nil, "", nil, nil}, ""},
					&ParentScope{"41"},
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
					}, "601598823", nil, nil, "", nil, nil}, ""},
					&ParentScope{"42"},
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
					}, "601598824", nil, nil, "", nil, nil}, ""},
					&ParentScope{"43"},
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
					}, "601598845", nil, nil, "", nil, nil}, ""},
					&ParentScope{"44"},
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
					}, "601598843", nil, nil, "", nil, nil}, ""},
					&ParentScope{"45"},
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
					}, "601598846", nil, nil, "", nil, nil}, ""},
					&ParentScope{"46"},
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
					}, "601598844", nil, nil, "", nil, nil}, ""},
					&ParentScope{"47"},
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
					}, "601598847", nil, nil, "", nil, nil}, ""},
					&ParentScope{"48"},
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
					}, "601598837", nil, nil, "", nil, nil}, ""},
					&ParentScope{"49"},
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
					}, "601598831", nil, nil, "", nil, nil}, ""},
					&ParentScope{"50"},
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
					}, "601598830", nil, nil, "", nil, nil}, ""},
					&ParentScope{"51"},
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
					}, "601598848", nil, nil, "", nil, nil}, ""},
					&ParentScope{"52"},
					&Assign{"601598823", nil, "1"},
					&Assign{"601598824", nil, "2"},
					&Assign{"601598825", nil, "3"},
					&Assign{"601598826", nil, "4"},
					&Assign{"601598827", nil, "5"},
					&Assign{"601598828", nil, "6"},
					&Assign{"601598829", nil, "7"},
					&Assign{"601598830", nil, "8"},
					&Assign{"601598831", nil, "9"},
					&Assign{"601598832", nil, "10"},
					&Assign{"601598833", nil, "11"},
					&Assign{"601598834", nil, "12"},
					&Assign{"601598835", nil, "13"},
					&Assign{"601598836", nil, "14"},
					&Assign{"601598837", nil, "15"},
					&Assign{"601598838", nil, "16"},
					&Assign{"601598839", nil, "17"},
					&Assign{"601598840", nil, "18"},
					&Assign{"601598841", nil, "19"},
					&Assign{"601598842", nil, "20"},
					&Assign{"601598843", nil, "21"},
					&Assign{"601598844", nil, "22"},
					&Assign{"601598845", nil, "23"},
					&Assign{"601598846", nil, "24"},
					&Assign{"601598847", nil, "25"},
					&Assign{"601598848", nil, "26"},
					&Assign{"Contains", nil, "27"},
					&Assign{"HasPrefix", nil, "28"},
					&Assign{"HasSuffix", nil, "29"},
					&Assign{"Index", nil, "30"},
					&Assign{"IndexAfter", nil, "31"},
					&Assign{"Join", nil, "32"},
					&Assign{"LastIndex", nil, "33"},
					&Assign{"LastIndexBefore", nil, "34"},
					&Assign{"PadLeft", nil, "35"},
					&Assign{"PadRight", nil, "36"},
					&Assign{"Repeat", nil, "37"},
					&Assign{"ReplaceAll", nil, "38"},
					&Assign{"Reverse", nil, "39"},
					&Assign{"Split", nil, "40"},
					&Assign{"Substr", nil, "41"},
					&Assign{"ToLower", nil, "42"},
					&Assign{"ToUpper", nil, "43"},
					&Assign{"Trim", nil, "44"},
					&Assign{"TrimLeft", nil, "45"},
					&Assign{"TrimPrefix", nil, "46"},
					&Assign{"TrimRight", nil, "47"},
					&Assign{"TrimSuffix", nil, "48"},
					&Assign{"createPad", nil, "49"},
					&Assign{"max", nil, "50"},
					&Assign{"min", nil, "51"},
					&Assign{"substrFrom", nil, "52"},
					&Return{Registers{"0"}},
				},
				Registers: 52,
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
				Pos:  "lib/strings/trim.ok:3:1",
			},
		},
		"time": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"3492229393": &CompiledFunc{
					Arguments: []string{"t", "duration"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229411", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"t"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Seconds", nil, nil, "", nil, nil}, ""},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Add{"4", "7", "8"},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						}, "3492229412", nil, nil, "", nil, nil}, ""},
						&Call{"*9", Registers{"8"}, Registers{"10"}, &types.Type{
							Kind: 1,
							Name: "Time",
							Properties: map[string]*types.Type{
								"Day": &types.Type{
									Kind: 6,
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Month": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Year": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Return{Registers{"10"}},
					},
					Registers: 10,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "Add",
					UniqueName: "3492229393",
					Pos:        "lib/time/add.ok:3:1",
				},
				"3492229394": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229411", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"a"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229411", nil, nil, "", nil, nil}, ""},
						&Call{"*5", Registers{"b"}, Registers{"6"}, &types.Type{
							Kind: 2,
						}},
						&Subtract{"4", "6", "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3492229398", nil, nil, "", nil, nil}, ""},
						&Call{"*8", Registers{"7"}, Registers{"9"}, &types.Type{
							Kind: 1,
							Name: "Duration",
							Properties: map[string]*types.Type{
								"Hours": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Microseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Milliseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Minutes": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Nanoseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Seconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
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
					Name:       "Sub",
					UniqueName: "3492229394",
					Pos:        "lib/time/compare.ok:4:1",
				},
				"3492229395": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3492229394", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, &types.Type{
							Kind: 1,
							Name: "Duration",
							Properties: map[string]*types.Type{
								"Hours": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Microseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Milliseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Minutes": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Nanoseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Seconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Assign{"duration", nil, "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Seconds", nil, nil, "", nil, nil}, ""},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/compare.ok:13:34", nil, nil}, ""},
						&EqualNumber{"7", "8", "9"},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "Equal",
					UniqueName: "3492229395",
					Pos:        "lib/time/compare.ok:10:1",
				},
				"3492229396": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3492229394", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, &types.Type{
							Kind: 1,
							Name: "Duration",
							Properties: map[string]*types.Type{
								"Hours": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Microseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Milliseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Minutes": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Nanoseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Seconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Assign{"duration", nil, "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Seconds", nil, nil, "", nil, nil}, ""},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/compare.ok:20:33", nil, nil}, ""},
						&LessThanNumber{"7", "8", "9"},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "Before",
					UniqueName: "3492229396",
					Pos:        "lib/time/compare.ok:17:1",
				},
				"3492229397": &CompiledFunc{
					Arguments: []string{"a", "b"},
					Instructions: []Instruction{
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
								&types.Type{
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 1,
									Name: "Duration",
									Properties: map[string]*types.Type{
										"Hours": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Microseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Milliseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Minutes": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Nanoseconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"Seconds": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 6,
												},
											},
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
									},
								},
							},
						}, "3492229394", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, &types.Type{
							Kind: 1,
							Name: "Duration",
							Properties: map[string]*types.Type{
								"Hours": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Microseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Milliseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Minutes": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Nanoseconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"Seconds": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
							},
						}},
						&Assign{"duration", nil, "4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Seconds", nil, nil, "", nil, nil}, ""},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/compare.ok:27:33", nil, nil}, ""},
						&GreaterThanNumber{"7", "8", "9"},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					},
					Name:       "After",
					UniqueName: "3492229397",
					Pos:        "lib/time/compare.ok:24:1",
				},
				"3492229398": &CompiledFunc{
					Arguments: []string{"seconds"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229405", nil, nil, "", nil, nil}, ""},
						&ParentScope{"2"},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229404", nil, nil, "", nil, nil}, ""},
						&ParentScope{"3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229403", nil, nil, "", nil, nil}, ""},
						&ParentScope{"4"},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229402", nil, nil, "", nil, nil}, ""},
						&ParentScope{"5"},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229401", nil, nil, "", nil, nil}, ""},
						&ParentScope{"6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229400", nil, nil, "", nil, nil}, ""},
						&ParentScope{"7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
						}, "3492229399", nil, nil, "", nil, nil}, ""},
						&ParentScope{"8"},
						&Assign{"String", nil, "2"},
						&Assign{"Hours", nil, "3"},
						&Assign{"Minutes", nil, "4"},
						&Assign{"Seconds", nil, "5"},
						&Assign{"Milliseconds", nil, "6"},
						&Assign{"Microseconds", nil, "7"},
						&Assign{"Nanoseconds", nil, "8"},
						&Return{Registers{"0"}},
					},
					Registers: 8,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
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
					Name:       "Duration",
					UniqueName: "3492229398",
					Pos:        "lib/time/duration.ok:11:1",
				},
				"3492229399": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.000000001", nil, nil, "", nil, nil}, ""},
						&Divide{"^seconds", "1", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Nanoseconds",
					UniqueName: "3492229399",
					Pos:        "lib/time/duration.ok:12:5",
				},
				"3492229400": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.000001", nil, nil, "", nil, nil}, ""},
						&Divide{"^seconds", "1", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Microseconds",
					UniqueName: "3492229400",
					Pos:        "lib/time/duration.ok:16:5",
				},
				"3492229401": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0.001", nil, nil, "", nil, nil}, ""},
						&Divide{"^seconds", "1", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Milliseconds",
					UniqueName: "3492229401",
					Pos:        "lib/time/duration.ok:20:5",
				},
				"3492229402": &CompiledFunc{
					Instructions: []Instruction{
						&Return{Registers{"^seconds"}},
					},
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Seconds",
					UniqueName: "3492229402",
					Pos:        "lib/time/duration.ok:24:5",
				},
				"3492229403": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 6,
						}, "60", nil, nil, "", nil, nil}, ""},
						&Divide{"^seconds", "1", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Minutes",
					UniqueName: "3492229403",
					Pos:        "lib/time/duration.ok:28:5",
				},
				"3492229404": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3600", nil, nil, "", nil, nil}, ""},
						&Divide{"^seconds", "1", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Hours",
					UniqueName: "3492229404",
					Pos:        "lib/time/duration.ok:32:5",
				},
				"3492229405": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"1", &ast.Literal{&types.Type{
							Kind: 7,
						}, "", nil, nil, "lib/time/duration.ok:37:13", nil, nil}, ""},
						&Assign{"s", nil, "1"},
						&Assign{"seconds", nil, "^seconds"},
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3600", nil, nil, "lib/time/duration.ok:40:23", nil, nil}, ""},
						&GreaterThanEqualNumber{"seconds", "2", "3"},
						&JumpUnless{"3", 16},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3600", nil, nil, "lib/time/duration.ok:41:37", nil, nil}, ""},
						&Divide{"seconds", "4", "5"},
						&Assign{"6", &ast.Literal{&types.Type{
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
						}, "3492229406", nil, nil, "", nil, nil}, ""},
						&Call{"*6", Registers{"5"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"hours", nil, "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "3600", nil, nil, "lib/time/duration.ok:42:32", nil, nil}, ""},
						&Multiply{"hours", "8", "9"},
						&Subtract{"seconds", "9", "seconds"},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 7,
						}, "h", nil, nil, "lib/time/duration.ok:43:28", nil, nil}, ""},
						&Interpolate{"10", Registers{"hours", "11"}},
						&Concat{"s", "10", "s"},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 6,
						}, "60", nil, nil, "lib/time/duration.ok:46:23", nil, nil}, ""},
						&GreaterThanEqualNumber{"seconds", "12", "13"},
						&JumpUnless{"13", 30},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 6,
						}, "60", nil, nil, "lib/time/duration.ok:47:39", nil, nil}, ""},
						&Divide{"seconds", "14", "15"},
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
						}, "3492229406", nil, nil, "", nil, nil}, ""},
						&Call{"*16", Registers{"15"}, Registers{"17"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"minutes", nil, "17"},
						&Assign{"18", &ast.Literal{&types.Type{
							Kind: 6,
						}, "60", nil, nil, "lib/time/duration.ok:48:34", nil, nil}, ""},
						&Multiply{"minutes", "18", "19"},
						&Subtract{"seconds", "19", "seconds"},
						&Assign{"21", &ast.Literal{&types.Type{
							Kind: 7,
						}, "m", nil, nil, "lib/time/duration.ok:49:30", nil, nil}, ""},
						&Interpolate{"20", Registers{"minutes", "21"}},
						&Concat{"s", "20", "s"},
						&Assign{"22", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/duration.ok:52:23", nil, nil}, ""},
						&NotEqualNumber{"seconds", "22", "23"},
						&JumpUnless{"23", 36},
						&Assign{"25", &ast.Literal{&types.Type{
							Kind: 7,
						}, "s", nil, nil, "lib/time/duration.ok:53:30", nil, nil}, ""},
						&Interpolate{"24", Registers{"seconds", "25"}},
						&Concat{"s", "24", "s"},
						&Return{Registers{"s"}},
					},
					Registers: 25,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "String",
					UniqueName: "3492229405",
					Pos:        "lib/time/duration.ok:36:5",
				},
				"3492229406": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/time/math.ok:3:16", nil, nil}, ""},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", nil, "3"},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/math.ok:4:16", nil, nil}, ""},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 6,
						}, "0", nil, nil, "lib/time/math.ok:8:12", nil, nil}, ""},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 13},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 6,
						}, "1", nil, nil, "lib/time/math.ok:9:27", nil, nil}, ""},
						&Add{"frac", "8", "9"},
						&Subtract{"x", "9", "10"},
						&Return{Registers{"10"}},
						&Subtract{"x", "frac", "11"},
						&Return{Registers{"11"}},
					},
					Registers: 11,
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
					Name:       "floor",
					UniqueName: "3492229406",
					Pos:        "lib/time/math.ok:2:1",
				},
				"3492229407": &CompiledFunc{
					Arguments: []string{"duration"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "Seconds", nil, nil, "", nil, nil}, ""},
						&MapGet{"duration", "2", "3"},
						&Call{"*3", nil, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Sleep{"4"},
					},
					Registers: 4,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
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
					Name:       "Sleep",
					UniqueName: "3492229407",
					Pos:        "lib/time/sleep.ok:2:1",
				},
				"3492229408": &CompiledFunc{
					Arguments: []string{"Year", "Month", "Day", "Hour", "Minute", "Second"},
					Instructions: []Instruction{
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229409", nil, nil, "", nil, nil}, ""},
						&ParentScope{"7"},
						&Assign{"String", nil, "7"},
						&Return{Registers{"0"}},
					},
					Registers: 7,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
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
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "Time",
					UniqueName: "3492229408",
					Pos:        "lib/time/time.ok:15:1",
				},
				"3492229409": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"month", nil, "^Month"},
						&Assign{"day", nil, "^Day"},
						&Assign{"hour", nil, "^Hour"},
						&Assign{"minute", nil, "^Minute"},
						&Assign{"second", nil, "^Second"},
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 7,
						}, "-", nil, nil, "lib/time/time.ok:27:24", nil, nil}, ""},
						&Assign{"3", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229413", nil, nil, "", nil, nil}, ""},
						&Call{"*3", Registers{"month"}, Registers{"4"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"5", &ast.Literal{&types.Type{
							Kind: 7,
						}, "-", nil, nil, "lib/time/time.ok:27:41", nil, nil}, ""},
						&Assign{"6", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229413", nil, nil, "", nil, nil}, ""},
						&Call{"*6", Registers{"day"}, Registers{"7"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 7,
						}, " ", nil, nil, "lib/time/time.ok:27:56", nil, nil}, ""},
						&Assign{"9", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229413", nil, nil, "", nil, nil}, ""},
						&Call{"*9", Registers{"hour"}, Registers{"10"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"11", &ast.Literal{&types.Type{
							Kind: 7,
						}, ":", nil, nil, "lib/time/time.ok:27:72", nil, nil}, ""},
						&Assign{"12", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229413", nil, nil, "", nil, nil}, ""},
						&Call{"*12", Registers{"minute"}, Registers{"13"}, &types.Type{
							Kind: 2,
						}},
						&Assign{"14", &ast.Literal{&types.Type{
							Kind: 7,
						}, ":", nil, nil, "lib/time/time.ok:27:90", nil, nil}, ""},
						&Assign{"15", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
							},
							Returns: []*types.Type{
								&types.Type{
									Kind: 7,
								},
							},
						}, "3492229413", nil, nil, "", nil, nil}, ""},
						&Call{"*15", Registers{"second"}, Registers{"16"}, &types.Type{
							Kind: 2,
						}},
						&Interpolate{"1", Registers{"^Year", "2", "4", "5", "7", "8", "10", "11", "13", "14", "16"}},
						&Return{Registers{"1"}},
					},
					Registers: 16,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					},
					Name:       "String",
					UniqueName: "3492229409",
					Pos:        "lib/time/time.ok:16:5",
				},
				"3492229410": &CompiledFunc{
					Instructions: []Instruction{
						&Now{"1", "2", "3", "4", "5", "6"},
						&Assign{"year", nil, "1"},
						&Assign{"month", nil, "2"},
						&Assign{"day", nil, "3"},
						&Assign{"hour", nil, "4"},
						&Assign{"minute", nil, "5"},
						&Assign{"second", nil, "6"},
						&Assign{"7", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
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
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						}, "3492229408", nil, nil, "", nil, nil}, ""},
						&Call{"*7", Registers{"year", "month", "day", "hour", "minute", "second"}, Registers{"8"}, &types.Type{
							Kind: 1,
							Name: "Time",
							Properties: map[string]*types.Type{
								"Day": &types.Type{
									Kind: 6,
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Month": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Year": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Return{Registers{"8"}},
					},
					Registers: 8,
					Type: &types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "Now",
					UniqueName: "3492229410",
					Pos:        "lib/time/time.ok:32:1",
				},
				"3492229411": &CompiledFunc{
					Arguments: []string{"t"},
					Instructions: []Instruction{
						&Unix{"t", "2"},
						&Return{Registers{"2"}},
					},
					Registers: 2,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					},
					Name:       "Unix",
					UniqueName: "3492229411",
					Pos:        "lib/time/unix.ok:3:1",
				},
				"3492229412": &CompiledFunc{
					Arguments: []string{"seconds"},
					Instructions: []Instruction{
						&FromUnix{"seconds", "2", "3", "4", "5", "6", "7"},
						&Assign{"year", nil, "2"},
						&Assign{"month", nil, "3"},
						&Assign{"day", nil, "4"},
						&Assign{"hour", nil, "5"},
						&Assign{"minute", nil, "6"},
						&Assign{"second", nil, "7"},
						&Assign{"8", &ast.Literal{&types.Type{
							Kind: 10,
							Arguments: []*types.Type{
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
								},
								&types.Type{
									Kind: 6,
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
									Kind: 1,
									Name: "Time",
									Properties: map[string]*types.Type{
										"Day": &types.Type{
											Kind: 6,
										},
										"Hour": &types.Type{
											Kind: 6,
										},
										"Minute": &types.Type{
											Kind: 6,
										},
										"Month": &types.Type{
											Kind: 6,
										},
										"Second": &types.Type{
											Kind: 6,
										},
										"String": &types.Type{
											Kind: 10,
											Returns: []*types.Type{
												&types.Type{
													Kind: 7,
												},
											},
										},
										"Year": &types.Type{
											Kind: 6,
										},
									},
								},
							},
						}, "3492229408", nil, nil, "", nil, nil}, ""},
						&Call{"*8", Registers{"year", "month", "day", "hour", "minute", "second"}, Registers{"9"}, &types.Type{
							Kind: 1,
							Name: "Time",
							Properties: map[string]*types.Type{
								"Day": &types.Type{
									Kind: 6,
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Month": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"String": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 7,
										},
									},
								},
								"Year": &types.Type{
									Kind: 6,
								},
							},
						}},
						&Return{Registers{"9"}},
					},
					Registers: 9,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					},
					Name:       "FromUnix",
					UniqueName: "3492229412",
					Pos:        "lib/time/unix.ok:10:1",
				},
				"3492229413": &CompiledFunc{
					Arguments: []string{"n"},
					Instructions: []Instruction{
						&Assign{"2", &ast.Literal{&types.Type{
							Kind: 6,
						}, "10", nil, nil, "lib/time/util.ok:2:12", nil, nil}, ""},
						&LessThanNumber{"n", "2", "3"},
						&JumpUnless{"3", 6},
						&Assign{"4", &ast.Literal{&types.Type{
							Kind: 7,
						}, "0", nil, nil, "lib/time/util.ok:3:16", nil, nil}, ""},
						&CastString{"n", "5"},
						&Concat{"4", "5", "6"},
						&Return{Registers{"6"}},
						&CastString{"n", "7"},
						&Return{Registers{"7"}},
					},
					Registers: 7,
					Type: &types.Type{
						Kind: 10,
						Arguments: []*types.Type{
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
					Name:       "zeroPad",
					UniqueName: "3492229413",
					Pos:        "lib/time/util.ok:1:1",
				},
			},
			Constants: map[string]*ast.Literal{
				"April": &ast.Literal{&types.Type{
					Kind: 6,
				}, "4", nil, nil, "lib/time/time.ok:4:9", nil, nil},
				"August": &ast.Literal{&types.Type{
					Kind: 6,
				}, "8", nil, nil, "lib/time/time.ok:8:10", nil, nil},
				"December": &ast.Literal{&types.Type{
					Kind: 6,
				}, "12", nil, nil, "lib/time/time.ok:12:12", nil, nil},
				"February": &ast.Literal{&types.Type{
					Kind: 6,
				}, "2", nil, nil, "lib/time/time.ok:2:12", nil, nil},
				"Hour": &ast.Literal{&types.Type{
					Kind: 6,
				}, "3600", nil, nil, "lib/time/duration.ok:8:15", nil, nil},
				"January": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1", nil, nil, "lib/time/time.ok:1:11", nil, nil},
				"July": &ast.Literal{&types.Type{
					Kind: 6,
				}, "7", nil, nil, "lib/time/time.ok:7:8", nil, nil},
				"June": &ast.Literal{&types.Type{
					Kind: 6,
				}, "6", nil, nil, "lib/time/time.ok:6:8", nil, nil},
				"March": &ast.Literal{&types.Type{
					Kind: 6,
				}, "3", nil, nil, "lib/time/time.ok:3:9", nil, nil},
				"May": &ast.Literal{&types.Type{
					Kind: 6,
				}, "5", nil, nil, "lib/time/time.ok:5:7", nil, nil},
				"Microsecond": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.000001", nil, nil, "lib/time/duration.ok:4:15", nil, nil},
				"Millisecond": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.001", nil, nil, "lib/time/duration.ok:5:15", nil, nil},
				"Minute": &ast.Literal{&types.Type{
					Kind: 6,
				}, "60", nil, nil, "lib/time/duration.ok:7:15", nil, nil},
				"Nanosecond": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.000000001", nil, nil, "lib/time/duration.ok:3:15", nil, nil},
				"November": &ast.Literal{&types.Type{
					Kind: 6,
				}, "11", nil, nil, "lib/time/time.ok:11:12", nil, nil},
				"October": &ast.Literal{&types.Type{
					Kind: 6,
				}, "10", nil, nil, "lib/time/time.ok:10:11", nil, nil},
				"Second": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1", nil, nil, "lib/time/duration.ok:6:15", nil, nil},
				"September": &ast.Literal{&types.Type{
					Kind: 6,
				}, "9", nil, nil, "lib/time/time.ok:9:13", nil, nil},
			},
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&Assign{"19", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229393", nil, nil, "", nil, nil}, ""},
					&ParentScope{"19"},
					&Assign{"20", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229394", nil, nil, "", nil, nil}, ""},
					&ParentScope{"20"},
					&Assign{"21", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229395", nil, nil, "", nil, nil}, ""},
					&ParentScope{"21"},
					&Assign{"22", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229396", nil, nil, "", nil, nil}, ""},
					&ParentScope{"22"},
					&Assign{"23", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229397", nil, nil, "", nil, nil}, ""},
					&ParentScope{"23"},
					&Assign{"24", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229398", nil, nil, "", nil, nil}, ""},
					&ParentScope{"24"},
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
					}, "3492229406", nil, nil, "", nil, nil}, ""},
					&ParentScope{"25"},
					&Assign{"26", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229407", nil, nil, "", nil, nil}, ""},
					&ParentScope{"26"},
					&Assign{"27", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
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
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229408", nil, nil, "", nil, nil}, ""},
					&ParentScope{"27"},
					&Assign{"28", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229410", nil, nil, "", nil, nil}, ""},
					&ParentScope{"28"},
					&Assign{"29", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3492229411", nil, nil, "", nil, nil}, ""},
					&ParentScope{"29"},
					&Assign{"30", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229412", nil, nil, "", nil, nil}, ""},
					&ParentScope{"30"},
					&Assign{"31", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3492229413", nil, nil, "", nil, nil}, ""},
					&ParentScope{"31"},
					&Assign{"32", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229393", nil, nil, "", nil, nil}, ""},
					&ParentScope{"32"},
					&Assign{"33", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229397", nil, nil, "", nil, nil}, ""},
					&ParentScope{"33"},
					&Assign{"34", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229396", nil, nil, "", nil, nil}, ""},
					&ParentScope{"34"},
					&Assign{"35", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229398", nil, nil, "", nil, nil}, ""},
					&ParentScope{"35"},
					&Assign{"36", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 3,
							},
						},
					}, "3492229395", nil, nil, "", nil, nil}, ""},
					&ParentScope{"36"},
					&Assign{"37", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229412", nil, nil, "", nil, nil}, ""},
					&ParentScope{"37"},
					&Assign{"38", &ast.Literal{&types.Type{
						Kind: 10,
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229410", nil, nil, "", nil, nil}, ""},
					&ParentScope{"38"},
					&Assign{"39", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229407", nil, nil, "", nil, nil}, ""},
					&ParentScope{"39"},
					&Assign{"40", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Duration",
								Properties: map[string]*types.Type{
									"Hours": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Microseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Milliseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Minutes": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Nanoseconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"Seconds": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 6,
											},
										},
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
								},
							},
						},
					}, "3492229394", nil, nil, "", nil, nil}, ""},
					&ParentScope{"40"},
					&Assign{"41", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
							},
							&types.Type{
								Kind: 6,
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
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
					}, "3492229408", nil, nil, "", nil, nil}, ""},
					&ParentScope{"41"},
					&Assign{"42", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 1,
								Name: "Time",
								Properties: map[string]*types.Type{
									"Day": &types.Type{
										Kind: 6,
									},
									"Hour": &types.Type{
										Kind: 6,
									},
									"Minute": &types.Type{
										Kind: 6,
									},
									"Month": &types.Type{
										Kind: 6,
									},
									"Second": &types.Type{
										Kind: 6,
									},
									"String": &types.Type{
										Kind: 10,
										Returns: []*types.Type{
											&types.Type{
												Kind: 7,
											},
										},
									},
									"Year": &types.Type{
										Kind: 6,
									},
								},
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
					}, "3492229411", nil, nil, "", nil, nil}, ""},
					&ParentScope{"42"},
					&Assign{"43", &ast.Literal{&types.Type{
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
					}, "3492229406", nil, nil, "", nil, nil}, ""},
					&ParentScope{"43"},
					&Assign{"44", &ast.Literal{&types.Type{
						Kind: 10,
						Arguments: []*types.Type{
							&types.Type{
								Kind: 6,
							},
						},
						Returns: []*types.Type{
							&types.Type{
								Kind: 7,
							},
						},
					}, "3492229413", nil, nil, "", nil, nil}, ""},
					&ParentScope{"44"},
					&Assign{"1", &ast.Literal{&types.Type{
						Kind: 6,
					}, "4", nil, nil, "lib/time/time.ok:4:9", nil, nil}, ""},
					&Assign{"April", nil, "1"},
					&Assign{"2", &ast.Literal{&types.Type{
						Kind: 6,
					}, "8", nil, nil, "lib/time/time.ok:8:10", nil, nil}, ""},
					&Assign{"August", nil, "2"},
					&Assign{"3", &ast.Literal{&types.Type{
						Kind: 6,
					}, "12", nil, nil, "lib/time/time.ok:12:12", nil, nil}, ""},
					&Assign{"December", nil, "3"},
					&Assign{"4", &ast.Literal{&types.Type{
						Kind: 6,
					}, "2", nil, nil, "lib/time/time.ok:2:12", nil, nil}, ""},
					&Assign{"February", nil, "4"},
					&Assign{"5", &ast.Literal{&types.Type{
						Kind: 6,
					}, "3600", nil, nil, "lib/time/duration.ok:8:15", nil, nil}, ""},
					&Assign{"Hour", nil, "5"},
					&Assign{"6", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1", nil, nil, "lib/time/time.ok:1:11", nil, nil}, ""},
					&Assign{"January", nil, "6"},
					&Assign{"7", &ast.Literal{&types.Type{
						Kind: 6,
					}, "7", nil, nil, "lib/time/time.ok:7:8", nil, nil}, ""},
					&Assign{"July", nil, "7"},
					&Assign{"8", &ast.Literal{&types.Type{
						Kind: 6,
					}, "6", nil, nil, "lib/time/time.ok:6:8", nil, nil}, ""},
					&Assign{"June", nil, "8"},
					&Assign{"9", &ast.Literal{&types.Type{
						Kind: 6,
					}, "3", nil, nil, "lib/time/time.ok:3:9", nil, nil}, ""},
					&Assign{"March", nil, "9"},
					&Assign{"10", &ast.Literal{&types.Type{
						Kind: 6,
					}, "5", nil, nil, "lib/time/time.ok:5:7", nil, nil}, ""},
					&Assign{"May", nil, "10"},
					&Assign{"11", &ast.Literal{&types.Type{
						Kind: 6,
					}, "0.000001", nil, nil, "lib/time/duration.ok:4:15", nil, nil}, ""},
					&Assign{"Microsecond", nil, "11"},
					&Assign{"12", &ast.Literal{&types.Type{
						Kind: 6,
					}, "0.001", nil, nil, "lib/time/duration.ok:5:15", nil, nil}, ""},
					&Assign{"Millisecond", nil, "12"},
					&Assign{"13", &ast.Literal{&types.Type{
						Kind: 6,
					}, "60", nil, nil, "lib/time/duration.ok:7:15", nil, nil}, ""},
					&Assign{"Minute", nil, "13"},
					&Assign{"14", &ast.Literal{&types.Type{
						Kind: 6,
					}, "0.000000001", nil, nil, "lib/time/duration.ok:3:15", nil, nil}, ""},
					&Assign{"Nanosecond", nil, "14"},
					&Assign{"15", &ast.Literal{&types.Type{
						Kind: 6,
					}, "11", nil, nil, "lib/time/time.ok:11:12", nil, nil}, ""},
					&Assign{"November", nil, "15"},
					&Assign{"16", &ast.Literal{&types.Type{
						Kind: 6,
					}, "10", nil, nil, "lib/time/time.ok:10:11", nil, nil}, ""},
					&Assign{"October", nil, "16"},
					&Assign{"17", &ast.Literal{&types.Type{
						Kind: 6,
					}, "1", nil, nil, "lib/time/duration.ok:6:15", nil, nil}, ""},
					&Assign{"Second", nil, "17"},
					&Assign{"18", &ast.Literal{&types.Type{
						Kind: 6,
					}, "9", nil, nil, "lib/time/time.ok:9:13", nil, nil}, ""},
					&Assign{"September", nil, "18"},
					&Assign{"3492229393", nil, "19"},
					&Assign{"3492229394", nil, "20"},
					&Assign{"3492229395", nil, "21"},
					&Assign{"3492229396", nil, "22"},
					&Assign{"3492229397", nil, "23"},
					&Assign{"3492229398", nil, "24"},
					&Assign{"3492229406", nil, "25"},
					&Assign{"3492229407", nil, "26"},
					&Assign{"3492229408", nil, "27"},
					&Assign{"3492229410", nil, "28"},
					&Assign{"3492229411", nil, "29"},
					&Assign{"3492229412", nil, "30"},
					&Assign{"3492229413", nil, "31"},
					&Assign{"Add", nil, "32"},
					&Assign{"After", nil, "33"},
					&Assign{"Before", nil, "34"},
					&Assign{"Duration", nil, "35"},
					&Assign{"Equal", nil, "36"},
					&Assign{"FromUnix", nil, "37"},
					&Assign{"Now", nil, "38"},
					&Assign{"Sleep", nil, "39"},
					&Assign{"Sub", nil, "40"},
					&Assign{"Time", nil, "41"},
					&Assign{"Unix", nil, "42"},
					&Assign{"floor", nil, "43"},
					&Assign{"zeroPad", nil, "44"},
					&Return{Registers{"0"}},
				},
				Registers: 44,
				Type: &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 1,
							Name: "..__time",
							Properties: map[string]*types.Type{
								"Add": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 1,
											Name: "Duration",
											Properties: map[string]*types.Type{
												"Hours": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Microseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Milliseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Minutes": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Nanoseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Seconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"After": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"April": &types.Type{
									Kind: 6,
								},
								"August": &types.Type{
									Kind: 6,
								},
								"Before": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"December": &types.Type{
									Kind: 6,
								},
								"Duration": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Duration",
											Properties: map[string]*types.Type{
												"Hours": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Microseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Milliseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Minutes": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Nanoseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Seconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"String": &types.Type{
													Kind: 10,
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
								"Equal": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 3,
										},
									},
								},
								"February": &types.Type{
									Kind: 6,
								},
								"FromUnix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"Hour": &types.Type{
									Kind: 6,
								},
								"January": &types.Type{
									Kind: 6,
								},
								"July": &types.Type{
									Kind: 6,
								},
								"June": &types.Type{
									Kind: 6,
								},
								"March": &types.Type{
									Kind: 6,
								},
								"May": &types.Type{
									Kind: 6,
								},
								"Microsecond": &types.Type{
									Kind: 6,
								},
								"Millisecond": &types.Type{
									Kind: 6,
								},
								"Minute": &types.Type{
									Kind: 6,
								},
								"Nanosecond": &types.Type{
									Kind: 6,
								},
								"November": &types.Type{
									Kind: 6,
								},
								"Now": &types.Type{
									Kind: 10,
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"October": &types.Type{
									Kind: 6,
								},
								"Second": &types.Type{
									Kind: 6,
								},
								"September": &types.Type{
									Kind: 6,
								},
								"Sleep": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Duration",
											Properties: map[string]*types.Type{
												"Hours": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Microseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Milliseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Minutes": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Nanoseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Seconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"String": &types.Type{
													Kind: 10,
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
								"Sub": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Duration",
											Properties: map[string]*types.Type{
												"Hours": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Microseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Milliseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Minutes": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Nanoseconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"Seconds": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 6,
														},
													},
												},
												"String": &types.Type{
													Kind: 10,
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
								"Time": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
										},
										&types.Type{
											Kind: 6,
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
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
								},
								"Unix": &types.Type{
									Kind: 10,
									Arguments: []*types.Type{
										&types.Type{
											Kind: 1,
											Name: "Time",
											Properties: map[string]*types.Type{
												"Day": &types.Type{
													Kind: 6,
												},
												"Hour": &types.Type{
													Kind: 6,
												},
												"Minute": &types.Type{
													Kind: 6,
												},
												"Month": &types.Type{
													Kind: 6,
												},
												"Second": &types.Type{
													Kind: 6,
												},
												"String": &types.Type{
													Kind: 10,
													Returns: []*types.Type{
														&types.Type{
															Kind: 7,
														},
													},
												},
												"Year": &types.Type{
													Kind: 6,
												},
											},
										},
									},
									Returns: []*types.Type{
										&types.Type{
											Kind: 6,
										},
									},
								},
							},
						},
					},
				},
				Name: "..__time",
				Pos:  "lib/time/util.ok:1:1",
			},
		},
	}
}
