package ast

import (
	"strings"

	"github.com/elliotchance/ok/util"
)

// Argument is used to define a name and type for a function argument.
type Argument struct {
	Name string
	Type string
}

func (arg *Argument) String() string {
	return strings.TrimSpace(arg.Name + " " + arg.Type)
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
	Returns []string

	// Statements can have zero or more elements for each of the ordered
	// discreet statements in the function.
	Statements []Node

	Pos string
}

func (f *Func) Interface() (map[string]string, error) {
	fields := map[string]string{}

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
	return f.signature(true)
}

func (f *Func) Type() string {
	return f.signature(false)
}

func (f *Func) signature(includeName bool) string {
	var args []string
	for _, arg := range f.Arguments {
		if arg.Name == "" {
			args = append(args, arg.Type)
		} else {
			args = append(args, arg.Name+" "+arg.Type)
		}
	}

	returnSignature := ""
	if len(f.Returns) == 1 {
		returnSignature = " " + f.Returns[0]
	}
	if len(f.Returns) > 1 {
		returnSignature = " (" + strings.Join(f.Returns, ", ") + ")"
	}

	prefix := "func"
	if f.Name != "" && includeName {
		prefix += " " + f.Name
	}

	return prefix + "(" + strings.Join(args, ", ") + ")" + returnSignature
}

// Position returns the position.
func (f *Func) Position() string {
	return f.Pos
}

func (f *Func) IsConstructor() bool {
	return f.Name != "" && f.Name == strings.Join(f.Returns, ",")
}

// NewFuncFromPrototype is a hack for now. It should be derived directly from
// the type itself.
func NewFuncFromPrototype(ty string) *Func {
	f := &Func{}

	// TODO(elliot): This is a bad solution. Fix me.
	parts := strings.Split(ty[5:], ")")

	if strings.TrimSpace(parts[0]) != "" {
		for _, a := range util.StringSliceMap(strings.Split(parts[0], ","), strings.TrimSpace) {
			f.Arguments = append(f.Arguments, &Argument{Name: "", Type: a})
		}
	}

	if strings.TrimSpace(parts[1]) != "" {
		f.Returns = util.StringSliceMap(strings.Split(parts[1], ","), strings.TrimSpace)
	}

	return f
}
