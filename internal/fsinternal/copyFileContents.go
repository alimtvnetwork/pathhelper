package fsinternal

import (
	"io"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdefer"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyFileContents(srcPath, dstPath string) (errWp *errorwrapper.Wrapper) {
	inFile, errOpen := os.Open(srcPath)

	defer errdefer.CloseFile(
		srcPath,
		errWp,
		inFile)

	if errOpen != nil {
		return errnew.Path(
			errtype.FileRead,
			errOpen,
			srcPath)
	}

	outFile, errCreate := os.Create(dstPath)

	defer errdefer.CloseFile(
		dstPath,
		errWp,
		outFile)

	if errCreate != nil {
		return errnew.Path(
			errtype.FileRead,
			errCreate,
			dstPath)
	}

	if _, err := io.Copy(outFile, inFile); err != nil {
		return errnew.SourceDestinationMessages(
			errtype.Copy,
			srcPath,
			dstPath,
			err.Error())
	}

	err2 := outFile.Sync()

	if err2 != nil {
		return errnew.SourceDestinationMessages(
			errtype.Sync,
			srcPath,
			dstPath,
			err2.Error(),
			errtype.Copy.String())
	}

	return errnew.EmptyPtr
}
