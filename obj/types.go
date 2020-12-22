package obj

import (
	"fmt"

	"github.com/elliotchance/ok/types"
)

type Types map[string]*types.Type

func (s Types) Add(ty *types.Type) {
	key := fmt.Sprintf("%d", len(s))
	s[key] = ty
}
