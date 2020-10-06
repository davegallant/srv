package file

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestExists(t *testing.T) {
	t.Run("expects file to exist", func(t *testing.T) {
		var FakeFs = afero.NewOsFs()
		input, _ := afero.TempFile(FakeFs, ".", "")
		defer FakeFs.Remove(input.Name())

		got := Exists(input.Name())
		expect := true

		if got != expect {
			t.Errorf("Expected file %s to exist.", input.Name())
		}
	})

	t.Run("expects file to not exist", func(t *testing.T) {
		input := fmt.Sprintf("test_file_%d", time.Now().UnixNano())

		got := Exists(fmt.Sprintf(input))
		expect := false

		if got != expect {
			t.Errorf("Expected file %s to not exist", input)
		}
	})
}
