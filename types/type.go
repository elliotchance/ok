package types

import (
	"sort"
	"strings"
)

// Type represents a data type.
type Type struct {
	// Kind will not be set when Ref is supplied.
	Kind Kind `json:",omitempty"`

	// Name is used as the descriptive name for the object.
	Name string `json:",omitempty"`

	// Element is used when Kind is an Array or Map.
	Element *Type `json:",omitempty"`

	// Argument and Returns are used when Kind is a Func. Either may be nil.
	Arguments, Returns []*Type `json:",omitempty"`

	// Properties is used for KindInterface
	Properties map[string]*Type `json:",omitempty"`

	// Ref is used when types are flattened for a Registry. It will point to an
	// index of another type in the registry.
	Ref string `json:",omitempty"`
}

// ToArray creates an array type using this element type.
func (t *Type) ToArray() *Type {
	return &Type{
		Kind:    KindArray,
		Element: t,
	}
}

// ToMap creates a map type using this element type.
func (t *Type) ToMap() *Type {
	return &Type{
		Kind:    KindMap,
		Element: t,
	}
}

func (t *Type) Copy() *Type {
	// This is so we don't have to check for nils on all the callers.
	if t == nil {
		return nil
	}

	ty := &Type{
		Kind:    t.Kind,
		Name:    t.Name,
		Ref:     t.Ref,
		Element: t.Element.Copy(),
	}

	for _, v := range t.Arguments {
		ty.Arguments = append(ty.Arguments, v.Copy())
	}

	for _, v := range t.Returns {
		ty.Returns = append(ty.Returns, v.Copy())
	}

	if len(t.Properties) > 0 {
		ty.Properties = map[string]*Type{}

		for k, v := range t.Properties {
			ty.Properties[k] = v.Copy()
		}
	}

	return ty
}

func (t *Type) String() string {
	if t == nil {
		return ""
	}

	switch t.Kind {
	case KindAny:
		return "any"

	case KindBool:
		return "bool"

	case KindChar:
		return "char"

	case KindData:
		return "data"

	case KindNumber:
		return "number"

	case KindString:
		return "string"

	case KindArray:
		return "[]" + t.Element.String()

	case KindMap:
		return "{}" + t.Element.String()

	case KindFunc:
		var args []string
		for _, arg := range t.Arguments {
			args = append(args, arg.String())
		}

		s := "func(" + strings.Join(args, ", ") + ")"

		switch len(t.Returns) {
		case 0:
			// Do nothing

		case 1:
			s += " " + t.Returns[0].String()

		default:
			var returns []string
			for _, r := range t.Returns {
				returns = append(returns, r.String())
			}

			s += " (" + strings.Join(returns, ", ") + ")"
		}

		return s
	}

	return t.Name
}

func NewArray(element *Type) *Type {
	return &Type{
		Kind:    KindArray,
		Element: element,
	}
}

func NewMap(element *Type) *Type {
	return &Type{
		Kind:    KindMap,
		Element: element,
	}
}

func NewRef(ref string) *Type {
	return &Type{
		Ref: ref,
	}
}

func NewFunc(args, returns []*Type) *Type {
	return &Type{
		Kind:      KindFunc,
		Arguments: args,
		Returns:   returns,
	}
}

// TypeFromString decodes a syntactically-valid type from a string.
func TypeFromString(s string) *Type {
	tokens := tokenize(s)
	ty, _ := parseType(tokens, 0)

	return ty
}

func NewUnresolvedInterface(name string) *Type {
	return &Type{
		Kind: KindUnresolvedInterface,
		Name: name,
	}
}

func NewInterface(name string, properties map[string]*Type) *Type {
	return &Type{
		Kind:       KindResolvedInterface,
		Name:       name,
		Properties: properties,
	}
}

func (t *Type) Interface() string {
	s := "{ "
	for i, key := range t.SortedPropertyNames() {
		if i > 0 {
			s += "; "
		}

		if t.Properties[key].Kind == KindFunc {
			s += key + strings.TrimPrefix(t.Properties[key].String(), "func")
		} else {
			s += key + " " + t.Properties[key].String()
		}
	}

	if s == "{ " {
		return "{}"
	}

	return s + " }"
}

func (t *Type) SortedPropertyNames() []string {
	// Avoid allocation in most cases.
	if len(t.Properties) == 0 {
		return nil
	}

	names := make([]string, len(t.Properties))
	i := 0
	for name := range t.Properties {
		names[i] = name
		i++
	}
	sort.Strings(names)

	return names
}
