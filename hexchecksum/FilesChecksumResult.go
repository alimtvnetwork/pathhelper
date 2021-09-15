package hexchecksum

import (
	"strconv"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

type FilesChecksumResult struct {
	HexFilesListChecksum     string
	HexFilesContentsChecksum string
	FilesCount               int
	Method                   hashas.Variant
	ErrorWrapper             *errorwrapper.Wrapper
}

func (it *FilesChecksumResult) CompileToSingle() *errstr.Result {
	if it == nil {
		return nil
	}

	var slice [5]string
	slice[0] = it.HexFilesListChecksum
	slice[1] = it.HexFilesContentsChecksum
	slice[2] = it.Method.Name()
	slice[3] = it.ErrorWrapper.String()
	slice[4] = strconv.Itoa(it.FilesCount)

	return it.Method.HexSumOfAny(slice)
}

func (it *FilesChecksumResult) IsEmpty() bool {
	return it.HasNoChecksum()
}

func (it *FilesChecksumResult) HasNoChecksum() bool {
	return it == nil ||
		it.HexFilesListChecksum == "" &&
			it.HexFilesContentsChecksum == ""
}

func (it *FilesChecksumResult) HasAnyChecksum() bool {
	return it != nil &&
		it.HexFilesListChecksum != "" ||
		it.HexFilesContentsChecksum != ""
}

func (it *FilesChecksumResult) HasFilesListChecksum() bool {
	return it != nil && it.HexFilesListChecksum != ""
}

func (it *FilesChecksumResult) HasContentsChecksum() bool {
	return it != nil && it.HexFilesContentsChecksum != ""
}

func (it *FilesChecksumResult) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *FilesChecksumResult) IsSuccess() bool {
	return it != nil && it.ErrorWrapper.IsSuccess()
}

func (it *FilesChecksumResult) IsFailed() bool {
	return it != nil && it.ErrorWrapper.IsFailed()
}

func (it *FilesChecksumResult) IsEqual(another *FilesChecksumResult) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.FilesCount != another.FilesCount {
		return false
	}

	if it.HexFilesListChecksum != another.HexFilesListChecksum {
		return false
	}

	if it.HexFilesContentsChecksum != another.HexFilesContentsChecksum {
		return false
	}

	if it.Method != another.Method {
		return false
	}

	return it.ErrorWrapper.IsEquals(another.ErrorWrapper)
}
