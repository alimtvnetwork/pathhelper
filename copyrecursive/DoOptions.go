package copyrecursive

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/ref"
	"gitlab.com/auk-go/pathhelper/deletepaths"
	"gitlab.com/auk-go/pathhelper/fs"
)

func DoOptions(
	src,
	dst string,
	options Options,
) *errorwrapper.Wrapper {
	var deleteErr *errorwrapper.Wrapper

	if options.IsClearDestination {
		deleteErr = deletepaths.RecursiveOnExist(
			dst)
	}

	if deleteErr.HasError() {
		return errnew.Ref.ManyUsingWrapper(deleteErr,
			ref.Value{
				Variable: "src",
				Value:    src,
			},
			ref.Value{
				Variable: "dst",
				Value:    dst,
			})
	}

	if isExists(dst) && options.IsSkipOnExist {
		return nil
	}

	isExist, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(src)
	if isExist && !fileInfo.IsDir() {
		// file
		return fs.CopyFile(src, dst)
	}

	if !isExist {
		return errnew.
			Path.
			Messages(
				errtype.PathMissingOrInvalid,
				src,
				"Source path is missing or access issues.",
				"Cannot copy to destination: ",
				dst)
	}

	err := DoSimple(
		src,
		dst)

	return errnew.Ref.TwoWithError(
		errtype.Copy,
		err,
		"src",
		src,
		"dst",
		dst)
}
