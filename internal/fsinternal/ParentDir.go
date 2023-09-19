package fsinternal

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func ParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return splitinternal.GetBaseDir(location)
}
