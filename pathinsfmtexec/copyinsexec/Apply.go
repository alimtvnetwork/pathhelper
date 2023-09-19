package copyinsexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/pathmodifier"
)

func Apply(copyPath *pathinsfmt.CopyPath) *errorwrapper.Wrapper {
	if copyPath == nil {
		return nil
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

	return nil
}
