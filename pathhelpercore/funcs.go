package pathhelpercore

import "gitlab.com/auk-go/errorwrapper"

type (
	InvokerFunc func(fileInfo *FileInfo) *errorwrapper.Wrapper
)
