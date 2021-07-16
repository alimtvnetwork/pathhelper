package copyrecursive

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string, fileMode os.FileMode) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}

	defer source.Close()

	// Create all the parent folder if needed
	if err := createDir(filepath.Dir(dst), fileMode); err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destination.Close()
	_, err = io.Copy(destination, source)

	return err
}
