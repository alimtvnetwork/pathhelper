package copyrecursive

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/ref"

	"gitlab.com/auk-go/pathhelper/deletepaths"
)

func Do(
	isClearBeforeCopy bool,
	src,
	dst string,
) *errorwrapper.Wrapper {
	var deleteErr *errorwrapper.Wrapper

	if isClearBeforeCopy {
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

	isExist, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(src)
	if isExist && !fileInfo.IsDir() {
		// file
		fileCopyErr := CopyFile(src, dst, defaultFileMode)

		if fileCopyErr == nil {
			return nil
		}

		return errnew.Ref.TwoWithError(
			errtype.Copy,
			fileCopyErr,
			"src",
			src,
			"dst",
			dst)
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

	if err == nil {
		return nil
	}

	return errnew.Ref.TwoWithError(
		errtype.Copy,
		err,
		"src",
		src,
		"dst",
		dst)
}
