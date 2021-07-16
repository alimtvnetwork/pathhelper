package copyrecursive

import (
	"fmt"
	"os"
)

func createDir(dir string, perm os.FileMode) error {
	err := os.MkdirAll(dir, perm)

	if err != nil {
		return fmt.Errorf(
			"failed to create directory: '%s' error: '%s'",
			dir,
			err.Error())
	}

	return nil
}
