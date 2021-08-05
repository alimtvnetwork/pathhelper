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

	if download.IsCreateDir {
		createdir.RemoveCreateAll(
			true,
			download.IsClearDir,
			download.Destination,
			download.FileModeDir)
	}

	ariaCommand := errcmd.ArgsJoin(
		Aria2C,
		download.URL,
		HyphenD,
		download.Destination,
		HyphenO,
		download.FileName,
	)

	return errcmd.
		BashScripts(ariaCommand).
		CompiledErrorWrapper()
}
