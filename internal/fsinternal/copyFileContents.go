package fsinternal

import (
	"io"
	"os"
)

func copyFileContents(srcPath, dstPath string) (err error) {
	inFile, errOpen := os.Open(srcPath)
	if errOpen != nil {
		return errOpen
	}

	defer inFile.Close()

	outFile, errCreate := os.Create(dstPath)
	if errCreate != nil {
		return errCreate
	}

	defer func() {
		outCloseErr := outFile.Close()
		if err == nil {
			err = outCloseErr
		}
	}()

	if _, err = io.Copy(outFile, inFile); err != nil {
		return err
	}

	err = outFile.Sync()

	return err
}
