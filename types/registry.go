package types

import (
	"fmt"
)

// A Registry holds a collection of unique types that can are referenced
// internally to allow for circular references but also be serializable.
type Registry map[string]*Type

func (registry Registry) EqualTypeSlices(a, b []*Type) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !registry.EqualTypes(a[i], b[i]) {
			return false
		}
	}

	return true
}

// EqualTypes ignores only the Name.
func (registry Registry) EqualTypes(a, b *Type) bool {
	// This covers cases where we compare Element.
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Ref != "" {
		a = registry.Get(a.Ref)
	}

	if b.Ref != "" {
		b = registry.Get(b.Ref)
	}

	if a.Kind != b.Kind {
		return false
	}

	if !registry.EqualTypes(a.Element, b.Element) {
		return false
	}

	if !registry.EqualTypeSlices(a.Arguments, b.Arguments) {
		return false
	}

	if !registry.EqualTypeSlices(a.Returns, b.Returns) {
		return false
	}

	if len(a.Properties) != len(b.Properties) {
		return false
	}

	for i := range a.Properties {
		if !registry.EqualTypes(a.Properties[i], b.Properties[i]) {
			return false
		}
	}

	return true
}

// Add will merge new types into the register. An error is returned only if an
// unresolved interface type is added and there's no way to resolve it right
// now.
func (registry Registry) Add(ty *Type) (string, error) {
	for key, ty2 := range registry {
		if registry.EqualTypes(ty, ty2) {
			return key, nil
		}
	}

	ty = ty.Copy()

	// Flatten dependent types. After this process the root type and all
	// recursive types will only contain a Ref. It is rewired with Get().
	if ty.Element != nil {
		newType, err := registry.Add(ty.Element)
		if err != nil {
			return "", err
		}

		ty.Element = NewRef(newType)
	}

	for i := range ty.Returns {
		newType, err := registry.Add(ty.Returns[i])
		if err != nil {
			return "", err
		}

		ty.Returns[i] = NewRef(newType)
	}

	for i := range ty.Arguments {
		newType, err := registry.Add(ty.Arguments[i])
		if err != nil {
			return "", err
		}

		ty.Arguments[i] = NewRef(newType)
	}

	for _, i := range ty.SortedPropertyNames() {
		newType, err := registry.Add(ty.Properties[i])
		if err != nil {
			return "", err
		}

		ty.Properties[i] = NewRef(newType)
	}

	key := fmt.Sprintf("%d", len(registry))
	registry[key] = ty

	return key, nil
}

func (registry Registry) Get(key string) *Type {
	// Let the following code panic if the type key does not exist. Get is only
	// used internally and must be from existing TypeRegister values.
	ty := registry[key].Copy()

	// Unflatten types.
	if ty.Element != nil {
		ty.Element = registry.Get(ty.Element.Ref)
	}

	for i := range ty.Arguments {
		ty.Arguments[i] = registry.Get(ty.Arguments[i].Ref)
	}

	for i := range ty.Returns {
		ty.Returns[i] = registry.Get(ty.Returns[i].Ref)
	}

	for i := range ty.Properties {
		ty.Properties[i] = registry.Get(ty.Properties[i].Ref)
	}

	return ty
}
