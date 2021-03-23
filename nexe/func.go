package nexe

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

type Func struct {
	UniqueName string
	Name       string
	Args       []string
	Stmts      []Stmt
	Type       *types.Type
	Pos        string
}

func (f *Func) JS(indent int) string {
	var lines []string
	for _, stmt := range f.Stmts {
		lines = append(lines, indentString(indent)+stmt.JS(indent)+";")
	}

	return fmt.Sprintf("function %s(%s) {\n%s\n%s}",
		f.UniqueName, strings.Join(f.Args, ", "),
		strings.Join(lines, "\n"), indentString(indent-1))
}

func TranspileFunc(fn *ast.Func, scope *Scope) (*Func, []*types.Type, error) {
	scope = scope.Subscope()

	f := &Func{
		Name: fn.Name,
		Type: fn.Type(),
		Pos:  fn.Position(),
	}
	f.UniqueName = scope.RegisterFunction(f)

	for _, arg := range fn.Arguments {
		realArgName := scope.Add(arg.Name, arg.Type)
		f.Args = append(f.Args, realArgName)
	}

	// First parse through to register all of the variables.
	for _, stmt := range fn.Statements {
		switch s := stmt.(type) {
		case *ast.Assign:
			// An assignment falls into one of two syntax, either the number of
			// elements on both sides is equal, or the right only has a single
			// expression that returns the same amount of types as the left.
			switch {
			case len(s.Lefts) == len(s.Rights):
				for i := range s.Lefts {
					if left, ok := s.Lefts[i].(*ast.Identifier); ok {
						name := left.Name

						// If this is a function assignment, either hoisted or assigning
						// a function literal we do not need to traverse because we
						// already know the type. Also, if we were go traverse now, we
						// might get a compiler error if this closure references
						// variables in the parent scope we haven't seen yet.
						if fn, ok := s.Rights[i].(*ast.Func); ok {
							scope.Add(name, fn.Type())

							f.Stmts = append(f.Stmts, &Variable{
								Name: scope.realName(name),
								Type: fn.Type(),
							})

							continue
						}

						if _, ty, _ := scope.Get(name); ty != nil {
							continue
						}

						_, ty, err := TranspileStmt(s.Rights[i], scope)
						if err != nil {
							return nil, nil, err
						}

						scope.Add(name, ty[0])
						f.Stmts = append(f.Stmts, &Variable{
							Name: scope.realName(name),
							Type: ty[0],
						})
					}
				}

			case len(s.Rights) == 1:
				_, tys, err := TranspileStmt(s.Rights[0], scope)
				if err != nil {
					return nil, nil, err
				}

				for i, ty := range tys {
					name := s.Lefts[i].(*ast.Identifier).Name
					scope.Add(name, ty)

					f.Stmts = append(f.Stmts, &Variable{
						Name: scope.realName(name),
						Type: ty,
					})
				}
			}
		}
	}

	for _, stmt := range fn.Statements {
		s, _, err := TranspileStmt(stmt, scope)
		if err != nil {
			return nil, nil, err
		}

		f.Stmts = append(f.Stmts, s)
	}

	if fn.IsConstructor() {
		var defs []string
		for name := range scope.variables {
			defs = append(defs, fmt.Sprintf("%s: %s", name, scope.realName(name)))
		}
		f.Stmts = append(f.Stmts, Expr(fmt.Sprintf("return [[{%s},'%s']]",
			strings.Join(defs, ", "), fn.Name)))
	} else if len(f.Stmts) == 0 {
		f.Stmts = append(f.Stmts, Expr("return []"))
	} else {
		if _, ok := f.Stmts[len(f.Stmts)-1].(*Return); !ok {
			f.Stmts = append(f.Stmts, Expr("return []"))
		}
	}

	return f, []*types.Type{fn.Type()}, nil
}
