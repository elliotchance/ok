package number

import (
	"math/big"
)

// Subtract returns the result of subtracting two numbers together. The result
// precision will be that of the highest precision input.
func Subtract(a, b *Number) *Number {
	return &Number{
		Rat:       new(big.Rat).Sub(a.Rat, b.Rat),
		Precision: max(a.Precision, b.Precision),
	}
}
