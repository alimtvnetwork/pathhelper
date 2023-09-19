package pathhelper

import (
	"gitlab.com/auk-go/core/constants"
)

func GetParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return GetBaseDir(location)
}
