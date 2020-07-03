package util

import (
	"fmt"
	"os"
)

// CheckErrorsWithExit is used in several places to bail out after a compilation
// that contains failures.
func CheckErrorsWithExit(errs []error) {
	for _, err := range errs {
		fmt.Println(err)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}
