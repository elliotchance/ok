package number

import (
	"github.com/cockroachdb/apd/v2"
)

// apdPrecision controls the accuracy of the calculations (in terms of
// significant figures). This value wasn't chosen for any specific reason other
// than it should be big enough to handle all reasonable cases without adding
// too much strain on the CPU.
const apdPrecision = 20

// context should be used for all calculations.
var context = apd.BaseContext.WithPrecision(apdPrecision)

// NewNumber creates a new number from a string that is well formatted.
func NewNumber(s string) *apd.Decimal {
	// We do not care about the error because we expect the number to be well
	// formatted.
	dec, _, _ := apd.NewFromString(s)

	return fix(dec)
}

func fix(dec *apd.Decimal) *apd.Decimal {
	// Correct for negative zeros.
	if dec.IsZero() {
		return new(apd.Decimal)
	}

	return dec
}

// IsZero returns true if the value is equivalent to zero (at any precision).
func IsZero(n *apd.Decimal) bool {
	return n.IsZero()
}
