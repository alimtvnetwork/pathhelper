package fileinfo

import (
	"encoding/json"
	"os"
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Wrapper struct {
	FileInfo     *os.FileInfo
	ErrorWrapper *errorwrapper.Wrapper
	RawPath      string
	IsDirectory  bool
	IsFile       bool
	IsEmptyPath  bool
	pathExists   issetter.Value
	Separator    string
	baseDir      *string
	parent       *Wrapper
}

func (wrapper *Wrapper) HasError() bool {
	return wrapper.ErrorWrapper.HasError()
}

func (wrapper *Wrapper) IsPathExists() bool {
	if wrapper.pathExists.IsUninitialized() {
		isPathExists := !wrapper.HasError() && (wrapper.IsDirectory || wrapper.IsFile)
		wrapper.pathExists = issetter.GetBool(isPathExists)
	}

	return wrapper.pathExists.IsTrue()
}

func (wrapper *Wrapper) BaseDir() string {
	if wrapper.baseDir != nil {
		return *wrapper.baseDir
	}

	baseDir := splitinternal.GetBaseDir(wrapper.RawPath)
	wrapper.baseDir = &baseDir

	return baseDir
}

func (wrapper *Wrapper) Parent() *Wrapper {
	if wrapper.parent != nil {
		return wrapper.parent
	}

	wrapper.parent = New(
		wrapper.BaseDir(),
		wrapper.Separator)

	return wrapper.parent
}

func (wrapper *Wrapper) GetBothExtensions() (dotExt, ext string) {
	return splitinternal.GetBothExtension(wrapper.RawPath)
}

func (wrapper *Wrapper) FileName() (filename string) {
	return (*wrapper.FileInfo).Name()
}

func (wrapper *Wrapper) FileNameWithoutExt() (filename string) {
	return splitinternal.GetFileNameWithoutExt(wrapper.RawPath)
}

func (wrapper *Wrapper) Size() int64 {
	return (*wrapper.FileInfo).Size()
}

func (wrapper *Wrapper) ModifyTime() time.Time {
	return (*wrapper.FileInfo).ModTime()
}

func (wrapper *Wrapper) AllSplits() *[]string {
	return splitinternal.GetAllSplitsWithSep(
		wrapper.RawPath,
		wrapper.Separator)
}

func (wrapper *Wrapper) String() string {
	return wrapper.RawPath
}

func (wrapper *Wrapper) ToString(sep string) string {
	return normalize.PathUsingSeparatorUsingSingleIf(
		true,
		sep,
		wrapper.RawPath,
	)
}

func (wrapper *Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(*wrapper.JsonModel())
}

func (wrapper *Wrapper) UnmarshalJSON(data []byte) error {
	var dataModel WrapperDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		wrapper.RawPath = dataModel.RawPath
		wrapper.IsDirectory = dataModel.IsDirectory
		wrapper.IsFile = dataModel.IsFile
		wrapper.IsEmptyPath = dataModel.IsEmptyPath
		wrapper.Separator = dataModel.Separator

		fileInfo, err2 := os.Stat(dataModel.RawPath)
		wrapper.FileInfo = &fileInfo
		wrapper.ErrorWrapper = errnew.ErrPtr(err2)
	}

	return err
}

func (wrapper *Wrapper) JsonModel() *WrapperDataModel {
	return &WrapperDataModel{
		RawPath:     wrapper.RawPath,
		IsDirectory: wrapper.IsDirectory,
		IsFile:      wrapper.IsFile,
		IsEmptyPath: wrapper.IsEmptyPath,
		Separator:   wrapper.Separator,
	}
}

func (wrapper *Wrapper) JsonModelAny() interface{} {
	return wrapper.JsonModel()
}

func (wrapper *Wrapper) Json() *corejson.Result {
	return corejson.NewFromAny(wrapper)
}

//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return nil, defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(*jsonResult.Bytes, &wrapper)

	if err != nil {
		return nil, err
	}

	return wrapper, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrapper {
	newUsingJson, err :=
		wrapper.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (wrapper *Wrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := wrapper.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (wrapper *Wrapper) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = wrapper

	return &jsoner
}

func (wrapper *Wrapper) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonInjector corejson.ParseSelfInjector = wrapper

	return &jsonInjector
}
