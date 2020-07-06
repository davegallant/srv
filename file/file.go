// Package file contains filesystem functions
package file

import (
	"github.com/spf13/afero"
)

// Exists returns true if a file exists
func Exists(filename string) bool {
	var AppFs = afero.NewOsFs()
	_, err := AppFs.Stat(filename)
	return err == nil
}
