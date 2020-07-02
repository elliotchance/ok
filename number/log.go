package number

import (
	"github.com/cockroachdb/apd/v2"
)

// Log is the natural logarithm (base e).
func Log(x *apd.Decimal) *apd.Decimal {
	d := new(apd.Decimal)
	_, _ = context.Ln(d, x)

	return fix(d)
}
