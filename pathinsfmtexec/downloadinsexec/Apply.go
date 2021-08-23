package downloadinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(download *pathinsfmt.Download) *errorwrapper.Wrapper {
	if download == nil {
		return errnew.EmptyPtr
	}

	if download.IsSkipOnExist && download.PathStat().IsExist {
		return errnew.EmptyPtr
	}

	createErr := download.
		CreateDirInstruction().
		CreateDefault()

	if createErr.HasError() {
		return createErr
	}

	bashCommandArg := aria2cBashCommandArg(download)

	scriptRunningErr := errcmd.
		BashScriptsErrorWrapper(bashCommandArg)

	if scriptRunningErr.HasError() {
		return scriptRunningErr
	}

	return downloadChecksumVerify(download)
}
