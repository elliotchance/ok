package types

import (
	"strings"
)

// Type represents a data type.
type Type struct {
	Kind Kind

	// Name is used as the descriptive name for the object.
	Name string

	// Element is used when Kind is an Array or Map.
	Element *Type

	// Argument and Returns are used when Kind is a Func. Either may be nil.
	Arguments, Returns []*Type

	// Properties is used for KindInterface
	Properties map[string]*Type
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
