package types

import (
	"regexp"
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
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "func") {
		ty := &Type{
			Kind: KindFunc,
		}

		// TODO(elliot): This is a bad solution. It doesn't work with nested
		//  func types. For that we will need to use a mini parser.
		parts := regexp.MustCompile(`\((.*?)\)\s*(\(?.*\)?)`).
			FindStringSubmatch(s)

		if strings.TrimSpace(parts[1]) != "" {
			for _, a := range StringSliceMap(strings.Split(parts[1], ","), strings.TrimSpace) {
				ty.Arguments = append(ty.Arguments, TypeFromString(a))
			}
		}

		parts[2] = strings.Trim(parts[2], "()")
		if strings.TrimSpace(parts[2]) != "" {
			for _, a := range StringSliceMap(strings.Split(parts[2], ","), strings.TrimSpace) {
				ty.Returns = append(ty.Returns, TypeFromString(a))
			}
		}

		return ty
	}

	if strings.HasPrefix(s, "[]") {
		return &Type{
			Kind:    KindArray,
			Element: TypeFromString(s[2:]),
		}
	}

	if strings.HasPrefix(s, "{}") {
		return &Type{
			Kind:    KindMap,
			Element: TypeFromString(s[2:]),
		}
	}

	t := &Type{
		Kind: kindFromString(s),
	}

	if t.Kind == KindUnresolvedInterface {
		t.Name = s
	}

	return t
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
