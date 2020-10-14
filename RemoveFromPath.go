package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// Given removingList array items will be replaced with "" empty string.
// If pathTemplate is given as empty string or nil or whitespace then returns as is.
func RemoveFromPath(pathTemplate string, removingList *[]string, isNormalizePath bool) string {
	if pathhelpercore.IsEmptyPath(pathTemplate) {
		return pathTemplate
	}

	if isNormalizePath {
		pathTemplate = NormalizePath(pathTemplate)
	}

	for _, value := range *removingList {
		pathTemplate = strings.Replace(
			pathTemplate,
			value,
			constants.EmptyString,
			constants.MinusOne)
	}

	return pathTemplate
}
