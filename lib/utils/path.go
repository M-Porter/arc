package utils

import (
	"fmt"
	"os"
)

// PathExists returns true if the given path exists.
// Returns an error if one occurred.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("error checking if path is valid and exists: %v", err)
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}
