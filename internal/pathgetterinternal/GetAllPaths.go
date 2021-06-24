package pathgetterinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
)

func GetAllPaths(isFixPaths bool, separator, rootPath string) *errstr.Results {
	if rootPath == "" {
		return &errstr.Results{
			Values:       &[]string{},
			ErrorWrapper: errnew.EmptyPtr,
		}
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return &errstr.Results{
			Values: &[]string{},
			ErrorWrapper: errnew.Path(
				errtype.PathStatusCannotRead,
				err,
				rootPath),
		}
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

	return &errstr.Results{
		Values:       &slice,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
