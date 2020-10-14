package pathhelper

// PathFromEnvVariable function takes a string input and replaces any word that starts with "$" in the input
// with its expanded path (if exists) and returns the new string.
func PathFromEnvVariable(stringToCheck string) string {
	keyNameArray := GetVariables(stringToCheck)

	replacementMap := expandEnvironmentVariable(keyNameArray)

	return GetCompiledPath(stringToCheck, &replacementMap)
}
