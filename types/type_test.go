package types_test

import (
	"testing"

	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestTypeFromString(t *testing.T) {
	for typeString, tt := range map[string]*types.Type{
		// basic types
		"bool":    {Kind: types.KindBool},
		"char":    {Kind: types.KindChar},
		"data":    {Kind: types.KindData},
		"number":  {Kind: types.KindNumber},
		"string":  {Kind: types.KindString},
		" char":   {Kind: types.KindChar},
		"\tdata ": {Kind: types.KindData},

		// arrays and maps
		"[]bool":    {Kind: types.KindArray, Element: &types.Type{Kind: types.KindBool}},
		"{} string": {Kind: types.KindMap, Element: &types.Type{Kind: types.KindString}},

		// functions
		"func()": {
			Kind: types.KindFunc,
		},
		"func() number": {
			Kind:    types.KindFunc,
			Returns: []*types.Type{{Kind: types.KindNumber}},
		},
		"func (number)": {
			Kind:      types.KindFunc,
			Arguments: []*types.Type{{Kind: types.KindNumber}},
		},
		" func  (number) (string, bool)": {
			Kind:      types.KindFunc,
			Arguments: []*types.Type{{Kind: types.KindNumber}},
			Returns:   []*types.Type{{Kind: types.KindString}, {Kind: types.KindBool}},
		},
		"[]func()": {
			Kind:    types.KindArray,
			Element: &types.Type{Kind: types.KindFunc},
		},

		// interfaces/objects
		"Person": {Kind: types.KindUnresolvedInterface, Name: "Person"},
		"[] fooBar": {
			Kind: types.KindArray,
			Element: &types.Type{
				Kind: types.KindUnresolvedInterface,
				Name: "fooBar",
			},
		},
	} {
		t.Run(typeString, func(t *testing.T) {
			assert.Equal(t, tt, types.TypeFromString(typeString))
		})
	}
}

func TestType_String(t *testing.T) {
	for typeString, tt := range map[string]*types.Type{
		// basic types
		"any":    {Kind: types.KindAny},
		"bool":   {Kind: types.KindBool},
		"char":   {Kind: types.KindChar},
		"data":   {Kind: types.KindData},
		"number": {Kind: types.KindNumber},
		"string": {Kind: types.KindString},

		// arrays and maps
		"[]bool":   {Kind: types.KindArray, Element: &types.Type{Kind: types.KindBool}},
		"{}string": {Kind: types.KindMap, Element: &types.Type{Kind: types.KindString}},

		// functions
		"func()": {
			Kind: types.KindFunc,
		},
		"func() number": {
			Kind:    types.KindFunc,
			Returns: []*types.Type{{Kind: types.KindNumber}},
		},
		"func(number)": {
			Kind:      types.KindFunc,
			Arguments: []*types.Type{{Kind: types.KindNumber}},
		},
		"func(number) (string, bool)": {
			Kind:      types.KindFunc,
			Arguments: []*types.Type{{Kind: types.KindNumber}},
			Returns:   []*types.Type{{Kind: types.KindString}, {Kind: types.KindBool}},
		},

		// interfaces/objects
		"Person": {Kind: types.KindUnresolvedInterface, Name: "Person"},
		"[]fooBar": {
			Kind: types.KindArray,
			Element: &types.Type{
				Kind: types.KindUnresolvedInterface,
				Name: "fooBar",
			},
		},
		"PersonA": {Kind: types.KindResolvedInterface, Name: "PersonA"},
	} {
		t.Run(typeString, func(t *testing.T) {
			assert.Equal(t, typeString, tt.String())
		})
	}
}
