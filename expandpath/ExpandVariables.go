package expandpath

// ExpandVariables
//
// function takes a string input and replaces any word that starts with "$" in the input
// with its expanded path (if exists) and returns the new string.
//
// Acceptable Env paths:
// ${Java_home} $java_home %{java_home} %java_home all will be expand e
func ExpandVariables(pathContainsEnvVariables string) string {
	envInfos := GetDollarOrPercentSymbolIdentifierEnvInfos(
		pathContainsEnvVariables)

	if len(envInfos) == 0 {
		return pathContainsEnvVariables
	}

	replacementMap := ExpandEnvironmentVariable(&envInfos)

	return GetCompiledPath(
		pathContainsEnvVariables,
		*replacementMap)
}
