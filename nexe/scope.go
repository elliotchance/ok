package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/types"
)

type Scope struct {
	parent          *Scope
	variables       map[string]*types.Type
	packageVariable string
	prefix          string
	funcs           []*Func
}

func NewScope() *Scope {
	return &Scope{
		variables: make(map[string]*types.Type),
	}
}

func (s *Scope) PackageScope(packageVariable string) *Scope {
	return &Scope{
		parent:          s,
		variables:       make(map[string]*types.Type),
		packageVariable: packageVariable,
	}
}

func (s *Scope) Subscope() *Scope {
	return &Scope{
		parent:          s,
		variables:       make(map[string]*types.Type),
		packageVariable: s.packageVariable,
		prefix:          s.prefix + "_",
	}
}

func (s *Scope) Get(name string) (string, *types.Type, error) {
	if strings.HasPrefix(name, "^") {
		return s.parent.Get(name[1:])
	}

	if ty, ok := s.variables[name]; ok {
		return s.realName(name), ty, nil
	}

	// We can't walk up the scope (like some other languages do). However, the
	// variable is allowed to also be stored at the root.
	if ty, ok := s.packageRoot().variables[name]; ok {
		return s.PackageVariable() + "[0]." + name, ty, nil
	}

	// If that fails, we can look at the true root which contains the built-in
	// language functions/definitions.
	if ty, ok := s.root().variables[name]; ok {
		return name, ty, nil
	}

	return "", nil, fmt.Errorf("no such variable: %s", name)
}

func (s *Scope) root() *Scope {
	if s.parent != nil {
		return s.parent.root()
	}

	return s
}

func (s *Scope) packageRoot() *Scope {
	if s.parent != nil && s.parent.parent != nil {
		return s.parent.packageRoot()
	}

	return s
}

func (s *Scope) Add(name string, ty *types.Type) string {
	s.variables[name] = ty

	return s.realName(name)
}

func (s *Scope) PackageVariable() string {
	return s.packageVariable
}

func (s *Scope) realName(name string) string {
	return s.prefix + name
}

func (s *Scope) String() string {
	return strings.Join(s.string(0), "\n")
}

func (s *Scope) string(depth int) []string {
	var lines []string
	for name, ty := range s.variables {
		lines = append(lines, fmt.Sprintf("%s%s (%s) %s",
			indentString(depth), name, s.realName(name), ty))
	}

	if s.parent != nil {
		lines = append(lines, s.parent.string(depth+1)...)
	}

	return lines
}

func (s *Scope) RegisterFunction(fn *Func) string {
	return s.packageRoot().registerFunction(fn)
}

func (s *Scope) registerFunction(fn *Func) string {
	uniqueName := fmt.Sprintf("_fn%d", len(s.funcs))
	s.funcs = append(s.funcs, fn)
	return uniqueName
}
