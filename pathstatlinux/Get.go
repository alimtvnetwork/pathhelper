package pathstatlinux

import (
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func Get(location string) *Info {
	if !fsinternal.IsPathExists(location) {
		return InvalidInfo(location)
	}

	pathStat := errcmd.BashArgs("stat", location)
	errorWrapper := pathStat.CompiledErrorWrapper()

	if errorWrapper.HasError() {
		return InvalidInfoUsingErr(
			location,
			errorWrapper)
	}

	lines := *pathStat.CompiledTrimmedOutputLines()

	return ProcessLinesToInfo(
		lines,
		location,
		errorWrapper)
}
