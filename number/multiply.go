package number

import (
	"math/big"
)

// Multiply returns the product of two numbers. The result precision will be
// that of the highest precision input.
func Multiply(a, b *Number) *Number {
	return &Number{
		Rat:       new(big.Rat).Mul(a.Rat, b.Rat),
		Precision: max(a.Precision, b.Precision),
	}
}
