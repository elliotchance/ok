package nexe

import (
	"fmt"
	"sort"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Object struct {
	Elements map[Stmt]Stmt
	Type     *types.Type
}

func NewObject(ty *types.Type) *Object {
	return &Object{
		Elements: map[Stmt]Stmt{},
		Type:     ty,
	}
}

func (o *Object) JS(indent int) string {
	var lines []string
	for k, v := range o.Elements {
		lines = append(lines, fmt.Sprintf("%s%s: %s,", indentString(indent), k.JS(0), v.JS(indent+1)))
	}
	sort.Strings(lines)

	return fmt.Sprintf("[ {\n%s\n%s}, '%s' ]", strings.Join(lines, "\n"),
		indentString(indent-1), o.Type.String())
}

func TranspileMap(m *ast.Map, scope *Scope) (*Object, []*types.Type, error) {
	o := NewObject(types.AnyMap)
	for _, kv := range m.Elements {
		key, _, err := TranspileStmt(kv.Key, scope)
		if err != nil {
			return nil, nil, err
		}

		value, _, err := TranspileStmt(kv.Value, scope)
		if err != nil {
			return nil, nil, err
		}

		o.Elements[key] = value
	}

	return o, []*types.Type{o.Type}, nil
}
