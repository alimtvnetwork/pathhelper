package pathhelper

func PathFromEnvVariable(stringToCheck string) string {
	keyNameArray := getVariables(stringToCheck)

	replacementMap := expandEnvironmentVariable(keyNameArray)

	return GetCompiledPath(stringToCheck, &replacementMap)
}
