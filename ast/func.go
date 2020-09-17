package ast

import (
	"strings"

	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
)

// Argument is used to define a name and type for a function argument.
type Argument struct {
	Name string
	Type *types.Type
}

func (arg *Argument) String() string {
	return strings.TrimSpace(arg.Name + " " + arg.Type.String())
}

// Func represents the definition of a function.
type Func struct {
	// Name is the name of the function being declared.
	Name string

	// Arguments may contain zero or more elements. They will always be in the
	// order in which their are declared.
	Arguments []*Argument

	// Returns may contain zero or more types. They will always be in the order
	// in which they are declared.
	Returns []*types.Type

	// Statements can have zero or more elements for each of the ordered
	// discreet statements in the function.
	Statements []Node

	Pos string
}

func (f *Func) Interface() (map[string]*types.Type, error) {
	fields := map[string]*types.Type{}

	for _, arg := range f.Arguments {
		if util.IsPublic(arg.Name) {
			fields[arg.Name] = arg.Type
		}
	}

	for _, stmt := range f.Statements {
		if a, ok := stmt.(*Assign); ok {
			if len(a.Lefts) == 1 {
				if b, ok := a.Lefts[0].(*Identifier); ok {
					if util.IsPublic(b.Name) {
						ty, err := TypeOf(a.Rights[0])
						if err != nil {
							return nil, err
						}

						fields[b.Name] = ty
					}
				}
			}
		}
	}

	return fields, nil
}

// String returns the signature, like:
//
//   func Foo(x number, y number) (string, string)
//
func (f *Func) String() string {
	var args []string
	for _, arg := range f.Arguments {
		args = append(args, arg.Name+" "+arg.Type.String())
	}

	returnSignature := ""
	if len(f.Returns) == 1 {
		returnSignature = " " + f.Returns[0].String()
	}
	if len(f.Returns) > 1 {
		var returns []string
		for _, r := range f.Returns {
			returns = append(returns, r.String())
		}
		returnSignature = " (" + strings.Join(returns, ", ") + ")"
	}

	prefix := "func"
	if f.Name != "" {
		prefix += " " + f.Name
	}

	return prefix + "(" + strings.Join(args, ", ") + ")" + returnSignature
}

func (f *Func) Type() *types.Type {
	var args, returns []*types.Type
	for _, arg := range f.Arguments {
		args = append(args, arg.Type)
	}
	for _, r := range f.Returns {
		returns = append(returns, r)
	}

	return types.NewFunc(args, returns)
}

// Position returns the position.
func (f *Func) Position() string {
	return f.Pos
}

func (f *Func) IsConstructor() bool {
	return f.Name != "" && len(f.Returns) == 1 && f.Name == f.Returns[0].Name
}

// NewFuncFromPrototype is a hack for now. It should be derived directly from
// the type itself.
func NewFuncFromPrototype(ty *types.Type) *Func {
	f := &Func{}

	for _, a := range ty.Arguments {
		f.Arguments = append(f.Arguments, &Argument{Name: "", Type: a})
	}

	f.Returns = ty.Returns

	return f
}
