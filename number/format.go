package number

import (
	"strings"

	"github.com/cockroachdb/apd/v2"
)

// Format returns the string representation with the provided precision. If
// precision is less than 0 then the number is returned exactly. Otherwise, the
// number will be rounded or zero padded to the precision provided.
func Format(n *apd.Decimal, precision int) string {
	d := new(apd.Decimal).Set(n)

	if precision == 0 {
		_, _ = apd.BaseContext.RoundToIntegralExact(d, d)
	} else if precision > 0 {
		p := place(n)
		var scale int

		switch {
		case p == 0: // value = 0
			scale = p + precision

		case p > 0: // value > 1.
			scale = p + precision + 1

		case p < 0: // 0 < value < 1.
			scale = precision - p - 2
		}

		_, _ = apd.BaseContext.WithPrecision(uint32(scale)).Round(d, n)
	}

	parts := strings.Split(d.String()+".", ".")

	l := len(parts[1])
	switch {
	case precision < 0:
		// Trim off any extra trailing zeros.
		parts[1] = strings.TrimRight(parts[1], "0")

	case l < precision:
		// Pad with extra zeros to reach precision.
		parts[1] += strings.Repeat("0", precision-l)
	}

	if parts[1] == "" {
		return parts[0]
	}

	return parts[0] + "." + parts[1]
}

// place is used in conjunction with precision to correctly round.
//
// It determines at what position the first significant figure appears as
// compared to the decimal point. The negative sign is ignored.
//
//   place(123)                 = 3
//   place(12.3) = place(12)    = 2
//   place(1)                   = 1
//   place(0)                   = 0
//   place(0.1)                 = -1
//   place(0.053) = place(0.05) = -2
//
func place(n *apd.Decimal) int {
	d := new(apd.Decimal).Set(n)
	d.Abs(d)

	// Zero is a special case, not only because we cannot take the log.
	if d.IsZero() {
		return 0
	}

	_, _ = context.Log10(d, d)

	return Int(d)
}
