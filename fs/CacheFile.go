package fs

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/isany"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type CacheFile struct {
	ChmodWrapper              pathchmod.Wrapper
	AbsParentDir, AbsFilePath string
	IsAcquireLock             bool
	IsCollectReadError        bool
	IsCollectWriteError       bool
	IsWriteEmptyOnNull        bool
	IsRemoveFileBeforeWrite   bool
	isCompiled                bool
	cacheData                 interface{}
	lastJsonResult            *corejson.Result
	compileErr                *errorwrapper.Wrapper
	OnInvalidGenerateFunc     func(toPtr interface{}) *errorwrapper.Wrapper
}

// GetOnce
//
//  on compile success reflect set to ToPtr
//  on compile error returns previous compiling error
func (it *CacheFile) GetOnce(toPtr interface{}) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isCompiled && it.compileErr.IsEmpty() {
		return errnew.
			Reflect.
			SetFromTo(it.cacheData, toPtr)
	}

	if it.isCompiled && it.compileErr.HasError() {
		return it.compileErr
	}

	if it.IsAcquireLock {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	// not generated yet
	// check in file first
	var readFromFileErrWrap *errorwrapper.Wrapper
	isFileExist := it.IsFileExist()
	if it.IsFileExist() {
		// read from it
		readFromFileErrWrap = it.readFromFileInternal(
			toPtr)
	}

	if isFileExist && readFromFileErrWrap.IsEmpty() {
		// all success
		return nil
	}

	if it.isCollectReadErrorWrap(isFileExist, readFromFileErrWrap) {
		// error needs to return
		return readFromFileErrWrap
	}

	// may not exist in file or read error
	// generate
	generateErrWrap := it.OnInvalidGenerateFunc(toPtr)
	it.isCompiled = true
	it.compileErr = generateErrWrap
	it.cacheData = toPtr

	if generateErrWrap.HasError() {
		return generateErrWrap
	}

	// clear, save
	savingErrorWrap := it.Save(toPtr)
	if it.IsCollectWriteError && savingErrorWrap.HasError() {
		return savingErrorWrap
	}

	return nil
}

func (it *CacheFile) isCollectReadErrorWrap(
	isFileExist bool,
	readFromFileErrWrap *errorwrapper.Wrapper,
) bool {
	return it.IsCollectReadError &&
		isFileExist &&
		readFromFileErrWrap.HasError()
}

func (it *CacheFile) ReadFromFile(
	toPtr interface{},
) *errorwrapper.Wrapper {
	if it.IsAcquireLock {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if it.IsNotFileExist() {
		return errnew.Path.TypeMsg(
			errtype.FileNotExist,
			"file not exist to read and unmarshal",
			it.AbsFilePath)
	}

	// file exist
	return JsonReadUnmarshal(it.AbsFilePath, toPtr)
}

func (it *CacheFile) readFromFileInternal(
	toPtr interface{},
) *errorwrapper.Wrapper {
	// file exist
	return JsonReadUnmarshal(it.AbsFilePath, toPtr)
}

func (it *CacheFile) IsFileExist() bool {
	return chmodhelper.IsPathExists(it.AbsFilePath)
}

func (it *CacheFile) IsNotFileExist() bool {
	return !chmodhelper.IsPathExists(it.AbsFilePath)
}

func (it *CacheFile) IsCompiled() bool {
	return it != nil && it.isCompiled
}

func (it *CacheFile) IsApplicable() bool {
	return it != nil &&
		it.isCompiled &&
		it.compileErr.IsEmpty()
}

func (it *CacheFile) IsErrorOnNull() bool {
	return it != nil &&
		!it.IsWriteEmptyOnNull
}

func (it *CacheFile) Save(
	toPtr interface{},
) *errorwrapper.Wrapper {
	isNull := isany.Null(toPtr)

	if it.IsErrorOnNull() && isNull {
		return errnew.Null.WithMessage(
			"cannot save cache data on nil given",
			toPtr)
	}

	// can be null
	if isNull {
		// on null write empty
		return WriteAllParams(
			true,
			it.ChmodWrapper.IsSkipOnInvalid,
			it.ChmodWrapper.IsKeepExistingChmod,
			it.ChmodWrapper.DirChmod,
			it.ChmodWrapper.FileChmod,
			it.AbsFilePath,
			[]byte(""))
	}

	toJsonResult := corejson.NewPtr(toPtr)
	if toJsonResult.HasError() {
		return errnew.Error.Default(
			errtype.Serialize,
			toJsonResult.MeaningfulError())
	}

	if it.IsRemoveFileBeforeWrite {
		// swallow is fine here.
		os.RemoveAll(it.AbsFilePath)
		// next error will be caught in write.
	}

	return WriteAllParams(
		true,
		it.ChmodWrapper.IsSkipOnInvalid,
		it.ChmodWrapper.IsKeepExistingChmod,
		it.ChmodWrapper.DirChmod,
		it.ChmodWrapper.FileChmod,
		it.AbsFilePath,
		toJsonResult.Bytes)
}
