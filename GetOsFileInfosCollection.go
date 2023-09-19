package pathhelper

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"

	"gitlab.com/auk-go/pathhelper/internal/fileinfogetter"
	"gitlab.com/auk-go/pathhelper/osfileinfos"
)

func GetOsFileInfosCollection(
	allPaths *[]string,
) (
	infos *osfileinfos.Collection,
	errsCollection *errwrappers.Collection,
) {
	rawInfos, errWrappersCollection := fileinfogetter.GetWithErrors(allPaths)

	return osfileinfos.New(rawInfos), errWrappersCollection
}
