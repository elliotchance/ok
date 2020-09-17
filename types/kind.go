package types

type Kind int

const (
	// An unresolved interface (object type) is named, but the properties will
	// be nil because they are not known. This is a distinction from a resolved
	// interface which allowed to have zero or more properties.
	//
	// KindUnresolvedInterface = 0 on purpose it acts as a fallthrough/unknown
	// type.
	KindUnresolvedInterface Kind = iota
	KindResolvedInterface

	// Basic types.
	KindAny
	KindBool
	KindChar
	KindData
	KindNumber
	KindString

	// Other.
	KindArray
	KindMap
	KindFunc
)

func kindFromString(s string) Kind {
	switch s {
	case "any":
		return KindAny

	case "bool":
		return KindBool

	case "char":
		return KindChar

	case "data":
		return KindData

	case "number":
		return KindNumber

	case "string":
		return KindString
	}

	return KindUnresolvedInterface
}
