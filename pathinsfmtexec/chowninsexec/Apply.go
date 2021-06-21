package chowninsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
)

func Apply(
	isContinueOnError bool,
	chown *pathinsfmt.Chown,
	flatPaths []string,
) *errorwrapper.Wrapper {
	if chown == nil || len(flatPaths) == 0 {
		return errnew.EmptyPtr
	}

	return namegroup.Apply(
		chown.IsRecursive,
		isContinueOnError,
		&chown.BaseUserNamePlusGroupName,
		flatPaths...)
}
