// Package file contains filesystem functions
package file

import (
	"os"
)

// Exists returns true if a file exists
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
