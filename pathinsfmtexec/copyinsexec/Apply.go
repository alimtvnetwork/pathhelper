package copyinsexec

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func Apply(copyPath *pathinsfmt.CopyPath) *errorwrapper.Wrapper {
	if copyPath == nil {
		return errnew.EmptyPtr
	}

	delCondition := deletepaths.Condition{
		IsRemove:           copyPath.IsClearBeforeCopy,
		IsRecursive:        copyPath.IsRecursive,
		IsExistBeforeClear: true,
	}

	destination := copyPath.DestinationFixedPath()

	delErr := deletepaths.
		SingleOrRecursiveOnCondition(
			delCondition,
			destination)

	if delErr.HasError() {
		return delErr
	}

	src := copyPath.SourceFixedPath()

	if !copyPath.IsRecursive {
		return fs.CopyFile(
			src,
			destination)
	}

	recursivePaths, errCollection := recursiveinternal.GetPaths(
		osconsts.PathSeparator,
		src,
		false)

	if errCollection.HasError() {
		return errCollection.GetAsErrorWrapperPtr()
	}

	copyErr := fs.CopyFile(src, destination)

	if copyErr.HasError() {
		return copyErr
	}

	for _, s := range *recursivePaths {
		if s == src {
			continue
		}

		relPath := strings.ReplaceAll(
			s,
			src,
			constants.EmptyString)

		if relPath == constants.EmptyString {
			continue
		}

		dest := pathjoin.JoinNormalizedIf(
			copyPath.IsNormalize,
			destination,
			relPath)

		createDir := fs.CreateDirectoryAllUptoParent(
			dest)

		if createDir.HasError() {
			return createDir
		}

		copyErr := fs.CopyFile(s, dest)

		if copyErr.HasError() {
			return copyErr
		}
	}

	return errnew.EmptyPtr
}
