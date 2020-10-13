package pathhelper

func PathFromEnvVariable(stringToCheck string) string {
	keyNameArray := GetVariables(stringToCheck)

	replacementMap := expandEnvironmentVariable(keyNameArray)

	return GetCompiledPath(stringToCheck, &replacementMap)
}
