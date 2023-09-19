package downloadinsexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func Apply(download *pathinsfmt.Download) *errorwrapper.Wrapper {
	if download == nil {
		return nil
	}

	if download.IsSkipOnExist && download.PathStat().IsExist {
		return nil
	}

	createErr := download.
		CreateDirInstruction().
		CreateDefault()

	if createErr.HasError() {
		return createErr
	}

	bashCommandArg := aria2cBashCommandArg(download)

	scriptRunningErr := errcmd.
		New.BashScript.ArgsErr(bashCommandArg)

	if scriptRunningErr.HasError() {
		return scriptRunningErr
	}

	return downloadChecksumVerify(download)
}
