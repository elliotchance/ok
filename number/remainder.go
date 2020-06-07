package number

import (
	"errors"
	"math/big"
)

// Remainder returns the remainder from the division of two numbers. This is not
// to be confused with a modulo which also returns the remainder but as an
// absolute number. The remainder may be negative if onf of the inputs are.
//
// The result precision will be that of the highest precision input. If b == 0
// an error will be returned and the result will be zero.
func Remainder(a, b *Number) (*Number, error) {
	if b.IsZero() {
		return NewNumber("0"), errors.New("division by zero")
	}

	// It's important we always ensure the times is rounded down. "123.7" will
	// be manually truncated to "123" since we know FloatString will always
	// return a fixed amount of digits.
	times := new(big.Rat).Quo(a.Rat, b.Rat).FloatString(1)
	times = times[:len(times)-2]

	offset := Multiply(NewNumber(times), b)

	return Subtract(a, offset), nil
}
