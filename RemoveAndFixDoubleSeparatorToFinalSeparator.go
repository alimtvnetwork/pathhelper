package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
)

var removeAndFixDoubleSeparatorToFinalSeparatorMap = map[string]string{
	constants.ForwardSlash:       constants.BackSlash,
	constants.DoubleForwardSlash: constants.BackSlash,
	constants.DoubleBackSlash:    constants.BackSlash,
}

// Replace both double slashes to single slash (// -> /, \\ -> \) and finally all slashes to finalSeparator
func RemoveAndFixDoubleSeparatorToFinalSeparator(finalSeparator, path string) string {
	pathUsingBackSlash := GetCompiledPath(path, &removeAndFixDoubleSeparatorToFinalSeparatorMap)
	doubleSeparatorPath := ChangeSeparator(pathUsingBackSlash, constants.TripleBackSlash, constants.BackSlash)
	singleBackSlashesPath := ChangeDoubleBackSlash(doubleSeparatorPath, constants.BackSlash)
	finalPath := ChangeSeparator(singleBackSlashesPath, constants.BackSlash, finalSeparator)

	return finalPath
}
