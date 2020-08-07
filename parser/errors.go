package parser

import (
	"strings"
)

// Errors is a collection of errors.
type Errors []error

func (errs Errors) String() string {
	var ss []string
	for _, err := range errs {
		ss = append(ss, err.Error())
	}

	return strings.Join(ss, "; ")
}
