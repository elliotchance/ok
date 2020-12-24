package types_test

import (
	"testing"

	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestRegistry_Add(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.Number)
		assert.Equal(t, types.Registry{
			"0": types.Number,
		}, registry)
	})

	t.Run("multiple", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.Number)
		registry.Add(types.String)
		assert.Equal(t, types.Registry{
			"0": types.Number,
			"1": types.String,
		}, registry)
	})

	t.Run("duplicate", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.Number)
		registry.Add(types.String)
		registry.Add(types.Number)
		assert.Equal(t, types.Registry{
			"0": types.Number,
			"1": types.String,
		}, registry)
	})

	t.Run("array", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NumberArray)
		assert.Equal(t, types.Registry{
			"0": types.Number,
			"1": types.NewArray(types.NewRef("0")),
		}, registry)
	})

	t.Run("map", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.StringMap)
		assert.Equal(t, types.Registry{
			"0": types.String,
			"1": types.NewMap(types.NewRef("0")),
		}, registry)
	})

	t.Run("duplicate func", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewFunc(nil, nil))
		registry.Add(types.NewFunc(nil, nil))
		assert.Equal(t, types.Registry{
			"0": types.NewFunc(nil, nil),
		}, registry)
	})

	t.Run("not duplicate func args", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewFunc(nil, nil))
		registry.Add(types.NewFunc([]*types.Type{types.Number}, nil))
		assert.Equal(t, types.Registry{
			"0": types.NewFunc(nil, nil),
			"1": types.Number,
			"2": types.NewFunc([]*types.Type{types.NewRef("1")}, nil),
		}, registry)
	})

	t.Run("not duplicate func returns", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewFunc(nil, []*types.Type{types.Number}))
		registry.Add(types.NewFunc(nil, nil))
		assert.Equal(t, types.Registry{
			"0": types.Number,
			"1": types.NewFunc(nil, []*types.Type{types.NewRef("0")}),
			"2": types.NewFunc(nil, nil),
		}, registry)
	})

	t.Run("duplicate interface", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewInterface("Foo", nil))
		registry.Add(types.NewInterface("Foo", nil))
		assert.Equal(t, types.Registry{
			"0": types.NewInterface("Foo", nil),
		}, registry)
	})

	t.Run("not duplicate interface", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewInterface("Foo", nil))
		registry.Add(types.NewInterface("Bar", map[string]*types.Type{
			"Bar": types.String,
		}))
		assert.Equal(t, types.Registry{
			"0": types.NewInterface("Foo", nil),
			"1": types.String,
			"2": types.NewInterface("Bar", map[string]*types.Type{
				"Bar": types.NewRef("1"),
			}),
		}, registry)
	})

	t.Run("duplicate interface with different name", func(t *testing.T) {
		registry := types.Registry{}
		registry.Add(types.NewInterface("Bar", map[string]*types.Type{
			"Foo": types.String,
		}))
		registry.Add(types.NewInterface("Bar", map[string]*types.Type{
			"Foo": types.String,
		}))
		assert.Equal(t, types.Registry{
			"0": types.String,
			"1": types.NewInterface("Bar", map[string]*types.Type{
				"Foo": types.NewRef("0"),
			}),
		}, registry)
	})
}

func TestRegistry_Get(t *testing.T) {
	registry := types.Registry{
		"0": types.Number,
		"1": types.String,
		"2": types.NewFunc(nil, []*types.Type{types.NewRef("0")}),
		"3": types.NewArray(types.NewRef("0")),
		"4": types.NewMap(types.NewRef("1")),
	}

	t.Run("simple type", func(t *testing.T) {
		assert.Equal(t, types.String, registry.Get("1"))
	})

	t.Run("array", func(t *testing.T) {
		assert.Equal(t, types.NumberArray, registry.Get("3"))
	})

	t.Run("map", func(t *testing.T) {
		assert.Equal(t, types.StringMap, registry.Get("4"))
	})

	t.Run("func", func(t *testing.T) {
		assert.Equal(t,
			types.NewFunc(nil, []*types.Type{types.Number}),
			registry.Get("2"))
	})
}
