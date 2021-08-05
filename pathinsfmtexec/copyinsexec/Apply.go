package copyinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifier"
)

func Apply(copyPath *pathinsfmt.CopyPath) *errorwrapper.Wrapper {
	if copyPath == nil {
		return errnew.EmptyPtr
	}

	copyErr := copyPath.LazyCopier().Copy()

	if copyErr.HasError() {
		return copyErr
	}

	if copyPath.HasApplyPathModifier() {
		return pathmodifier.ApplySimpleSingle(
			copyPath.Options.ApplyPathModifier,
			copyPath.DestinationFixedPath())
	}

	return errnew.EmptyPtr
}
