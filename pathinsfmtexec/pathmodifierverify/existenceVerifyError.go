package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func existenceVerifyError(
	isVerify bool,
	verifier *pathinsfmt.PathVerifier,
	location string,
) *errbool.Result {
	if !isVerify || verifier == nil || len(location) == 0 {
		return errbool.EmptyErrorResultPtr(
			false)
	}

	isFileExist := fsinternal.IsPathExists(location)
	isFileMissing := !isFileExist

	if verifier.IsSkipCheckingOnNonExist && isFileMissing {
		return errbool.EmptyErrorResultPtr(
			false)
	}

	if !verifier.IsSkipCheckingOnNonExist && isFileMissing {
		errWp := errnew.PathMessages(
			errtype.PathNotFound,
			location,
			"Use IsSkipCheckingOnNonExist to true skip the error.")

		return errbool.NewUsingWrapperPtr(
			isFileExist,
			errWp)
	}

	return errbool.EmptyErrorResultPtr(
		isFileExist)
}
