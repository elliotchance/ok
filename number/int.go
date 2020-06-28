package number

import (
	"github.com/cockroachdb/apd/v2"
)

// Int64 returns the int64 value. This is only for internal use when we need the
// actual int for things like array sizes. Non-integers will be truncated
// (rounded down).
func Int64(x *apd.Decimal) int64 {
	// Extract the fractional part and ignore it. If we don't do this Int64 will
	// throw an error.
	integ, frac := new(apd.Decimal), new(apd.Decimal)
	x.Modf(integ, frac)

	i64, _ := integ.Int64()

	return i64
}

// Int returns the int value. This is only for internal use when we need the
// actual int for things like array lengths. Non-integers will be truncated
// (rounded down).
func Int(n *apd.Decimal) int {
	return int(Int64(n))
}
