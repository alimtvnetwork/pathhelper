package pathhelper

import (
	"gitlab.com/auk-go/core/constants"
)

func FileNameJoiner(
	fileName,
	extensionWithOrWithoutDot string,
) string {
	if extensionWithOrWithoutDot == "" {
		return fileName
	}

	if extensionWithOrWithoutDot[constants.Zero] == constants.DotChar {
		return fileName + extensionWithOrWithoutDot
	}

	return fileName + constants.Dot + extensionWithOrWithoutDot
}
