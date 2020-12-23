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
					&AssignFunc{"1", "func(string) Error", "1657625601"},
					&ParentScope{"1"},
					&AssignFunc{"2", "func(string) Error", "1657625601"},
					&ParentScope{"2"},
					&Assign{"1657625601", "1"},
					&Assign{"Error", "2"},
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
			Types:   nil,
			Symbols: nil,
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
						&AssignFunc{"2", "func(string)", "3624248086"},
						&ParentScope{"2"},
						&AssignFunc{"3", "func(string)", "3624248085"},
						&ParentScope{"3"},
						&AssignFunc{"4", "func(string)", "3624248084"},
						&ParentScope{"4"},
						&AssignFunc{"5", "func(string)", "3624248083"},
						&ParentScope{"5"},
						&AssignFunc{"6", "func(string)", "3624248082"},
						&ParentScope{"6"},
						&Assign{"Fatal", "2"},
						&Assign{"Error", "3"},
						&Assign{"Warn", "4"},
						&Assign{"Info", "5"},
						&Assign{"Debug", "6"},
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
						&AssignSymbol{"2", "stringDEBUG"},
						&Call{"*^Log", Registers{"2", "message"}, nil, "any"},
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
						&AssignSymbol{"2", "stringINFO"},
						&Call{"*^Log", Registers{"2", "message"}, nil, "any"},
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
						&AssignSymbol{"2", "stringWARN"},
						&Call{"*^Log", Registers{"2", "message"}, nil, "any"},
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
						&AssignSymbol{"2", "stringERROR"},
						&Call{"*^Log", Registers{"2", "message"}, nil, "any"},
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
						&AssignSymbol{"2", "stringFATAL"},
						&Call{"*^Log", Registers{"2", "message"}, nil, "any"},
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
						&AssignSymbol{"3", "string "},
						&AssignSymbol{"4", "number5"},
						&LoadPackage{"5", "strings"},
						&AssignSymbol{"6", "stringPadLeft"},
						&MapGet{"5", "6", "7"},
						&Call{"*7", Registers{"level", "3", "4"}, Registers{"8"}, "any"},
						&Assign{"level", "8"},
						&LoadPackage{"10", "time"},
						&AssignSymbol{"11", "stringNow"},
						&MapGet{"10", "11", "12"},
						&Call{"*12", nil, Registers{"13"}, "Time"},
						&AssignSymbol{"14", "stringString"},
						&MapGet{"13", "14", "15"},
						&Call{"*15", nil, Registers{"16"}, "any"},
						&AssignSymbol{"17", "string "},
						&AssignSymbol{"18", "string "},
						&Interpolate{"9", Registers{"16", "17", "level", "18", "message"}},
						&Print{Registers{"9"}},
						&AssignSymbol{"19", "stringFATAL"},
						&Equal{"level", "19", "20"},
						&JumpUnless{"20", 25},
						&AssignSymbol{"21", "number1"},
						&LoadPackage{"22", "runtime"},
						&AssignSymbol{"23", "stringExit"},
						&MapGet{"22", "23", "24"},
						&Call{"*24", Registers{"21"}, nil, "any"},
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
					&AssignFunc{"6", "func(func(string, string)) Logger", "3624248081"},
					&ParentScope{"6"},
					&AssignFunc{"7", "func(string, string)", "3624248087"},
					&ParentScope{"7"},
					&AssignFunc{"8", "func(string, string)", "3624248087"},
					&ParentScope{"8"},
					&AssignFunc{"9", "func(func(string, string)) Logger", "3624248081"},
					&ParentScope{"9"},
					&AssignSymbol{"1", "stringDEBUG"},
					&Assign{"LevelDebug", "1"},
					&AssignSymbol{"2", "stringERROR"},
					&Assign{"LevelError", "2"},
					&AssignSymbol{"3", "stringFATAL"},
					&Assign{"LevelFatal", "3"},
					&AssignSymbol{"4", "stringINFO"},
					&Assign{"LevelInfo", "4"},
					&AssignSymbol{"5", "stringWARN"},
					&Assign{"LevelWarn", "5"},
					&Assign{"3624248081", "6"},
					&Assign{"3624248087", "7"},
					&Assign{"Log", "8"},
					&Assign{"Logger", "9"},
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
			Types: map[TypeRegister]*types.Type{
				"Time": &types.Type{
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
				"any": &types.Type{
					Kind: 2,
				},
				"func(string)": &types.Type{
					Kind: 10,
					Arguments: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number5": &Symbol{
					Type:  "number",
					Value: "5",
				},
				"string ": &Symbol{
					Type:  "string",
					Value: " ",
				},
				"stringDEBUG": &Symbol{
					Type:  "string",
					Value: "DEBUG",
				},
				"stringERROR": &Symbol{
					Type:  "string",
					Value: "ERROR",
				},
				"stringExit": &Symbol{
					Type:  "string",
					Value: "Exit",
				},
				"stringFATAL": &Symbol{
					Type:  "string",
					Value: "FATAL",
				},
				"stringINFO": &Symbol{
					Type:  "string",
					Value: "INFO",
				},
				"stringNow": &Symbol{
					Type:  "string",
					Value: "Now",
				},
				"stringPadLeft": &Symbol{
					Type:  "string",
					Value: "PadLeft",
				},
				"stringString": &Symbol{
					Type:  "string",
					Value: "String",
				},
				"stringWARN": &Symbol{
					Type:  "string",
					Value: "WARN",
				},
			},
		},
		"math": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"1725484959": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "number0"},
						&LessThanNumber{"x", "2", "3"},
						&JumpUnless{"3", 5},
						&AssignSymbol{"4", "number0"},
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
						&AssignSymbol{"3", "number10"},
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
						&AssignSymbol{"2", "number2.71828182845904523536028747135266249775724709369995957496696763"},
						&Power{"2", "x", "3"},
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
					Pos:        "lib/math/powers.ok:7:1",
				},
				"1725484964": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "number0.5"},
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
					Pos:        "lib/math/powers.ok:12:1",
				},
				"1725484965": &CompiledFunc{
					Arguments: []string{"x"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "number1"},
						&AssignSymbol{"3", "number3"},
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
					Pos:        "lib/math/powers.ok:17:1",
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
						&AssignSymbol{"2", "number1"},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", "3"},
						&AssignSymbol{"4", "number0"},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&AssignSymbol{"6", "number0"},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 11},
						&Subtract{"x", "frac", "8"},
						&Return{Registers{"8"}},
						&AssignSymbol{"9", "number1"},
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
						&AssignSymbol{"2", "number1"},
						&Remainder{"x", "2", "3"},
						&Assign{"frac", "3"},
						&AssignSymbol{"4", "number0"},
						&EqualNumber{"frac", "4", "5"},
						&JumpUnless{"5", 6},
						&Return{Registers{"x"}},
						&AssignSymbol{"6", "number0"},
						&LessThanNumber{"x", "6", "7"},
						&JumpUnless{"7", 13},
						&AssignSymbol{"8", "number1"},
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
						&AssignSymbol{"3", "number10"},
						&Power{"3", "prec", "4"},
						&Assign{"p", "4"},
						&Multiply{"x", "p", "5"},
						&Assign{"y", "5"},
						&AssignSymbol{"6", "number1"},
						&Remainder{"y", "6", "7"},
						&Assign{"diff", "7"},
						&AssignSymbol{"8", "number0.5"},
						&GreaterThanEqualNumber{"diff", "8", "9"},
						&JumpUnless{"9", 15},
						&AssignSymbol{"10", "number1"},
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
					&AssignFunc{"10", "func(number) number", "1725484959"},
					&ParentScope{"10"},
					&AssignFunc{"11", "func(number) number", "1725484960"},
					&ParentScope{"11"},
					&AssignFunc{"12", "func(number) number", "1725484961"},
					&ParentScope{"12"},
					&AssignFunc{"13", "func(number) number", "1725484962"},
					&ParentScope{"13"},
					&AssignFunc{"14", "func(number, number) number", "1725484963"},
					&ParentScope{"14"},
					&AssignFunc{"15", "func(number) number", "1725484964"},
					&ParentScope{"15"},
					&AssignFunc{"16", "func(number) number", "1725484965"},
					&ParentScope{"16"},
					&AssignFunc{"17", "func() number", "1725484966"},
					&ParentScope{"17"},
					&AssignFunc{"18", "func(number) number", "1725484967"},
					&ParentScope{"18"},
					&AssignFunc{"19", "func(number) number", "1725484968"},
					&ParentScope{"19"},
					&AssignFunc{"20", "func(number, number) number", "1725484969"},
					&ParentScope{"20"},
					&AssignFunc{"21", "func(number) number", "1725484959"},
					&ParentScope{"21"},
					&AssignFunc{"22", "func(number) number", "1725484965"},
					&ParentScope{"22"},
					&AssignFunc{"23", "func(number) number", "1725484967"},
					&ParentScope{"23"},
					&AssignFunc{"24", "func(number) number", "1725484962"},
					&ParentScope{"24"},
					&AssignFunc{"25", "func(number) number", "1725484968"},
					&ParentScope{"25"},
					&AssignFunc{"26", "func(number) number", "1725484961"},
					&ParentScope{"26"},
					&AssignFunc{"27", "func(number) number", "1725484960"},
					&ParentScope{"27"},
					&AssignFunc{"28", "func(number, number) number", "1725484963"},
					&ParentScope{"28"},
					&AssignFunc{"29", "func() number", "1725484966"},
					&ParentScope{"29"},
					&AssignFunc{"30", "func(number, number) number", "1725484969"},
					&ParentScope{"30"},
					&AssignFunc{"31", "func(number) number", "1725484964"},
					&ParentScope{"31"},
					&AssignSymbol{"1", "number2.71828182845904523536028747135266249775724709369995957496696763"},
					&Assign{"E", "1"},
					&AssignSymbol{"2", "number2.30258509299404568401799145468436420760110148862877297603332790"},
					&Assign{"Ln10", "2"},
					&AssignSymbol{"3", "number0.693147180559945309417232121458176568075500134360255254120680009"},
					&Assign{"Ln2", "3"},
					&AssignSymbol{"4", "number1.61803398874989484820458683436563811772030917980576286213544862"},
					&Assign{"Phi", "4"},
					&AssignSymbol{"5", "number3.14159265358979323846264338327950288419716939937510582097494459"},
					&Assign{"Pi", "5"},
					&AssignSymbol{"6", "number1.41421356237309504880168872420969807856967187537694807317667974"},
					&Assign{"Sqrt2", "6"},
					&AssignSymbol{"7", "number1.64872127070012814684865078781416357165377610071014801157507931"},
					&Assign{"SqrtE", "7"},
					&AssignSymbol{"8", "number1.27201964951406896425242246173749149171560804184009624861664038"},
					&Assign{"SqrtPhi", "8"},
					&AssignSymbol{"9", "number1.77245385090551602729816748334114518279754945612238712821380779"},
					&Assign{"SqrtPi", "9"},
					&Assign{"1725484959", "10"},
					&Assign{"1725484960", "11"},
					&Assign{"1725484961", "12"},
					&Assign{"1725484962", "13"},
					&Assign{"1725484963", "14"},
					&Assign{"1725484964", "15"},
					&Assign{"1725484965", "16"},
					&Assign{"1725484966", "17"},
					&Assign{"1725484967", "18"},
					&Assign{"1725484968", "19"},
					&Assign{"1725484969", "20"},
					&Assign{"Abs", "21"},
					&Assign{"Cbrt", "22"},
					&Assign{"Ceil", "23"},
					&Assign{"Exp", "24"},
					&Assign{"Floor", "25"},
					&Assign{"Log10", "26"},
					&Assign{"LogE", "27"},
					&Assign{"Pow", "28"},
					&Assign{"Rand", "29"},
					&Assign{"Round", "30"},
					&Assign{"Sqrt", "31"},
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
			Types: nil,
			Symbols: map[SymbolRegister]*Symbol{
				"number0": &Symbol{
					Type:  "number",
					Value: "0",
				},
				"number0.5": &Symbol{
					Type:  "number",
					Value: "0.5",
				},
				"number0.693147180559945309417232121458176568075500134360255254120680009": &Symbol{
					Type:  "number",
					Value: "0.693147180559945309417232121458176568075500134360255254120680009",
				},
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number1.27201964951406896425242246173749149171560804184009624861664038": &Symbol{
					Type:  "number",
					Value: "1.27201964951406896425242246173749149171560804184009624861664038",
				},
				"number1.41421356237309504880168872420969807856967187537694807317667974": &Symbol{
					Type:  "number",
					Value: "1.41421356237309504880168872420969807856967187537694807317667974",
				},
				"number1.61803398874989484820458683436563811772030917980576286213544862": &Symbol{
					Type:  "number",
					Value: "1.61803398874989484820458683436563811772030917980576286213544862",
				},
				"number1.64872127070012814684865078781416357165377610071014801157507931": &Symbol{
					Type:  "number",
					Value: "1.64872127070012814684865078781416357165377610071014801157507931",
				},
				"number1.77245385090551602729816748334114518279754945612238712821380779": &Symbol{
					Type:  "number",
					Value: "1.77245385090551602729816748334114518279754945612238712821380779",
				},
				"number10": &Symbol{
					Type:  "number",
					Value: "10",
				},
				"number2.30258509299404568401799145468436420760110148862877297603332790": &Symbol{
					Type:  "number",
					Value: "2.30258509299404568401799145468436420760110148862877297603332790",
				},
				"number2.71828182845904523536028747135266249775724709369995957496696763": &Symbol{
					Type:  "number",
					Value: "2.71828182845904523536028747135266249775724709369995957496696763",
				},
				"number3": &Symbol{
					Type:  "number",
					Value: "3",
				},
				"number3.14159265358979323846264338327950288419716939937510582097494459": &Symbol{
					Type:  "number",
					Value: "3.14159265358979323846264338327950288419716939937510582097494459",
				},
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
						&AssignFunc{"2", "func() string", "3819329021"},
						&ParentScope{"2"},
						&AssignFunc{"3", "func(number) string", "3819329020"},
						&ParentScope{"3"},
						&AssignFunc{"4", "func(number) data", "3819329019"},
						&ParentScope{"4"},
						&AssignFunc{"5", "func(number, number) number", "3819329018"},
						&ParentScope{"5"},
						&AssignFunc{"6", "func()", "3819329017"},
						&ParentScope{"6"},
						&AssignFunc{"7", "func(data)", "3819329016"},
						&ParentScope{"7"},
						&AssignFunc{"8", "func(string)", "3819329015"},
						&ParentScope{"8"},
						&Assign{"ReadLine", "2"},
						&Assign{"ReadString", "3"},
						&Assign{"ReadData", "4"},
						&Assign{"Seek", "5"},
						&Assign{"Close", "6"},
						&Assign{"WriteData", "7"},
						&Assign{"WriteString", "8"},
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
						&AssignSymbol{"1", "string"},
						&Assign{"line", "1"},
						&AssignSymbol{"2", "booltrue"},
						&JumpUnless{"2", 17},
						&AssignSymbol{"3", "number1"},
						&ReadString{"^fd", "3", "4"},
						&Assign{"chars", "4"},
						&Len{"chars", "5"},
						&AssignSymbol{"6", "number0"},
						&EqualNumber{"5", "6", "7"},
						&JumpUnless{"7", 11},
						&Jump{17},
						&Concat{"line", "chars", "line"},
						&AssignSymbol{"8", "string\\n"},
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
						&AssignSymbol{"3", "func(data) File3819329014"},
						&Call{"*3", Registers{"2"}, Registers{"4"}, "File"},
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
						&Assign{"name", "2"},
						&Assign{"size", "3"},
						&Assign{"mode", "4"},
						&Assign{"modTimeYear", "5"},
						&Assign{"modTimeMonth", "6"},
						&Assign{"modTimeDay", "7"},
						&Assign{"modTimeHour", "8"},
						&Assign{"modTimeMinute", "9"},
						&Assign{"modTimeSecond", "10"},
						&Assign{"isDir", "11"},
						&LoadPackage{"12", "time"},
						&AssignSymbol{"13", "stringTime"},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"modTimeYear", "modTimeMonth", "modTimeDay", "modTimeHour", "modTimeMinute", "modTimeSecond"}, Registers{"15"}, "Time"},
						&Assign{"modTime", "15"},
						&AssignSymbol{"16", "func(string, number, string, Time, bool) FileInfo3819329026"},
						&Call{"*16", Registers{"name", "size", "mode", "modTime", "isDir"}, Registers{"17"}, "FileInfo"},
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
						&AssignSymbol{"1", "string/tmp/ok."},
						&Rand{"2"},
						&AssignSymbol{"3", "number1000000000"},
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
					&AssignFunc{"1", "func(data) File", "3819329014"},
					&ParentScope{"1"},
					&AssignFunc{"2", "func(string) File", "3819329022"},
					&ParentScope{"2"},
					&AssignFunc{"3", "func(string)", "3819329023"},
					&ParentScope{"3"},
					&AssignFunc{"4", "func(string, string)", "3819329024"},
					&ParentScope{"4"},
					&AssignFunc{"5", "func(string)", "3819329025"},
					&ParentScope{"5"},
					&AssignFunc{"6", "func(string, number, string, Time, bool) FileInfo", "3819329026"},
					&ParentScope{"6"},
					&AssignFunc{"7", "func(string) FileInfo", "3819329027"},
					&ParentScope{"7"},
					&AssignFunc{"8", "func() string", "3819329028"},
					&ParentScope{"8"},
					&AssignFunc{"9", "func(string)", "3819329025"},
					&ParentScope{"9"},
					&AssignFunc{"10", "func(data) File", "3819329014"},
					&ParentScope{"10"},
					&AssignFunc{"11", "func(string, number, string, Time, bool) FileInfo", "3819329026"},
					&ParentScope{"11"},
					&AssignFunc{"12", "func(string) FileInfo", "3819329027"},
					&ParentScope{"12"},
					&AssignFunc{"13", "func(string) File", "3819329022"},
					&ParentScope{"13"},
					&AssignFunc{"14", "func(string)", "3819329023"},
					&ParentScope{"14"},
					&AssignFunc{"15", "func(string, string)", "3819329024"},
					&ParentScope{"15"},
					&AssignFunc{"16", "func() string", "3819329028"},
					&ParentScope{"16"},
					&Assign{"3819329014", "1"},
					&Assign{"3819329022", "2"},
					&Assign{"3819329023", "3"},
					&Assign{"3819329024", "4"},
					&Assign{"3819329025", "5"},
					&Assign{"3819329026", "6"},
					&Assign{"3819329027", "7"},
					&Assign{"3819329028", "8"},
					&Assign{"CreateDirectory", "9"},
					&Assign{"File", "10"},
					&Assign{"FileInfo", "11"},
					&Assign{"Info", "12"},
					&Assign{"Open", "13"},
					&Assign{"Remove", "14"},
					&Assign{"Rename", "15"},
					&Assign{"TempPath", "16"},
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
			Types: map[TypeRegister]*types.Type{
				"File": &types.Type{
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
				"FileInfo": &types.Type{
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
				"Time": &types.Type{
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
				"func()": &types.Type{
					Kind: 10,
				},
				"func() string": &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
				},
				"func(data)": &types.Type{
					Kind: 10,
					Arguments: []*types.Type{
						&types.Type{
							Kind: 5,
						},
					},
				},
				"func(number) data": &types.Type{
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
				"func(number) string": &types.Type{
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
				"func(number, number) number": &types.Type{
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
				"func(string)": &types.Type{
					Kind: 10,
					Arguments: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"booltrue": &Symbol{
					Type:  "bool",
					Value: "true",
				},
				"func(data) File3819329014": &Symbol{
					Type:  "func(data) File",
					Value: "3819329014",
				},
				"func(string, number, string, Time, bool) FileInfo3819329026": &Symbol{
					Type:  "func(string, number, string, Time, bool) FileInfo",
					Value: "3819329026",
				},
				"number0": &Symbol{
					Type:  "number",
					Value: "0",
				},
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number1000000000": &Symbol{
					Type:  "number",
					Value: "1000000000",
				},
				"string": &Symbol{
					Type: "string",
				},
				"string/tmp/ok.": &Symbol{
					Type:  "string",
					Value: "/tmp/ok.",
				},
				"stringTime": &Symbol{
					Type:  "string",
					Value: "Time",
				},
				"string\n": &Symbol{
					Type:  "string",
					Value: "\n",
				},
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
						&AssignSymbol{"2", "func(any) string1100080410"},
						&Call{"*2", Registers{"value"}, Registers{"3"}, "any"},
						&Assign{"type", "3"},
						&AssignSymbol{"4", "string[]"},
						&LoadPackage{"5", "strings"},
						&AssignSymbol{"6", "stringHasPrefix"},
						&MapGet{"5", "6", "7"},
						&Call{"*7", Registers{"type", "4"}, Registers{"8"}, "any"},
						&JumpUnless{"8", 11},
						&AssignSymbol{"9", "stringarray"},
						&Return{Registers{"9"}},
						&Jump{30},
						&AssignSymbol{"11", "string{}"},
						&Interpolate{"10", Registers{"11"}},
						&LoadPackage{"12", "strings"},
						&AssignSymbol{"13", "stringHasPrefix"},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"type", "10"}, Registers{"15"}, "any"},
						&JumpUnless{"15", 21},
						&AssignSymbol{"16", "stringmap"},
						&Return{Registers{"16"}},
						&Jump{30},
						&AssignSymbol{"17", "stringfunc("},
						&LoadPackage{"18", "strings"},
						&AssignSymbol{"19", "stringHasPrefix"},
						&MapGet{"18", "19", "20"},
						&Call{"*20", Registers{"type", "17"}, Registers{"21"}, "any"},
						&JumpUnless{"21", 30},
						&AssignSymbol{"22", "stringfunc"},
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
					&AssignFunc{"1", "func(any, []any) []any", "1100080403"},
					&ParentScope{"1"},
					&AssignFunc{"2", "func(any, any) any", "1100080404"},
					&ParentScope{"2"},
					&AssignFunc{"3", "func(any) string", "1100080405"},
					&ParentScope{"3"},
					&AssignFunc{"4", "func(any) string", "1100080406"},
					&ParentScope{"4"},
					&AssignFunc{"5", "func(any) number", "1100080407"},
					&ParentScope{"5"},
					&AssignFunc{"6", "func(any) []string", "1100080408"},
					&ParentScope{"6"},
					&AssignFunc{"7", "func(any, any, any) any", "1100080409"},
					&ParentScope{"7"},
					&AssignFunc{"8", "func(any) string", "1100080410"},
					&ParentScope{"8"},
					&AssignFunc{"9", "func(any, []any) []any", "1100080403"},
					&ParentScope{"9"},
					&AssignFunc{"10", "func(any, any) any", "1100080404"},
					&ParentScope{"10"},
					&AssignFunc{"11", "func(any) string", "1100080405"},
					&ParentScope{"11"},
					&AssignFunc{"12", "func(any) string", "1100080406"},
					&ParentScope{"12"},
					&AssignFunc{"13", "func(any) number", "1100080407"},
					&ParentScope{"13"},
					&AssignFunc{"14", "func(any) []string", "1100080408"},
					&ParentScope{"14"},
					&AssignFunc{"15", "func(any, any, any) any", "1100080409"},
					&ParentScope{"15"},
					&AssignFunc{"16", "func(any) string", "1100080410"},
					&ParentScope{"16"},
					&Assign{"1100080403", "1"},
					&Assign{"1100080404", "2"},
					&Assign{"1100080405", "3"},
					&Assign{"1100080406", "4"},
					&Assign{"1100080407", "5"},
					&Assign{"1100080408", "6"},
					&Assign{"1100080409", "7"},
					&Assign{"1100080410", "8"},
					&Assign{"Call", "9"},
					&Assign{"Get", "10"},
					&Assign{"Interface", "11"},
					&Assign{"Kind", "12"},
					&Assign{"Len", "13"},
					&Assign{"Properties", "14"},
					&Assign{"Set", "15"},
					&Assign{"Type", "16"},
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
			Types: map[TypeRegister]*types.Type{
				"any": &types.Type{
					Kind: 2,
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"func(any) string1100080410": &Symbol{
					Type:  "func(any) string",
					Value: "1100080410",
				},
				"stringHasPrefix": &Symbol{
					Type:  "string",
					Value: "HasPrefix",
				},
				"string[]": &Symbol{
					Type:  "string",
					Value: "[]",
				},
				"stringarray": &Symbol{
					Type:  "string",
					Value: "array",
				},
				"stringfunc": &Symbol{
					Type:  "string",
					Value: "func",
				},
				"stringfunc(": &Symbol{
					Type:  "string",
					Value: "func(",
				},
				"stringmap": &Symbol{
					Type:  "string",
					Value: "map",
				},
				"string{}": &Symbol{
					Type:  "string",
					Value: "{}",
				},
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
						&AssignSymbol{"2", "func(string) (string, bool)3243469103"},
						&Call{"*2", Registers{"name"}, Registers{"3", "4"}, "any"},
						&Assign{"value", "3"},
						&Assign{"_", "4"},
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
						&AssignSymbol{"2", "string\\n"},
						&LoadPackage{"3", "strings"},
						&AssignSymbol{"4", "stringSplit"},
						&MapGet{"3", "4", "5"},
						&Call{"*5", Registers{"1", "2"}, Registers{"6"}, "any"},
						&Assign{"elements", "6"},
						&AssignSymbol{"7", "number0"},
						&ArrayAlloc{"7", "8", "[]StackElement"},
						&Assign{"stack", "8"},
						&AssignSymbol{"9", "number0"},
						&NextArray{"elements", "9", "", "element", "10"},
						&JumpUnless{"10", 44},
						&AssignSymbol{"11", "string|"},
						&LoadPackage{"12", "strings"},
						&AssignSymbol{"13", "stringSplit"},
						&MapGet{"12", "13", "14"},
						&Call{"*14", Registers{"element", "11"}, Registers{"15"}, "any"},
						&Assign{"parts", "15"},
						&AssignSymbol{"16", "number0"},
						&ArrayGet{"parts", "16", "17"},
						&AssignSymbol{"18", "string:"},
						&LoadPackage{"19", "strings"},
						&AssignSymbol{"20", "stringSplit"},
						&MapGet{"19", "20", "21"},
						&Call{"*21", Registers{"17", "18"}, Registers{"22"}, "any"},
						&Assign{"locationParts", "22"},
						&AssignSymbol{"23", "number1"},
						&ArrayAlloc{"23", "24", "[]StackElement"},
						&AssignSymbol{"25", "number0"},
						&AssignSymbol{"26", "number0"},
						&ArrayGet{"locationParts", "26", "27"},
						&AssignSymbol{"28", "number1"},
						&ArrayGet{"locationParts", "28", "29"},
						&CastNumber{"29", "30"},
						&AssignSymbol{"31", "number2"},
						&ArrayGet{"locationParts", "31", "32"},
						&CastNumber{"32", "33"},
						&AssignSymbol{"34", "number1"},
						&ArrayGet{"parts", "34", "35"},
						&AssignSymbol{"36", "func(string, number, number, string) StackElement3243469107"},
						&Call{"*36", Registers{"27", "30", "33", "35"}, Registers{"37"}, "StackElement"},
						&ArraySet{"24", "25", "37"},
						&Append{"stack", "24", "stack"},
						&Jump{10},
						&AssignSymbol{"38", "func([]StackElement) StackTrace3243469109"},
						&Call{"*38", Registers{"stack"}, Registers{"39"}, "StackTrace"},
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
						&AssignFunc{"2", "func() string", "3243469110"},
						&ParentScope{"2"},
						&Assign{"String", "2"},
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
						&AssignSymbol{"1", "string"},
						&Assign{"s", "1"},
						&AssignSymbol{"2", "number1"},
						&Assign{"i", "2"},
						&AssignSymbol{"3", "number0"},
						&NextArray{"^Elements", "3", "", "element", "4"},
						&JumpUnless{"4", 21},
						&AssignSymbol{"6", "string "},
						&AssignSymbol{"7", "stringFunctionName"},
						&MapGet{"element", "7", "8"},
						&AssignSymbol{"9", "string "},
						&AssignSymbol{"10", "stringFile"},
						&MapGet{"element", "10", "11"},
						&AssignSymbol{"12", "string:"},
						&AssignSymbol{"13", "stringLineNumber"},
						&MapGet{"element", "13", "14"},
						&AssignSymbol{"15", "string\\n"},
						&Interpolate{"5", Registers{"i", "6", "8", "9", "11", "12", "14", "15"}},
						&Concat{"s", "5", "s"},
						&AssignSymbol{"16", "number1"},
						&Add{"i", "16", "i"},
						&Jump{4},
						&AssignSymbol{"17", "string\\n"},
						&LoadPackage{"18", "strings"},
						&AssignSymbol{"19", "stringTrimRight"},
						&MapGet{"18", "19", "20"},
						&Call{"*20", Registers{"s", "17"}, Registers{"21"}, "any"},
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
					&AssignFunc{"1", "func(string) string", "3243469102"},
					&ParentScope{"1"},
					&AssignFunc{"2", "func(string) (string, bool)", "3243469103"},
					&ParentScope{"2"},
					&AssignFunc{"3", "func(string, string)", "3243469104"},
					&ParentScope{"3"},
					&AssignFunc{"4", "func(string)", "3243469105"},
					&ParentScope{"4"},
					&AssignFunc{"5", "func(number)", "3243469106"},
					&ParentScope{"5"},
					&AssignFunc{"6", "func(string, number, number, string) StackElement", "3243469107"},
					&ParentScope{"6"},
					&AssignFunc{"7", "func() StackTrace", "3243469108"},
					&ParentScope{"7"},
					&AssignFunc{"8", "func([]StackElement) StackTrace", "3243469109"},
					&ParentScope{"8"},
					&AssignFunc{"9", "func(string) string", "3243469102"},
					&ParentScope{"9"},
					&AssignFunc{"10", "func(number)", "3243469106"},
					&ParentScope{"10"},
					&AssignFunc{"11", "func(string) (string, bool)", "3243469103"},
					&ParentScope{"11"},
					&AssignFunc{"12", "func(string, string)", "3243469104"},
					&ParentScope{"12"},
					&AssignFunc{"13", "func() StackTrace", "3243469108"},
					&ParentScope{"13"},
					&AssignFunc{"14", "func(string, number, number, string) StackElement", "3243469107"},
					&ParentScope{"14"},
					&AssignFunc{"15", "func([]StackElement) StackTrace", "3243469109"},
					&ParentScope{"15"},
					&AssignFunc{"16", "func(string)", "3243469105"},
					&ParentScope{"16"},
					&Assign{"3243469102", "1"},
					&Assign{"3243469103", "2"},
					&Assign{"3243469104", "3"},
					&Assign{"3243469105", "4"},
					&Assign{"3243469106", "5"},
					&Assign{"3243469107", "6"},
					&Assign{"3243469108", "7"},
					&Assign{"3243469109", "8"},
					&Assign{"Env", "9"},
					&Assign{"Exit", "10"},
					&Assign{"LookupEnv", "11"},
					&Assign{"SetEnv", "12"},
					&Assign{"Stack", "13"},
					&Assign{"StackElement", "14"},
					&Assign{"StackTrace", "15"},
					&Assign{"UnsetEnv", "16"},
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
			Types: map[TypeRegister]*types.Type{
				"StackElement": &types.Type{
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
				"StackTrace": &types.Type{
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
				"[]StackElement": &types.Type{
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
				},
				"any": &types.Type{
					Kind: 2,
				},
				"func() string": &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"func([]StackElement) StackTrace3243469109": &Symbol{
					Type:  "func([]StackElement) StackTrace",
					Value: "3243469109",
				},
				"func(string) (string, bool)3243469103": &Symbol{
					Type:  "func(string) (string, bool)",
					Value: "3243469103",
				},
				"func(string, number, number, string) StackElement3243469107": &Symbol{
					Type:  "func(string, number, number, string) StackElement",
					Value: "3243469107",
				},
				"number0": &Symbol{
					Type:  "number",
					Value: "0",
				},
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number2": &Symbol{
					Type:  "number",
					Value: "2",
				},
				"string": &Symbol{
					Type: "string",
				},
				"string ": &Symbol{
					Type:  "string",
					Value: " ",
				},
				"string:": &Symbol{
					Type:  "string",
					Value: ":",
				},
				"stringFile": &Symbol{
					Type:  "string",
					Value: "File",
				},
				"stringFunctionName": &Symbol{
					Type:  "string",
					Value: "FunctionName",
				},
				"stringLineNumber": &Symbol{
					Type:  "string",
					Value: "LineNumber",
				},
				"stringSplit": &Symbol{
					Type:  "string",
					Value: "Split",
				},
				"stringTrimRight": &Symbol{
					Type:  "string",
					Value: "TrimRight",
				},
				"string\n": &Symbol{
					Type:  "string",
					Value: "\n",
				},
				"string|": &Symbol{
					Type:  "string",
					Value: "|",
				},
			},
		},
		"strings": &File{
			Imports: nil,
			Funcs: map[string]*CompiledFunc{
				"601598823": &CompiledFunc{
					Arguments: []string{"s"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "string"},
						&Assign{"result", "2"},
						&AssignSymbol{"3", "number0"},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", "5"},
						&AssignSymbol{"6", "charA"},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&AssignSymbol{"9", "charZ"},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&AssignSymbol{"13", "number32"},
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
						&AssignSymbol{"2", "string"},
						&Assign{"result", "2"},
						&AssignSymbol{"3", "number0"},
						&NextString{"s", "3", "", "c", "4"},
						&JumpUnless{"4", 23},
						&CastNumber{"c", "5"},
						&Assign{"n", "5"},
						&AssignSymbol{"6", "chara"},
						&CastNumber{"6", "7"},
						&GreaterThanEqualNumber{"n", "7", "8"},
						&AssignSymbol{"9", "charz"},
						&CastNumber{"9", "10"},
						&LessThanEqualNumber{"n", "10", "11"},
						&And{"8", "11", "12"},
						&JumpUnless{"12", 20},
						&AssignSymbol{"13", "number32"},
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
						&AssignSymbol{"3", "func(string, string) number601598828"},
						&Call{"*3", Registers{"s", "substr"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "number-1"},
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
						&AssignSymbol{"6", "boolfalse"},
						&Return{Registers{"6"}},
						&AssignSymbol{"7", "number0"},
						&Assign{"i", "7"},
						&Len{"prefix", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 19},
						&StringIndex{"s", "i", "10"},
						&StringIndex{"prefix", "i", "11"},
						&NotEqual{"10", "11", "12"},
						&JumpUnless{"12", 16},
						&AssignSymbol{"13", "boolfalse"},
						&Return{Registers{"13"}},
						&AssignSymbol{"14", "number1"},
						&Add{"i", "14", "i"},
						&Jump{8},
						&AssignSymbol{"15", "booltrue"},
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
						&AssignSymbol{"6", "boolfalse"},
						&Return{Registers{"6"}},
						&Len{"s", "7"},
						&AssignSymbol{"8", "number1"},
						&Subtract{"7", "8", "9"},
						&Assign{"j", "9"},
						&Len{"suffix", "10"},
						&AssignSymbol{"11", "number1"},
						&Subtract{"10", "11", "12"},
						&Assign{"i", "12"},
						&AssignSymbol{"13", "number0"},
						&GreaterThanEqualNumber{"i", "13", "14"},
						&JumpUnless{"14", 27},
						&StringIndex{"s", "j", "15"},
						&StringIndex{"suffix", "i", "16"},
						&NotEqual{"15", "16", "17"},
						&JumpUnless{"17", 22},
						&AssignSymbol{"18", "boolfalse"},
						&Return{Registers{"18"}},
						&AssignSymbol{"19", "number1"},
						&Subtract{"j", "19", "j"},
						&AssignSymbol{"20", "number1"},
						&Subtract{"i", "20", "i"},
						&Jump{14},
						&AssignSymbol{"21", "booltrue"},
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
						&AssignSymbol{"3", "number-1"},
						&AssignSymbol{"4", "func(string, string, number) number601598829"},
						&Call{"*4", Registers{"s", "substr", "3"}, Registers{"5"}, "any"},
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
						&AssignSymbol{"4", "number-1"},
						&AssignSymbol{"5", "func(number, number) number601598831"},
						&Call{"*5", Registers{"offset", "4"}, Registers{"6"}, "any"},
						&Assign{"offset", "6"},
						&AssignSymbol{"7", "number1"},
						&Add{"offset", "7", "8"},
						&Assign{"i", "8"},
						&Len{"s", "9"},
						&Len{"substr", "10"},
						&Subtract{"9", "10", "11"},
						&LessThanEqualNumber{"i", "11", "12"},
						&JumpUnless{"12", 34},
						&AssignSymbol{"13", "booltrue"},
						&Assign{"found", "13"},
						&AssignSymbol{"14", "number0"},
						&Assign{"j", "14"},
						&Len{"substr", "15"},
						&LessThanNumber{"j", "15", "16"},
						&JumpUnless{"16", 29},
						&Add{"i", "j", "17"},
						&StringIndex{"s", "17", "18"},
						&StringIndex{"substr", "j", "19"},
						&NotEqual{"18", "19", "20"},
						&JumpUnless{"20", 26},
						&AssignSymbol{"21", "boolfalse"},
						&Assign{"found", "21"},
						&Jump{29},
						&AssignSymbol{"22", "number1"},
						&Add{"j", "22", "j"},
						&Jump{16},
						&JumpUnless{"found", 31},
						&Return{Registers{"i"}},
						&AssignSymbol{"23", "number1"},
						&Add{"i", "23", "i"},
						&Jump{9},
						&AssignSymbol{"24", "number-1"},
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
						&AssignSymbol{"3", "func(string) string601598840"},
						&Call{"*3", Registers{"s"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "func(string) string601598840"},
						&Call{"*5", Registers{"substr"}, Registers{"6"}, "any"},
						&AssignSymbol{"7", "func(string, string) number601598828"},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, "any"},
						&Assign{"index", "8"},
						&AssignSymbol{"9", "number-1"},
						&EqualNumber{"index", "9", "10"},
						&JumpUnless{"10", 11},
						&AssignSymbol{"11", "number-1"},
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
						&AssignSymbol{"6", "func(number, number) number601598830"},
						&Call{"*6", Registers{"offset", "5"}, Registers{"7"}, "any"},
						&AssignSymbol{"8", "number1"},
						&Add{"7", "8", "9"},
						&Subtract{"4", "9", "10"},
						&Assign{"offset", "10"},
						&AssignSymbol{"11", "func(string) string601598840"},
						&Call{"*11", Registers{"s"}, Registers{"12"}, "any"},
						&AssignSymbol{"13", "func(string) string601598840"},
						&Call{"*13", Registers{"substr"}, Registers{"14"}, "any"},
						&AssignSymbol{"15", "func(string, string, number) number601598829"},
						&Call{"*15", Registers{"12", "14", "offset"}, Registers{"16"}, "any"},
						&Assign{"index", "16"},
						&AssignSymbol{"17", "number-1"},
						&EqualNumber{"index", "17", "18"},
						&JumpUnless{"18", 19},
						&AssignSymbol{"19", "number-1"},
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
						&AssignSymbol{"3", "string"},
						&Assign{"result", "3"},
						&AssignSymbol{"4", "number0"},
						&NextArray{"strings", "4", "i", "s", "5"},
						&JumpUnless{"5", 10},
						&AssignSymbol{"6", "number0"},
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
						&AssignSymbol{"6", "string"},
						&Equal{"pad", "6", "7"},
						&Or{"5", "7", "8"},
						&JumpUnless{"8", 6},
						&Return{Registers{"s"}},
						&Len{"s", "9"},
						&Subtract{"toLen", "9", "10"},
						&AssignSymbol{"11", "func(string, number) string601598837"},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, "any"},
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
						&AssignSymbol{"6", "string"},
						&Equal{"pad", "6", "7"},
						&Or{"5", "7", "8"},
						&JumpUnless{"8", 6},
						&Return{Registers{"s"}},
						&Len{"s", "9"},
						&Subtract{"toLen", "9", "10"},
						&AssignSymbol{"11", "func(string, number) string601598837"},
						&Call{"*11", Registers{"pad", "10"}, Registers{"12"}, "any"},
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
						&AssignSymbol{"5", "func(string, number) string601598838"},
						&Call{"*5", Registers{"pad", "4"}, Registers{"6"}, "any"},
						&AssignSymbol{"7", "number0"},
						&AssignSymbol{"8", "func(string, number, number) string601598842"},
						&Call{"*8", Registers{"6", "7", "toLen"}, Registers{"9"}, "any"},
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
						&AssignSymbol{"3", "string"},
						&Assign{"result", "3"},
						&AssignSymbol{"4", "number0"},
						&Assign{"i", "4"},
						&LessThanNumber{"i", "times", "5"},
						&JumpUnless{"5", 9},
						&Concat{"result", "str", "result"},
						&AssignSymbol{"6", "number1"},
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
						&AssignSymbol{"4", "func(string, string) []string601598841"},
						&Call{"*4", Registers{"s", "find"}, Registers{"5"}, "any"},
						&AssignSymbol{"6", "func([]string, string) string601598834"},
						&Call{"*6", Registers{"5", "replace"}, Registers{"7"}, "any"},
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
						&AssignSymbol{"2", "string"},
						&Assign{"result", "2"},
						&Len{"s", "3"},
						&AssignSymbol{"4", "number1"},
						&Subtract{"3", "4", "5"},
						&Assign{"i", "5"},
						&AssignSymbol{"6", "number0"},
						&GreaterThanEqualNumber{"i", "6", "7"},
						&JumpUnless{"7", 14},
						&StringIndex{"s", "i", "8"},
						&CastString{"8", "9"},
						&Concat{"result", "9", "result"},
						&AssignSymbol{"10", "number1"},
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
						&AssignSymbol{"3", "number0"},
						&ArrayAlloc{"3", "4", "[]string"},
						&Assign{"elements", "4"},
						&AssignSymbol{"5", "string"},
						&Equal{"delimiter", "5", "6"},
						&JumpUnless{"6", 21},
						&AssignSymbol{"7", "number0"},
						&Assign{"i", "7"},
						&Len{"s", "8"},
						&LessThanNumber{"i", "8", "9"},
						&JumpUnless{"9", 20},
						&AssignSymbol{"10", "number1"},
						&ArrayAlloc{"10", "11", "[]string"},
						&AssignSymbol{"12", "number0"},
						&StringIndex{"s", "i", "13"},
						&CastString{"13", "14"},
						&ArraySet{"11", "12", "14"},
						&Append{"elements", "11", "elements"},
						&AssignSymbol{"15", "number1"},
						&Add{"i", "15", "i"},
						&Jump{8},
						&Jump{59},
						&AssignSymbol{"16", "string"},
						&Assign{"element", "16"},
						&AssignSymbol{"17", "number0"},
						&Assign{"i", "17"},
						&Len{"s", "18"},
						&LessThanNumber{"i", "18", "19"},
						&JumpUnless{"19", 54},
						&AssignSymbol{"20", "number1"},
						&Subtract{"i", "20", "21"},
						&AssignSymbol{"22", "func(string, string, number) number601598829"},
						&Call{"*22", Registers{"s", "delimiter", "21"}, Registers{"23"}, "any"},
						&Subtract{"23", "i", "24"},
						&AssignSymbol{"25", "number0"},
						&EqualNumber{"24", "25", "26"},
						&JumpUnless{"26", 48},
						&AssignSymbol{"27", "number1"},
						&ArrayAlloc{"27", "28", "[]string"},
						&AssignSymbol{"29", "number0"},
						&ArraySet{"28", "29", "element"},
						&Append{"elements", "28", "elements"},
						&AssignSymbol{"30", "string"},
						&Assign{"element", "30"},
						&Len{"delimiter", "31"},
						&AssignSymbol{"32", "number1"},
						&Subtract{"31", "32", "33"},
						&Add{"i", "33", "i"},
						&Jump{51},
						&StringIndex{"s", "i", "34"},
						&CastString{"34", "35"},
						&Concat{"element", "35", "element"},
						&AssignSymbol{"36", "number1"},
						&Add{"i", "36", "i"},
						&Jump{26},
						&AssignSymbol{"37", "number1"},
						&ArrayAlloc{"37", "38", "[]string"},
						&AssignSymbol{"39", "number0"},
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
						&AssignSymbol{"4", "string"},
						&Assign{"r", "4"},
						&Assign{"i", "fromIndex"},
						&LessThanNumber{"i", "toIndex", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "i", "6"},
						&CastString{"6", "7"},
						&Concat{"r", "7", "r"},
						&AssignSymbol{"8", "number1"},
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
						&AssignSymbol{"3", "number0"},
						&Assign{"offset", "3"},
						&Len{"s", "4"},
						&LessThanNumber{"offset", "4", "5"},
						&JumpUnless{"5", 17},
						&StringIndex{"s", "offset", "6"},
						&CastString{"6", "7"},
						&AssignSymbol{"8", "func(string, string) number601598828"},
						&Call{"*8", Registers{"cutset", "7"}, Registers{"9"}, "any"},
						&AssignSymbol{"10", "number-1"},
						&EqualNumber{"9", "10", "11"},
						&JumpUnless{"11", 14},
						&AssignSymbol{"12", "func(string, number) string601598848"},
						&Call{"*12", Registers{"s", "offset"}, Registers{"13"}, "any"},
						&Return{Registers{"13"}},
						&AssignSymbol{"14", "number1"},
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
						&AssignSymbol{"3", "func(string) string601598840"},
						&Call{"*3", Registers{"s"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "func(string, string) string601598843"},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, "any"},
						&AssignSymbol{"7", "func(string) string601598840"},
						&Call{"*7", Registers{"6"}, Registers{"8"}, "any"},
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
						&AssignSymbol{"3", "func(string, string) string601598843"},
						&Call{"*3", Registers{"s", "cutset"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "func(string, string) string601598844"},
						&Call{"*5", Registers{"4", "cutset"}, Registers{"6"}, "any"},
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
						&AssignSymbol{"3", "func(string, string) bool601598826"},
						&Call{"*3", Registers{"s", "prefix"}, Registers{"4"}, "any"},
						&JumpUnless{"4", 6},
						&Len{"prefix", "5"},
						&AssignSymbol{"6", "func(string, number) string601598848"},
						&Call{"*6", Registers{"s", "5"}, Registers{"7"}, "any"},
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
						&AssignSymbol{"3", "func(string) string601598840"},
						&Call{"*3", Registers{"s"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "func(string) string601598840"},
						&Call{"*5", Registers{"suffix"}, Registers{"6"}, "any"},
						&AssignSymbol{"7", "func(string, string) string601598846"},
						&Call{"*7", Registers{"4", "6"}, Registers{"8"}, "any"},
						&AssignSymbol{"9", "func(string) string601598840"},
						&Call{"*9", Registers{"8"}, Registers{"10"}, "any"},
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
						&AssignSymbol{"3", "string"},
						&Assign{"result", "3"},
						&Len{"s", "4"},
						&LessThanNumber{"index", "4", "5"},
						&JumpUnless{"5", 10},
						&StringIndex{"s", "index", "6"},
						&CastString{"6", "7"},
						&Concat{"result", "7", "result"},
						&AssignSymbol{"8", "number1"},
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
					&AssignFunc{"1", "func(string) string", "601598823"},
					&ParentScope{"1"},
					&AssignFunc{"2", "func(string) string", "601598824"},
					&ParentScope{"2"},
					&AssignFunc{"3", "func(string, string) bool", "601598825"},
					&ParentScope{"3"},
					&AssignFunc{"4", "func(string, string) bool", "601598826"},
					&ParentScope{"4"},
					&AssignFunc{"5", "func(string, string) bool", "601598827"},
					&ParentScope{"5"},
					&AssignFunc{"6", "func(string, string) number", "601598828"},
					&ParentScope{"6"},
					&AssignFunc{"7", "func(string, string, number) number", "601598829"},
					&ParentScope{"7"},
					&AssignFunc{"8", "func(number, number) number", "601598830"},
					&ParentScope{"8"},
					&AssignFunc{"9", "func(number, number) number", "601598831"},
					&ParentScope{"9"},
					&AssignFunc{"10", "func(string, string) number", "601598832"},
					&ParentScope{"10"},
					&AssignFunc{"11", "func(string, string, number) number", "601598833"},
					&ParentScope{"11"},
					&AssignFunc{"12", "func([]string, string) string", "601598834"},
					&ParentScope{"12"},
					&AssignFunc{"13", "func(string, string, number) string", "601598835"},
					&ParentScope{"13"},
					&AssignFunc{"14", "func(string, string, number) string", "601598836"},
					&ParentScope{"14"},
					&AssignFunc{"15", "func(string, number) string", "601598837"},
					&ParentScope{"15"},
					&AssignFunc{"16", "func(string, number) string", "601598838"},
					&ParentScope{"16"},
					&AssignFunc{"17", "func(string, string, string) string", "601598839"},
					&ParentScope{"17"},
					&AssignFunc{"18", "func(string) string", "601598840"},
					&ParentScope{"18"},
					&AssignFunc{"19", "func(string, string) []string", "601598841"},
					&ParentScope{"19"},
					&AssignFunc{"20", "func(string, number, number) string", "601598842"},
					&ParentScope{"20"},
					&AssignFunc{"21", "func(string, string) string", "601598843"},
					&ParentScope{"21"},
					&AssignFunc{"22", "func(string, string) string", "601598844"},
					&ParentScope{"22"},
					&AssignFunc{"23", "func(string, string) string", "601598845"},
					&ParentScope{"23"},
					&AssignFunc{"24", "func(string, string) string", "601598846"},
					&ParentScope{"24"},
					&AssignFunc{"25", "func(string, string) string", "601598847"},
					&ParentScope{"25"},
					&AssignFunc{"26", "func(string, number) string", "601598848"},
					&ParentScope{"26"},
					&AssignFunc{"27", "func(string, string) bool", "601598825"},
					&ParentScope{"27"},
					&AssignFunc{"28", "func(string, string) bool", "601598826"},
					&ParentScope{"28"},
					&AssignFunc{"29", "func(string, string) bool", "601598827"},
					&ParentScope{"29"},
					&AssignFunc{"30", "func(string, string) number", "601598828"},
					&ParentScope{"30"},
					&AssignFunc{"31", "func(string, string, number) number", "601598829"},
					&ParentScope{"31"},
					&AssignFunc{"32", "func([]string, string) string", "601598834"},
					&ParentScope{"32"},
					&AssignFunc{"33", "func(string, string) number", "601598832"},
					&ParentScope{"33"},
					&AssignFunc{"34", "func(string, string, number) number", "601598833"},
					&ParentScope{"34"},
					&AssignFunc{"35", "func(string, string, number) string", "601598835"},
					&ParentScope{"35"},
					&AssignFunc{"36", "func(string, string, number) string", "601598836"},
					&ParentScope{"36"},
					&AssignFunc{"37", "func(string, number) string", "601598838"},
					&ParentScope{"37"},
					&AssignFunc{"38", "func(string, string, string) string", "601598839"},
					&ParentScope{"38"},
					&AssignFunc{"39", "func(string) string", "601598840"},
					&ParentScope{"39"},
					&AssignFunc{"40", "func(string, string) []string", "601598841"},
					&ParentScope{"40"},
					&AssignFunc{"41", "func(string, number, number) string", "601598842"},
					&ParentScope{"41"},
					&AssignFunc{"42", "func(string) string", "601598823"},
					&ParentScope{"42"},
					&AssignFunc{"43", "func(string) string", "601598824"},
					&ParentScope{"43"},
					&AssignFunc{"44", "func(string, string) string", "601598845"},
					&ParentScope{"44"},
					&AssignFunc{"45", "func(string, string) string", "601598843"},
					&ParentScope{"45"},
					&AssignFunc{"46", "func(string, string) string", "601598846"},
					&ParentScope{"46"},
					&AssignFunc{"47", "func(string, string) string", "601598844"},
					&ParentScope{"47"},
					&AssignFunc{"48", "func(string, string) string", "601598847"},
					&ParentScope{"48"},
					&AssignFunc{"49", "func(string, number) string", "601598837"},
					&ParentScope{"49"},
					&AssignFunc{"50", "func(number, number) number", "601598831"},
					&ParentScope{"50"},
					&AssignFunc{"51", "func(number, number) number", "601598830"},
					&ParentScope{"51"},
					&AssignFunc{"52", "func(string, number) string", "601598848"},
					&ParentScope{"52"},
					&Assign{"601598823", "1"},
					&Assign{"601598824", "2"},
					&Assign{"601598825", "3"},
					&Assign{"601598826", "4"},
					&Assign{"601598827", "5"},
					&Assign{"601598828", "6"},
					&Assign{"601598829", "7"},
					&Assign{"601598830", "8"},
					&Assign{"601598831", "9"},
					&Assign{"601598832", "10"},
					&Assign{"601598833", "11"},
					&Assign{"601598834", "12"},
					&Assign{"601598835", "13"},
					&Assign{"601598836", "14"},
					&Assign{"601598837", "15"},
					&Assign{"601598838", "16"},
					&Assign{"601598839", "17"},
					&Assign{"601598840", "18"},
					&Assign{"601598841", "19"},
					&Assign{"601598842", "20"},
					&Assign{"601598843", "21"},
					&Assign{"601598844", "22"},
					&Assign{"601598845", "23"},
					&Assign{"601598846", "24"},
					&Assign{"601598847", "25"},
					&Assign{"601598848", "26"},
					&Assign{"Contains", "27"},
					&Assign{"HasPrefix", "28"},
					&Assign{"HasSuffix", "29"},
					&Assign{"Index", "30"},
					&Assign{"IndexAfter", "31"},
					&Assign{"Join", "32"},
					&Assign{"LastIndex", "33"},
					&Assign{"LastIndexBefore", "34"},
					&Assign{"PadLeft", "35"},
					&Assign{"PadRight", "36"},
					&Assign{"Repeat", "37"},
					&Assign{"ReplaceAll", "38"},
					&Assign{"Reverse", "39"},
					&Assign{"Split", "40"},
					&Assign{"Substr", "41"},
					&Assign{"ToLower", "42"},
					&Assign{"ToUpper", "43"},
					&Assign{"Trim", "44"},
					&Assign{"TrimLeft", "45"},
					&Assign{"TrimPrefix", "46"},
					&Assign{"TrimRight", "47"},
					&Assign{"TrimSuffix", "48"},
					&Assign{"createPad", "49"},
					&Assign{"max", "50"},
					&Assign{"min", "51"},
					&Assign{"substrFrom", "52"},
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
			Types: map[TypeRegister]*types.Type{
				"[]string": &types.Type{
					Kind: 8,
					Element: &types.Type{
						Kind: 7,
					},
				},
				"any": &types.Type{
					Kind: 2,
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"boolfalse": &Symbol{
					Type:  "bool",
					Value: "false",
				},
				"booltrue": &Symbol{
					Type:  "bool",
					Value: "true",
				},
				"charA": &Symbol{
					Type:  "char",
					Value: "A",
				},
				"charZ": &Symbol{
					Type:  "char",
					Value: "Z",
				},
				"chara": &Symbol{
					Type:  "char",
					Value: "a",
				},
				"charz": &Symbol{
					Type:  "char",
					Value: "z",
				},
				"func([]string, string) string601598834": &Symbol{
					Type:  "func([]string, string) string",
					Value: "601598834",
				},
				"func(number, number) number601598830": &Symbol{
					Type:  "func(number, number) number",
					Value: "601598830",
				},
				"func(number, number) number601598831": &Symbol{
					Type:  "func(number, number) number",
					Value: "601598831",
				},
				"func(string) string601598840": &Symbol{
					Type:  "func(string) string",
					Value: "601598840",
				},
				"func(string, number) string601598837": &Symbol{
					Type:  "func(string, number) string",
					Value: "601598837",
				},
				"func(string, number) string601598838": &Symbol{
					Type:  "func(string, number) string",
					Value: "601598838",
				},
				"func(string, number) string601598848": &Symbol{
					Type:  "func(string, number) string",
					Value: "601598848",
				},
				"func(string, number, number) string601598842": &Symbol{
					Type:  "func(string, number, number) string",
					Value: "601598842",
				},
				"func(string, string) []string601598841": &Symbol{
					Type:  "func(string, string) []string",
					Value: "601598841",
				},
				"func(string, string) bool601598826": &Symbol{
					Type:  "func(string, string) bool",
					Value: "601598826",
				},
				"func(string, string) number601598828": &Symbol{
					Type:  "func(string, string) number",
					Value: "601598828",
				},
				"func(string, string) string601598843": &Symbol{
					Type:  "func(string, string) string",
					Value: "601598843",
				},
				"func(string, string) string601598844": &Symbol{
					Type:  "func(string, string) string",
					Value: "601598844",
				},
				"func(string, string) string601598846": &Symbol{
					Type:  "func(string, string) string",
					Value: "601598846",
				},
				"func(string, string, number) number601598829": &Symbol{
					Type:  "func(string, string, number) number",
					Value: "601598829",
				},
				"number-1": &Symbol{
					Type:  "number",
					Value: "-1",
				},
				"number0": &Symbol{
					Type:  "number",
					Value: "0",
				},
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number32": &Symbol{
					Type:  "number",
					Value: "32",
				},
				"string": &Symbol{
					Type: "string",
				},
			},
		},
		"time": &File{
			Imports: map[string]*types.Type{
				"math": &types.Type{
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
			Funcs: map[string]*CompiledFunc{
				"3492229393": &CompiledFunc{
					Arguments: []string{"t", "duration"},
					Instructions: []Instruction{
						&AssignSymbol{"3", "func(Time) number3492229410"},
						&Call{"*3", Registers{"t"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "stringSeconds"},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, "any"},
						&Add{"4", "7", "8"},
						&AssignSymbol{"9", "func(number) Time3492229411"},
						&Call{"*9", Registers{"8"}, Registers{"10"}, "Time"},
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
						&AssignSymbol{"3", "func(Time) number3492229410"},
						&Call{"*3", Registers{"a"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "func(Time) number3492229410"},
						&Call{"*5", Registers{"b"}, Registers{"6"}, "any"},
						&Subtract{"4", "6", "7"},
						&AssignSymbol{"8", "func(number) Duration3492229398"},
						&Call{"*8", Registers{"7"}, Registers{"9"}, "Duration"},
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
						&AssignSymbol{"3", "func(Time, Time) Duration3492229394"},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, "Duration"},
						&Assign{"duration", "4"},
						&AssignSymbol{"5", "stringSeconds"},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, "any"},
						&AssignSymbol{"8", "number0"},
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
						&AssignSymbol{"3", "func(Time, Time) Duration3492229394"},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, "Duration"},
						&Assign{"duration", "4"},
						&AssignSymbol{"5", "stringSeconds"},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, "any"},
						&AssignSymbol{"8", "number0"},
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
						&AssignSymbol{"3", "func(Time, Time) Duration3492229394"},
						&Call{"*3", Registers{"a", "b"}, Registers{"4"}, "Duration"},
						&Assign{"duration", "4"},
						&AssignSymbol{"5", "stringSeconds"},
						&MapGet{"duration", "5", "6"},
						&Call{"*6", nil, Registers{"7"}, "any"},
						&AssignSymbol{"8", "number0"},
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
						&AssignFunc{"2", "func() string", "3492229405"},
						&ParentScope{"2"},
						&AssignFunc{"3", "func() number", "3492229404"},
						&ParentScope{"3"},
						&AssignFunc{"4", "func() number", "3492229403"},
						&ParentScope{"4"},
						&AssignFunc{"5", "func() number", "3492229402"},
						&ParentScope{"5"},
						&AssignFunc{"6", "func() number", "3492229401"},
						&ParentScope{"6"},
						&AssignFunc{"7", "func() number", "3492229400"},
						&ParentScope{"7"},
						&AssignFunc{"8", "func() number", "3492229399"},
						&ParentScope{"8"},
						&Assign{"String", "2"},
						&Assign{"Hours", "3"},
						&Assign{"Minutes", "4"},
						&Assign{"Seconds", "5"},
						&Assign{"Milliseconds", "6"},
						&Assign{"Microseconds", "7"},
						&Assign{"Nanoseconds", "8"},
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
					Pos:        "lib/time/duration.ok:13:1",
				},
				"3492229399": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "number0.000000001"},
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
					Pos:        "lib/time/duration.ok:14:5",
				},
				"3492229400": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "number0.000001"},
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
					Pos:        "lib/time/duration.ok:18:5",
				},
				"3492229401": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "number0.001"},
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
					Pos:        "lib/time/duration.ok:22:5",
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
					Pos:        "lib/time/duration.ok:26:5",
				},
				"3492229403": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "number60"},
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
					Pos:        "lib/time/duration.ok:30:5",
				},
				"3492229404": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "number3600"},
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
					Pos:        "lib/time/duration.ok:34:5",
				},
				"3492229405": &CompiledFunc{
					Instructions: []Instruction{
						&AssignSymbol{"1", "string"},
						&Assign{"s", "1"},
						&Assign{"seconds", "^seconds"},
						&AssignSymbol{"2", "number3600"},
						&GreaterThanEqualNumber{"seconds", "2", "3"},
						&JumpUnless{"3", 18},
						&AssignSymbol{"4", "number3600"},
						&Divide{"seconds", "4", "5"},
						&LoadPackage{"6", "math"},
						&AssignSymbol{"7", "stringFloor"},
						&MapGet{"6", "7", "8"},
						&Call{"*8", Registers{"5"}, Registers{"9"}, "any"},
						&Assign{"hours", "9"},
						&AssignSymbol{"10", "number3600"},
						&Multiply{"hours", "10", "11"},
						&Subtract{"seconds", "11", "seconds"},
						&AssignSymbol{"13", "stringh"},
						&Interpolate{"12", Registers{"hours", "13"}},
						&Concat{"s", "12", "s"},
						&AssignSymbol{"14", "number60"},
						&GreaterThanEqualNumber{"seconds", "14", "15"},
						&JumpUnless{"15", 34},
						&AssignSymbol{"16", "number60"},
						&Divide{"seconds", "16", "17"},
						&LoadPackage{"18", "math"},
						&AssignSymbol{"19", "stringFloor"},
						&MapGet{"18", "19", "20"},
						&Call{"*20", Registers{"17"}, Registers{"21"}, "any"},
						&Assign{"minutes", "21"},
						&AssignSymbol{"22", "number60"},
						&Multiply{"minutes", "22", "23"},
						&Subtract{"seconds", "23", "seconds"},
						&AssignSymbol{"25", "stringm"},
						&Interpolate{"24", Registers{"minutes", "25"}},
						&Concat{"s", "24", "s"},
						&AssignSymbol{"26", "number0"},
						&NotEqualNumber{"seconds", "26", "27"},
						&JumpUnless{"27", 40},
						&AssignSymbol{"29", "strings"},
						&Interpolate{"28", Registers{"seconds", "29"}},
						&Concat{"s", "28", "s"},
						&Return{Registers{"s"}},
					},
					Registers: 29,
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
					Pos:        "lib/time/duration.ok:38:5",
				},
				"3492229406": &CompiledFunc{
					Arguments: []string{"duration"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "stringSeconds"},
						&MapGet{"duration", "2", "3"},
						&Call{"*3", nil, Registers{"4"}, "any"},
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
					UniqueName: "3492229406",
					Pos:        "lib/time/sleep.ok:2:1",
				},
				"3492229407": &CompiledFunc{
					Arguments: []string{"Year", "Month", "Day", "Hour", "Minute", "Second"},
					Instructions: []Instruction{
						&AssignFunc{"7", "func() string", "3492229408"},
						&ParentScope{"7"},
						&Assign{"String", "7"},
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
					UniqueName: "3492229407",
					Pos:        "lib/time/time.ok:15:1",
				},
				"3492229408": &CompiledFunc{
					Instructions: []Instruction{
						&Assign{"month", "^Month"},
						&Assign{"day", "^Day"},
						&Assign{"hour", "^Hour"},
						&Assign{"minute", "^Minute"},
						&Assign{"second", "^Second"},
						&AssignSymbol{"2", "string-"},
						&AssignSymbol{"3", "func(number) string3492229412"},
						&Call{"*3", Registers{"month"}, Registers{"4"}, "any"},
						&AssignSymbol{"5", "string-"},
						&AssignSymbol{"6", "func(number) string3492229412"},
						&Call{"*6", Registers{"day"}, Registers{"7"}, "any"},
						&AssignSymbol{"8", "string "},
						&AssignSymbol{"9", "func(number) string3492229412"},
						&Call{"*9", Registers{"hour"}, Registers{"10"}, "any"},
						&AssignSymbol{"11", "string:"},
						&AssignSymbol{"12", "func(number) string3492229412"},
						&Call{"*12", Registers{"minute"}, Registers{"13"}, "any"},
						&AssignSymbol{"14", "string:"},
						&AssignSymbol{"15", "func(number) string3492229412"},
						&Call{"*15", Registers{"second"}, Registers{"16"}, "any"},
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
					UniqueName: "3492229408",
					Pos:        "lib/time/time.ok:16:5",
				},
				"3492229409": &CompiledFunc{
					Instructions: []Instruction{
						&Now{"1", "2", "3", "4", "5", "6"},
						&Assign{"year", "1"},
						&Assign{"month", "2"},
						&Assign{"day", "3"},
						&Assign{"hour", "4"},
						&Assign{"minute", "5"},
						&Assign{"second", "6"},
						&AssignSymbol{"7", "func(number, number, number, number, number, number) Time3492229407"},
						&Call{"*7", Registers{"year", "month", "day", "hour", "minute", "second"}, Registers{"8"}, "Time"},
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
					UniqueName: "3492229409",
					Pos:        "lib/time/time.ok:32:1",
				},
				"3492229410": &CompiledFunc{
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
					UniqueName: "3492229410",
					Pos:        "lib/time/unix.ok:3:1",
				},
				"3492229411": &CompiledFunc{
					Arguments: []string{"seconds"},
					Instructions: []Instruction{
						&FromUnix{"seconds", "2", "3", "4", "5", "6", "7"},
						&Assign{"year", "2"},
						&Assign{"month", "3"},
						&Assign{"day", "4"},
						&Assign{"hour", "5"},
						&Assign{"minute", "6"},
						&Assign{"second", "7"},
						&AssignSymbol{"8", "func(number, number, number, number, number, number) Time3492229407"},
						&Call{"*8", Registers{"year", "month", "day", "hour", "minute", "second"}, Registers{"9"}, "Time"},
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
					UniqueName: "3492229411",
					Pos:        "lib/time/unix.ok:10:1",
				},
				"3492229412": &CompiledFunc{
					Arguments: []string{"n"},
					Instructions: []Instruction{
						&AssignSymbol{"2", "number10"},
						&LessThanNumber{"n", "2", "3"},
						&JumpUnless{"3", 6},
						&AssignSymbol{"4", "string0"},
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
					UniqueName: "3492229412",
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
				}, "3600", nil, nil, "lib/time/duration.ok:10:15", nil, nil},
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
				}, "0.000001", nil, nil, "lib/time/duration.ok:6:15", nil, nil},
				"Millisecond": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.001", nil, nil, "lib/time/duration.ok:7:15", nil, nil},
				"Minute": &ast.Literal{&types.Type{
					Kind: 6,
				}, "60", nil, nil, "lib/time/duration.ok:9:15", nil, nil},
				"Nanosecond": &ast.Literal{&types.Type{
					Kind: 6,
				}, "0.000000001", nil, nil, "lib/time/duration.ok:5:15", nil, nil},
				"November": &ast.Literal{&types.Type{
					Kind: 6,
				}, "11", nil, nil, "lib/time/time.ok:11:12", nil, nil},
				"October": &ast.Literal{&types.Type{
					Kind: 6,
				}, "10", nil, nil, "lib/time/time.ok:10:11", nil, nil},
				"Second": &ast.Literal{&types.Type{
					Kind: 6,
				}, "1", nil, nil, "lib/time/duration.ok:8:15", nil, nil},
				"September": &ast.Literal{&types.Type{
					Kind: 6,
				}, "9", nil, nil, "lib/time/time.ok:9:13", nil, nil},
			},
			PackageFunc: &CompiledFunc{
				Instructions: []Instruction{
					&AssignFunc{"19", "func(Time, Duration) Time", "3492229393"},
					&ParentScope{"19"},
					&AssignFunc{"20", "func(Time, Time) Duration", "3492229394"},
					&ParentScope{"20"},
					&AssignFunc{"21", "func(Time, Time) bool", "3492229395"},
					&ParentScope{"21"},
					&AssignFunc{"22", "func(Time, Time) bool", "3492229396"},
					&ParentScope{"22"},
					&AssignFunc{"23", "func(Time, Time) bool", "3492229397"},
					&ParentScope{"23"},
					&AssignFunc{"24", "func(number) Duration", "3492229398"},
					&ParentScope{"24"},
					&AssignFunc{"25", "func(Duration)", "3492229406"},
					&ParentScope{"25"},
					&AssignFunc{"26", "func(number, number, number, number, number, number) Time", "3492229407"},
					&ParentScope{"26"},
					&AssignFunc{"27", "func() Time", "3492229409"},
					&ParentScope{"27"},
					&AssignFunc{"28", "func(Time) number", "3492229410"},
					&ParentScope{"28"},
					&AssignFunc{"29", "func(number) Time", "3492229411"},
					&ParentScope{"29"},
					&AssignFunc{"30", "func(number) string", "3492229412"},
					&ParentScope{"30"},
					&AssignFunc{"31", "func(Time, Duration) Time", "3492229393"},
					&ParentScope{"31"},
					&AssignFunc{"32", "func(Time, Time) bool", "3492229397"},
					&ParentScope{"32"},
					&AssignFunc{"33", "func(Time, Time) bool", "3492229396"},
					&ParentScope{"33"},
					&AssignFunc{"34", "func(number) Duration", "3492229398"},
					&ParentScope{"34"},
					&AssignFunc{"35", "func(Time, Time) bool", "3492229395"},
					&ParentScope{"35"},
					&AssignFunc{"36", "func(number) Time", "3492229411"},
					&ParentScope{"36"},
					&AssignFunc{"37", "func() Time", "3492229409"},
					&ParentScope{"37"},
					&AssignFunc{"38", "func(Duration)", "3492229406"},
					&ParentScope{"38"},
					&AssignFunc{"39", "func(Time, Time) Duration", "3492229394"},
					&ParentScope{"39"},
					&AssignFunc{"40", "func(number, number, number, number, number, number) Time", "3492229407"},
					&ParentScope{"40"},
					&AssignFunc{"41", "func(Time) number", "3492229410"},
					&ParentScope{"41"},
					&AssignFunc{"42", "func(number) string", "3492229412"},
					&ParentScope{"42"},
					&AssignSymbol{"1", "number4"},
					&Assign{"April", "1"},
					&AssignSymbol{"2", "number8"},
					&Assign{"August", "2"},
					&AssignSymbol{"3", "number12"},
					&Assign{"December", "3"},
					&AssignSymbol{"4", "number2"},
					&Assign{"February", "4"},
					&AssignSymbol{"5", "number3600"},
					&Assign{"Hour", "5"},
					&AssignSymbol{"6", "number1"},
					&Assign{"January", "6"},
					&AssignSymbol{"7", "number7"},
					&Assign{"July", "7"},
					&AssignSymbol{"8", "number6"},
					&Assign{"June", "8"},
					&AssignSymbol{"9", "number3"},
					&Assign{"March", "9"},
					&AssignSymbol{"10", "number5"},
					&Assign{"May", "10"},
					&AssignSymbol{"11", "number0.000001"},
					&Assign{"Microsecond", "11"},
					&AssignSymbol{"12", "number0.001"},
					&Assign{"Millisecond", "12"},
					&AssignSymbol{"13", "number60"},
					&Assign{"Minute", "13"},
					&AssignSymbol{"14", "number0.000000001"},
					&Assign{"Nanosecond", "14"},
					&AssignSymbol{"15", "number11"},
					&Assign{"November", "15"},
					&AssignSymbol{"16", "number10"},
					&Assign{"October", "16"},
					&AssignSymbol{"17", "number1"},
					&Assign{"Second", "17"},
					&AssignSymbol{"18", "number9"},
					&Assign{"September", "18"},
					&Assign{"3492229393", "19"},
					&Assign{"3492229394", "20"},
					&Assign{"3492229395", "21"},
					&Assign{"3492229396", "22"},
					&Assign{"3492229397", "23"},
					&Assign{"3492229398", "24"},
					&Assign{"3492229406", "25"},
					&Assign{"3492229407", "26"},
					&Assign{"3492229409", "27"},
					&Assign{"3492229410", "28"},
					&Assign{"3492229411", "29"},
					&Assign{"3492229412", "30"},
					&Assign{"Add", "31"},
					&Assign{"After", "32"},
					&Assign{"Before", "33"},
					&Assign{"Duration", "34"},
					&Assign{"Equal", "35"},
					&Assign{"FromUnix", "36"},
					&Assign{"Now", "37"},
					&Assign{"Sleep", "38"},
					&Assign{"Sub", "39"},
					&Assign{"Time", "40"},
					&Assign{"Unix", "41"},
					&Assign{"zeroPad", "42"},
					&Return{Registers{"0"}},
				},
				Registers: 42,
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
			Types: map[TypeRegister]*types.Type{
				"Duration": &types.Type{
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
				"Time": &types.Type{
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
				"any": &types.Type{
					Kind: 2,
				},
				"func() number": &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 6,
						},
					},
				},
				"func() string": &types.Type{
					Kind: 10,
					Returns: []*types.Type{
						&types.Type{
							Kind: 7,
						},
					},
				},
			},
			Symbols: map[SymbolRegister]*Symbol{
				"func(Time) number3492229410": &Symbol{
					Type:  "func(Time) number",
					Value: "3492229410",
				},
				"func(Time, Time) Duration3492229394": &Symbol{
					Type:  "func(Time, Time) Duration",
					Value: "3492229394",
				},
				"func(number) Duration3492229398": &Symbol{
					Type:  "func(number) Duration",
					Value: "3492229398",
				},
				"func(number) Time3492229411": &Symbol{
					Type:  "func(number) Time",
					Value: "3492229411",
				},
				"func(number) string3492229412": &Symbol{
					Type:  "func(number) string",
					Value: "3492229412",
				},
				"func(number, number, number, number, number, number) Time3492229407": &Symbol{
					Type:  "func(number, number, number, number, number, number) Time",
					Value: "3492229407",
				},
				"number0": &Symbol{
					Type:  "number",
					Value: "0",
				},
				"number0.000000001": &Symbol{
					Type:  "number",
					Value: "0.000000001",
				},
				"number0.000001": &Symbol{
					Type:  "number",
					Value: "0.000001",
				},
				"number0.001": &Symbol{
					Type:  "number",
					Value: "0.001",
				},
				"number1": &Symbol{
					Type:  "number",
					Value: "1",
				},
				"number10": &Symbol{
					Type:  "number",
					Value: "10",
				},
				"number11": &Symbol{
					Type:  "number",
					Value: "11",
				},
				"number12": &Symbol{
					Type:  "number",
					Value: "12",
				},
				"number2": &Symbol{
					Type:  "number",
					Value: "2",
				},
				"number3": &Symbol{
					Type:  "number",
					Value: "3",
				},
				"number3600": &Symbol{
					Type:  "number",
					Value: "3600",
				},
				"number4": &Symbol{
					Type:  "number",
					Value: "4",
				},
				"number5": &Symbol{
					Type:  "number",
					Value: "5",
				},
				"number6": &Symbol{
					Type:  "number",
					Value: "6",
				},
				"number60": &Symbol{
					Type:  "number",
					Value: "60",
				},
				"number7": &Symbol{
					Type:  "number",
					Value: "7",
				},
				"number8": &Symbol{
					Type:  "number",
					Value: "8",
				},
				"number9": &Symbol{
					Type:  "number",
					Value: "9",
				},
				"string": &Symbol{
					Type: "string",
				},
				"string ": &Symbol{
					Type:  "string",
					Value: " ",
				},
				"string-": &Symbol{
					Type:  "string",
					Value: "-",
				},
				"string0": &Symbol{
					Type:  "string",
					Value: "0",
				},
				"string:": &Symbol{
					Type:  "string",
					Value: ":",
				},
				"stringFloor": &Symbol{
					Type:  "string",
					Value: "Floor",
				},
				"stringSeconds": &Symbol{
					Type:  "string",
					Value: "Seconds",
				},
				"stringh": &Symbol{
					Type:  "string",
					Value: "h",
				},
				"stringm": &Symbol{
					Type:  "string",
					Value: "m",
				},
				"strings": &Symbol{
					Type:  "string",
					Value: "s",
				},
			},
		},
	}
}
