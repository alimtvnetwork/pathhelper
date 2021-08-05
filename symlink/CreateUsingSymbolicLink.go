package symlink

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func CreateUsingSymbolicLink(symLink *pathinsfmt.SymbolicLink) *errorwrapper.Wrapper {
	if symLink == nil {
		return errnew.EmptyPtr
	}

	isFileExist := fsinternal.IsPathExists(symLink.Src)
	isFileMissing := !isFileExist

	if symLink.IsSkipOnSrcMissing && isFileMissing {
		return errnew.EmptyPtr
	}

	if !symLink.IsSkipOnSrcMissing && isFileMissing {
		return errnew.PathMessages(
			errtype.PathNotFound,
			symLink.Src,
			"Cannot apply or create symbolic link when path is missing. Please select IsSkipOnSrcMissing to ignore.")
	}

	if symLink.IsClearBefore {
		errW := fsinternal.SafeRemove(symLink.Dst)

		if errW.HasError() {
			return errW
		}
	}

	if symLink.IsSkipOnExist && isFileExist {
		return errnew.EmptyPtr
	}

	if symLink.IsMkDirAll {
		errW := fsinternal.CreateDirectoryAllUptoParentDefault(
			symLink.Dst)

		if errW.HasError() {
			return errW
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
		return errnew.NewRefs(
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

	return errnew.EmptyPtr
}
