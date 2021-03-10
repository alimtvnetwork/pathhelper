package fileinfo

import (
	"encoding/json"
	"os"
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
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

func (wrapper *Wrapper) JsonModel() *Wrapper {
	return wrapper
}

func (wrapper *Wrapper) JsonModelAny() interface{} {
	return wrapper.JsonModel()
}

//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) Json() *corejson.Result {
	jsonBytes, err := json.Marshal(wrapper)

	return corejson.NewPtr(jsonBytes, err)
}

//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return nil, nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &wrapper)

	if err != nil {
		return nil, nil
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
