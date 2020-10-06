package file

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/afero"
)

var FakeFs = afero.NewOsFs()

func TestExists(t *testing.T) {
	t.Run("expects file to exist", func(t *testing.T) {
		input, _ := afero.TempFile(FakeFs, ".", "")
		defer removeFile(input.Name(), t)

		got := Exists(input.Name())
		expect := true

		if got != expect {
			t.Errorf("Expected file %s to exist.", input.Name())
		}
	})

	t.Run("expects file to not exist", func(t *testing.T) {
		input := fmt.Sprintf("test_file_%d", time.Now().UnixNano())

		got := Exists(input)
		expect := false

		if got != expect {
			t.Errorf("Expected file %s to not exist", input)
		}
	})
}

func removeFile(name string, t *testing.T) {
	err := FakeFs.Remove(name)
	if err != nil {
		t.Errorf("Did not cleanly delete file %s", name)
	}
}
