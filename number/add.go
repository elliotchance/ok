package number

import (
	"math/big"
)

// Add returns the result of adding two numbers together. The result precision
// will be that of the highest precision input.
func Add(a, b *Number) *Number {
	return &Number{
		Rat:       new(big.Rat).Add(a.Rat, b.Rat),
		Precision: max(a.Precision, b.Precision),
	}
}
