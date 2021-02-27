package expandpath

import (
	"gitlab.com/evatix-go/pathhelper"
)

// EnvironmentVarExpand function takes a string input and replaces any word that starts with "$" in the input
// with its expanded path (if exists) and returns the new string.
func EnvironmentVarExpand(pathContainsEnvVariablesStartingDollar string) string {
	keyNameArray := GetEnvironmentVariables(pathContainsEnvVariablesStartingDollar)
	replacementMap := expandEnvironmentVariable(keyNameArray)

	return pathhelper.GetCompiledPath(
		pathContainsEnvVariablesStartingDollar,
		replacementMap)
}
