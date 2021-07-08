package expandpath

// EnvironmentVarExpand function takes a string input and replaces any word that starts with "$" in the input
// with its expanded path (if exists) and returns the new string.
func EnvironmentVarExpand(pathContainsEnvVariablesStartingDollar string) string {
	keyNameArray, _ := getDollarOrPercentSymbolIdentifierEnvInfos(
		pathContainsEnvVariablesStartingDollar)

	if len(keyNameArray) == 0 {
		return pathContainsEnvVariablesStartingDollar
	}

	replacementMap := expandEnvironmentVariable(&keyNameArray)

	return GetCompiledPath(
		pathContainsEnvVariablesStartingDollar,
		replacementMap)
}
