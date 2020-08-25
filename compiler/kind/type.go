package kind

import "strings"

// IsArray tests for an array kind.
func IsArray(ty string) bool {
	return strings.HasPrefix(ty, "[]")
}

// IsMap tests for a map kind.
func IsMap(ty string) bool {
	return strings.HasPrefix(ty, "{}")
}

// IsFunc tests for a function literal.
func IsFunc(ty string) bool {
	return strings.HasPrefix(ty, "func(")
}

// ElementType returns the element type of a value of an array or map.
func ElementType(ty string) string {
	return ty[2:]
}

// IsObject returns true for an object type.
func IsObject(ty string) bool {
	return !IsLiteral(ty) && !IsMap(ty) && !IsArray(ty)
}

// IsLiteral return true for any literal (including functions).
func IsLiteral(ty string) bool {
	return ty == "bool" ||
		ty == "char" ||
		ty == "data" ||
		ty == "number" ||
		ty == "string" ||
		IsFunc(ty)
}
