package pathstatlinux

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func Get(location string) *Info {
	if !fsinternal.IsPathExists(location) {
		return InvalidInfo(location)
	}

	if osconsts.IsWindows {
		return InvalidInfoUsingErr(
			location,
			errnew.
				Path.
				Messages(
					errtype.NotSupportInWindows,
					location,
					"pathstatlinux package is not supported in windows."))
	}

	pathStat := errcmd.New.BashScript.ArgsDefault("stat", location)
	errorWrapper := pathStat.CompiledErrorWrapper()

	if errorWrapper.HasError() {
		return InvalidInfoUsingErr(
			location,
			errorWrapper)
	}

	lines := pathStat.CompiledTrimmedOutputLines()

	return ProcessLinesToInfo(
		lines,
		location,
		errorWrapper)
}
