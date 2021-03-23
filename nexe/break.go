package nexe

import (
	"github.com/elliotchance/ok/types"
)

func TranspileBreak() (Expr, []*types.Type, error) {
	return "break", nil, nil
}
