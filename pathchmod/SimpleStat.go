package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errinf"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

type SimpleStat struct {
	Location        string
	FileInfo        os.FileInfo
	HasFileInfo     bool
	InvalidFileInfo bool
	IsNotExist      bool
	IsExist         bool
	IsDir           bool
	IsFile          bool
	errinf.ErrWrapper
}

func (it *SimpleStat) ReadString() *errstr.Result {
	errWp := it.notFileError()
	if errWp.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: errWp,
		}
	}

	errBytes := fsinternal.ReadFile(it.Location)

	return &errstr.Result{
		Value:        errBytes.String(),
		ErrorWrapper: errBytes.ErrorWrapper,
	}
}

func (it *SimpleStat) notFileError() *errorwrapper.Wrapper {
	if !it.IsExist || it.IsDir {
		return errnew.PathMessages(
			errtype.File,
			it.Location,
			"Cannot read invalid path or a directory.")
	}

	return errnew.EmptyPtr
}

func (it *SimpleStat) ReadBytes() *errbyte.Results {
	errWp := it.notFileError()
	if errWp.HasError() {
		return &errbyte.Results{
			Values:       &[]byte{},
			ErrorWrapper: errWp,
		}
	}

	return fsinternal.ReadFile(it.Location)
}

func (it *SimpleStat) CheckSum(hashType hashas.Variant) *errbyte.Results {
	errWp := it.notFileError()
	if errWp.HasError() {
		return &errbyte.Results{
			Values:       &[]byte{},
			ErrorWrapper: errWp,
		}
	}

	allBytes := it.ReadBytes()

	if allBytes.HasError() {
		return &errbyte.Results{
			Values:       &[]byte{},
			ErrorWrapper: allBytes.ErrorWrapper,
		}
	}

	return hashType.SumOf(*allBytes.Values)
}

func (it *SimpleStat) CheckSumHexString(hashType hashas.Variant) *errstr.Result {
	errWp := it.notFileError()
	if errWp.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: errWp,
		}
	}

	allBytes := it.ReadBytes()

	if allBytes.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: allBytes.ErrorWrapper,
		}
	}

	return hashType.StringSumOf(*allBytes.Values)
}

func (it *SimpleStat) GetChmodWithError() *ChmodWithError {
	return ExistingChmodWithError(it.Location)
}

func (it *SimpleStat) GetRwxWithError() *RwxWrapperWithError {
	return ExistingRwxWrapperWithError(it.Location)
}
