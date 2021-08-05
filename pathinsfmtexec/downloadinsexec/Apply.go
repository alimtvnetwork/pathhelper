package downloadinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/createdir"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(download *pathinsfmt.Download) *errorwrapper.Wrapper {
	if download == nil {
		return errnew.EmptyPtr
	}

	createErr := createdir.RemoveCreateAll(
		download.IsCreateDir,
		true,
		download.IsClearDir,
		download.Destination,
		download.FileModeDir)

	if createErr.HasError() {
		return createErr
	}

	return errcmd.
		BashArgsErrorWrapper(
			Aria2C,
			download.URL,
			HyphenD,
			download.Destination,
			HyphenO,
			download.FileName)
}
