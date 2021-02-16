package pathext

import (
	"path/filepath"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type Wrapper struct {
	fullPath                string
	dotExtension, extension *string
}

func New(path string) Wrapper {
	return Wrapper{
		fullPath: path,
	}
}

func NewPtr(path string) *Wrapper {
	return &Wrapper{
		fullPath: path,
	}
}

func (receiver *Wrapper) IsPathEquals(path string, ignoreCase bool) bool {
	if ignoreCase {
		return strings.EqualFold(receiver.fullPath, path)
	}

	return receiver.fullPath == path
}

func (receiver *Wrapper) IsPathContains(path string) bool {
	return strings.Contains(receiver.fullPath, path)
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (receiver *Wrapper) DotExtension() *string {
	if receiver.dotExtension == nil {
		dotExt := filepath.Ext(receiver.fullPath)
		receiver.dotExtension = &dotExt
	}

	return receiver.dotExtension
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (receiver *Wrapper) Extension() *string {
	if receiver.extension != nil {
		return receiver.extension
	}

	dotExt := *receiver.DotExtension()

	if len(dotExt) > 0 && dotExt[0] == constants.Dot[0] {
		ext := dotExt[1:]
		receiver.extension = &ext
	} else {
		receiver.extension = &dotExt
	}

	return receiver.extension
}

func (receiver *Wrapper) IsExtension(extension string) bool {
	return *receiver.Extension() == extension
}

func (receiver *Wrapper) IsDotExtension(dotExtension string) bool {
	return *receiver.DotExtension() == dotExtension
}

func (receiver *Wrapper) IsAnyOfExtension(extensions ...string) bool {
	if extensions == nil {
		return false
	}

	return receiver.IsAnyOfExtensionPtr(&extensions)
}

func (receiver *Wrapper) IsAnyOfExtensionPtr(
	extensions *[]string,
) bool {
	if extensions == nil || len(*extensions) == 0 {
		return false
	}

	currentExt := *receiver.Extension()
	for _, ext := range *extensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}

func (receiver *Wrapper) IsAnyOfDotExtension(
	dotExtensions ...string,
) bool {
	if dotExtensions == nil {
		return false
	}

	return receiver.IsAnyOfDotExtensionPtr(&dotExtensions)
}

func (receiver *Wrapper) IsAnyOfDotExtensionPtr(
	dotExtensions *[]string,
) bool {
	if dotExtensions == nil || len(*dotExtensions) == 0 {
		return false
	}

	currentExt := *receiver.DotExtension()
	for _, ext := range *dotExtensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}
