package util

import "os"

// OKPath returns the root of where packages can be found. This can be provided
// through the $OKPATH environment variable. However, if its not set (or is
// empty) then the current directory will be used.
func OKPath() (string, error) {
	okPath := os.Getenv("OKPATH")
	if okPath == "" {
		var err error
		okPath, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	return okPath, nil
}
