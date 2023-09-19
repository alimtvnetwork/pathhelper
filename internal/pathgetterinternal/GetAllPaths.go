package pathgetterinternal

import (
	"io/ioutil"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/normalizeinternal"
)

func GetAllPaths(isFixPaths bool, separator, rootPath string) *errstr.Results {
	if rootPath == constants.EmptyString {
		return errstr.Empty.Results()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.PathStatusCannotRead,
					err,
					rootPath))
	}

	slice := make(
		[]string,
		len(fileInfos))

	for i, info := range fileInfos {
		currentPath := rootPath +
			separator +
			info.Name()

		slice[i] = normalizeinternal.JoinPathsFixIf(
			isFixPaths, currentPath)
	}

	return errstr.New.Results.Strings(
		slice)
}
