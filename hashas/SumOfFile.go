package hashas

import (
	"io"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SumOfFile(method Variant, filePath string) *errbyte.Results {
	if filePath == constants.EmptyString {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyString,
				"File name is empty"))
	}

	isExist, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(filePath)

	if !isExist || fileInfo == nil || fileInfo.IsDir() {
		return errbyte.EmptyResultsWithError(
			errnew.PathMessages(
				errtype.InvalidPath,
				filePath,
				"File path either invalid or has permission issue or a folder for hash-checksum."))
	}

	hashWriter, errWp := method.NewHash()

	if errWp.HasError() {
		return errbyte.EmptyResultsWithError(errWp)
	}

	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		return errbyte.EmptyResultsWithError(
			errnew.Path(
				errtype.FileRead,
				errOpen,
				"Error opening file : "+filePath,
			))
	}

	defer file.Close()

	_, errCopy := io.Copy(hashWriter, file)
	if errCopy != nil {
		return errbyte.EmptyResultsWithError(
			errnew.Path(
				errtype.Copy,
				errOpen,
				"Error copying to  file : "+filePath,
			))
	}

	hashedBytes := hashWriter.Sum(nil)

	return errbyte.EmptyErrorResults(hashedBytes...)
}
