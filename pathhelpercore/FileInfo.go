package pathhelpercore

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/enum/strtype"
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"
)

type FileInfo struct {
	StateType           strtype.Variant
	RootInfo            *coretaskinfo.Info
	FilePath            string
	FileRawBytesResults *errbyte.Results
}

func (it *FileInfo) Json() corejson.Result {
	return corejson.New(it)
}

func (it *FileInfo) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *FileInfo) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it FileInfo) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
