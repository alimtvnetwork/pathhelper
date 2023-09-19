package pathhelper

import (
	"os"

	"gitlab.com/auk-go/errorwrapper/errwrappers"

	"gitlab.com/auk-go/pathhelper/internal/fileinfogetter"
)

func GetOsFileInfos(
	allPaths []string,
) (
	infos []os.FileInfo,
	errsCollection *errwrappers.Collection,
) {
	return fileinfogetter.GetWithErrors(allPaths)
}
