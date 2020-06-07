package number

import (
	"errors"
	"math/big"
)

// Divide returns the division of two numbers. The result precision will be that
// of the highest precision input. If b == 0 an error will be returned and the
// result is zero.
func Divide(a, b *Number) (*Number, error) {
	if b.IsZero() {
		return NewNumber("0"), errors.New("division by zero")
	}

	return &Number{
		Rat:       new(big.Rat).Quo(a.Rat, b.Rat),
		Precision: max(a.Precision, b.Precision),
	}, nil
}
