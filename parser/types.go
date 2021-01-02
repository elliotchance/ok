package parser

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// ResolveTypes is a necessary part between the parsing finishing and the
// compiler starting. It is the job of this process to translate all the
// keywords that represent types (that would appear in argument types, return
// types, etc) into the actual types within this package.
//
// Types referring to an external types (ie. "foo.Bar") are also resolved here,
// and all imports are expected to be provided to this function.
func (parser *Parser) ResolveTypes(
	registry types.Registry,
	imports map[string]*types.Type,
) error {
	for _, funcName := range parser.SortedFuncNames() {
		fn := parser.funcs[funcName]
		err := parser.resolveTypes(fn, registry, imports)
		if err != nil {
			return err
		}
	}

	for key := range parser.finalizers {
		for i := range parser.finalizers[key] {
			err := parser.resolveTypes(parser.finalizers[key][i], registry, imports)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (parser *Parser) resolveTypes(
	node ast.Node,
	registry types.Registry,
	imports map[string]*types.Type,
) error {
	elem := reflect.ValueOf(node).Elem()
	if !elem.IsValid() {
		return nil
	}

	var err error
	for i := 0; i < elem.NumField(); i++ {
		switch t := elem.Field(i).Interface().(type) {
		case ast.Node:
			err = parser.resolveTypes(t, registry, imports)
			if err != nil {
				return err
			}

		case []ast.Node:
			for _, n2 := range t {
				err = parser.resolveTypes(n2, registry, imports)
				if err != nil {
					return err
				}
			}

		case *types.Type:
			t, err = parser.ResolveType(node, t, registry, imports)
			if err != nil {
				return err
			}

		case []*ast.Argument:
			for _, arg := range t {
				arg.Type, err = parser.ResolveType(node, arg.Type, registry,
					imports)
				if err != nil {
					return err
				}
			}

		case []*types.Type:
			for i := range t {
				t[i], err = parser.ResolveType(node, t[i], registry, imports)
				if err != nil {
					return err
				}
			}

		case []*ast.On:
			for i := range t {
				err = parser.resolveTypes(t[i], registry, imports)
				if err != nil {
					return err
				}
			}

		case []*ast.KeyValue:
			for i := range t {
				err = parser.resolveTypes(t[i], registry, imports)
				if err != nil {
					return err
				}
			}

		case []*ast.Case:
			for i := range t {
				err = parser.resolveTypes(t[i], registry, imports)
				if err != nil {
					return err
				}
			}

		case nil, bool, string, int,
			[]*ast.Literal, map[string]*ast.Literal,
			*os.File, *bufio.Reader:
			// These are all types that may appear in the parsers AST nodes that
			// we can ignore during resolving interfaces.

		default:
			// TODO(elliot): This is here as a safety net. At some point in the
			//  future it could be removed as the AST will be stable enough that
			//  won't be reachable anymore.
			panic(fmt.Sprintf("%T", t))
		}
	}

	return nil
}

func (parser *Parser) ResolveType(
	node ast.Node,
	typ *types.Type,
	registry types.Registry,
	imports map[string]*types.Type,
) (*types.Type, error) {
	// There may not be a type in some cases, such as an array that doesn't have
	// an explicit type. This is fine.
	if typ == nil {
		return nil, nil
	}

	var err error

	switch typ.Kind {
	case types.KindArray, types.KindMap:
		typ.Element, err = parser.ResolveType(node, typ.Element, registry, imports)
		if err != nil {
			return nil, err
		}

	case types.KindFunc:
		for i := range typ.Arguments {
			typ.Arguments[i], err = parser.ResolveType(node, typ.Arguments[i],
				registry, imports)
			if err != nil {
				return nil, err
			}
		}

		for i := range typ.Returns {
			typ.Returns[i], err = parser.ResolveType(node, typ.Returns[i],
				registry, imports)
			if err != nil {
				return nil, err
			}
		}

		// TODO(elliot): Is this even needed?
		for _, i := range typ.SortedPropertyNames() {
			typ.Properties[i], err = parser.ResolveType(node, typ.Properties[i],
				registry, imports)
			if err != nil {
				return nil, err
			}
		}

	case types.KindUnresolvedInterface:
		// Check for imported type.
		parts := strings.Split(typ.Name, ".")
		if len(parts) == 2 {
			// TODO(elliot): Check all these exist in the path.
			return imports[parts[0]].Properties[parts[1]].Returns[0], nil
		}

		// Find the constructor.
		var constructorFn *ast.Func
		for _, fn := range parser.funcs {
			if fn.IsConstructor() && fn.Name == typ.Name {
				constructorFn = fn
				break
			}
		}

		if constructorFn == nil {
			return nil, fmt.Errorf("cannot find constructor for %s", typ.Name)
		}

		ty, err := constructorFn.Interface()
		if err != nil {
			return typ, fmt.Errorf("%v %s", node.Position(), err)
		}

		return types.NewInterface(typ.Name, ty), nil
	}

	return typ, nil
}
