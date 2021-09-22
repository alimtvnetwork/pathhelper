package fileinfopath

import (
	"os"
	"time"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompare"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

type Instance struct {
	FileInfo os.FileInfo
	FullPath string
	Error    error
}

func New(location string) *Instance {
	fileInfo, err := os.Stat(location)

	return &Instance{
		FileInfo: fileInfo,
		FullPath: location,
		Error:    err,
	}
}

func (it *Instance) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *Instance) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *Instance) IsInvalidFileInfo() bool {
	return it == nil || it.FileInfo == nil
}

func (it *Instance) HasFileInfo() bool {
	return it != nil && it.FileInfo != nil
}

func (it *Instance) IsFile() bool {
	return it.IsExist() && !it.FileInfo.IsDir()
}

func (it *Instance) IsDir() bool {
	return it.IsExist() && it.FileInfo.IsDir()
}

func (it *Instance) IsExist() bool {
	return it != nil && it.FileInfo != nil && (it.Error == nil || !os.IsNotExist(it.Error))
}

func (it *Instance) IsInvalidPath() bool {
	return it == nil || it.FileInfo == nil || it.Error != nil
}

// FileName
// Returns file name with extension
func (it *Instance) FileName() string {
	if it.HasFileInfo() {
		return it.FileInfo.Name()
	}

	return constants.EmptyString
}

func (it *Instance) FileNameWithoutExt() string {
	if it.HasFileInfo() {
		return splitinternal.GetFileNameWithoutExt(it.FileInfo.Name())
	}

	return constants.EmptyString
}

func (it *Instance) BothExtension() (dotExt, ext string) {
	if it.HasFileInfo() {
		return splitinternal.GetBothExtension(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *Instance) DotExtension() (fileName, dotExt string) {
	if it.HasFileInfo() {
		return splitinternal.GetFileNameDotExt(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *Instance) FileNameExt() string {
	return it.FileName()
}

func (it *Instance) Mode() os.FileMode {
	if it.HasFileInfo() {
		return it.FileInfo.Mode()
	}

	return constants.Zero
}

func (it *Instance) LastModifiedAt() *time.Time {
	if it.HasFileInfo() {
		mod := it.FileInfo.ModTime()

		return &mod
	}

	return nil
}

func (it *Instance) Size() *int64 {
	if it.HasFileInfo() {
		sz := it.FileInfo.Size()

		return &sz
	}

	return nil
}

func (it *Instance) CompareFileInfo(right os.FileInfo) corecomparator.Compare {
	return pathcompare.FileInfo(it.FileInfo, right)
}

func (it *Instance) CompareSize(anotherInstance *Instance) corecomparator.Compare {
	return pathcompare.Size(it.Size(), anotherInstance.Size())
}

func (it *Instance) CompareLastModified(anotherInstance *Instance) corecomparator.Compare {
	return pathcompare.LastModified(it.LastModifiedAt(), anotherInstance.LastModifiedAt())
}

func (it *Instance) NotFileError() *errorwrapper.Wrapper {
	if it.IsFile() {
		return errnew.EmptyPtr
	}

	return errnew.PathMessages(
		errtype.File,
		it.FullPath,
		"Cannot read invalid path or a directory. (required file)")
}

func (it *Instance) NotDirError() *errorwrapper.Wrapper {
	if it.IsDir() {
		return errnew.EmptyPtr
	}

	return errnew.PathMessages(
		errtype.Directory,
		it.FullPath,
		"Cannot read invalid path or a file. (required directory)")
}

func (it *Instance) ErrorWrapper(errType errtype.Variation) *errorwrapper.Wrapper {
	if it.HasError() {
		return errnew.Path(errType, it.Error, it.FullPath)
	}

	return errnew.EmptyPtr
}
