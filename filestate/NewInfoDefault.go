package filestate

import "gitlab.com/auk-go/errorwrapper"

func NewInfoDefault(
	filePath string,
) (*Info, *errorwrapper.Wrapper) {
	return NewInfo(
		DefaultHashMethod,
		true,
		filePath)
}
