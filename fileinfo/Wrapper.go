package fileinfo

import (
	"os"

	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"
)

type Wrapper struct {
	FileInfo    *os.FileInfo
	Error       errorwrapper.Wrapper
	RawPath     string
	IsDirectory bool
	IsFile      bool
	IsEmptyPath bool
	pathExists  issetter.Value
}

func (wrapper *Wrapper) HasError() bool {
	return wrapper.Error.HasError()
}

func (wrapper *Wrapper) IsPathExists() bool {
	if wrapper.pathExists.IsUninitialized() {
		isPathExists := !wrapper.HasError() && (wrapper.IsDirectory || wrapper.IsFile)
		wrapper.pathExists = issetter.GetBool(isPathExists)
	}

	return wrapper.pathExists.IsTrue()
}
