package number

import (
	"errors"

	"github.com/cockroachdb/apd/v2"
)

// Add returns the result of adding two numbers together.
func Add(a, b *apd.Decimal) *apd.Decimal {
	d := new(apd.Decimal)
	_, _ = context.Add(d, a, b)

	return fix(d)
}

// Subtract returns the result of subtracting two numbers together.
func Subtract(a, b *apd.Decimal) *apd.Decimal {
	d := new(apd.Decimal)
	_, _ = context.Sub(d, a, b)

	return fix(d)
}

// Multiply returns the product (multiplication) of two numbers.
func Multiply(a, b *apd.Decimal) *apd.Decimal {
	d := new(apd.Decimal)
	_, _ = context.Mul(d, a, b)

	return fix(d)
}

// Divide returns the division of two numbers. If b == 0 an error will be
// returned and the result is zero.
func Divide(a, b *apd.Decimal) (*apd.Decimal, error) {
	if b.IsZero() {
		return new(apd.Decimal), errors.New("division by zero")
	}

	d := new(apd.Decimal)
	_, _ = context.Quo(d, a, b)

	return fix(d), nil
}

// Remainder returns the remainder from the division of two numbers. This is not
// to be confused with a modulo which also returns the remainder but as an
// absolute number. The remainder may be negative if one of the inputs are.
//
// Like Divide, if b == 0 an error will be returned and the result will be zero.
func Remainder(a, b *apd.Decimal) (*apd.Decimal, error) {
	if b.IsZero() {
		return new(apd.Decimal), errors.New("division by zero")
	}

	d := new(apd.Decimal)
	_, _ = context.Rem(d, a, b)

	return fix(d), nil
}
