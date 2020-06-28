package number

import (
	"github.com/cockroachdb/apd/v2"
)

// Pow returns the result of applying a to the power of b.
func Pow(a, b *apd.Decimal) *apd.Decimal {
	d := new(apd.Decimal)
	_, _ = context.Pow(d, a, b)

	return fix(d)
}
