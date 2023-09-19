package filestate

import (
	"gitlab.com/auk-go/pathhelper/hashas"
)

func InvalidInfo(hashMethod hashas.Variant, filePath string) *Info {
	return &Info{
		FullPath:   filePath,
		IsFile:     false,
		IsInvalid:  true,
		HashMethod: hashMethod,
	}
}
