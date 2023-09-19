package chowninsexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/namegroup"
)

func Apply(
	isContinueOnError bool,
	chown *pathinsfmt.Chown,
	flatPaths []string,
) *errorwrapper.Wrapper {
	if chown == nil || len(flatPaths) == 0 {
		return nil
	}

	return namegroup.Apply(
		chown.IsRecursive,
		isContinueOnError,
		&chown.UserGroupName,
		flatPaths...)
}
