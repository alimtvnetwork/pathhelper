package symlink

import (
	"os"

	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/ref"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
	"gitlab.com/auk-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func CreateUsingSymbolicLink(symLink *pathinsfmt.SymbolicLink) *errorwrapper.Wrapper {
	if symLink == nil {
		return nil
	}

	isFileExist := fsinternal.IsPathExists(symLink.Src)
	isFileMissing := !isFileExist

	if symLink.IsSkipOnSrcMissing && isFileMissing {
		return nil
	}

	if !symLink.IsSkipOnSrcMissing && isFileMissing {
		return errnew.
			Path.
			Messages(
				errtype.PathNotFound,
				symLink.Src,
				"Cannot apply or create symbolic link when path is missing. Please select IsSkipOnSrcMissing to ignore.")
	}

	if symLink.IsClearBefore {
		errWrap := fsinternal.SafeRemove(symLink.Dst)

		if errWrap.HasError() {
			return errWrap
		}
	}

	if symLink.IsSkipOnExist && isFileExist {
		return nil
	}

	if symLink.IsMkDirAll {
		errWrap := fsinternal.CreateDirectoryAllUptoParentDefault(
			symLink.Dst)

		if errWrap.HasError() {
			return errWrap
		}
	}

	dst := symLink.Dst

	if fsinternal.IsDirectory(dst) {
		sourceFileName := fsinternal.GetFileName(symLink.Src)
		dst += osconsts.PathSeparator + sourceFileName
		dst = normalizeinternal.Fix(dst)
	}

	err := os.Symlink(
		symLink.Src,
		dst)

	if err != nil {
		return errnew.Ref.ManyWithError(
			errtype.SymbolicLink,
			err,
			ref.Value{
				Variable: "Source Symbolic Link",
				Value:    symLink.Src,
			},
			ref.Value{
				Variable: "Destination Symbolic Link",
				Value:    dst,
			},
			ref.Value{
				Variable: "Full Symbolic Link Request",
				Value:    symLink,
			})
	}

	return nil
}
