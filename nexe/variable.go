package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/types"
)

type Variable struct {
	Name string
	Type *types.Type
}

func (v *Variable) JS(indent int) string {
	return fmt.Sprintf("let %s", v.Name)
}
