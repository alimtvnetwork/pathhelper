package expandpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/regexnew"
)

func GetDollarOrPercentSymbolIdentifierEnvInfos(
	stringToCheck string,
) []EnvKeyInfo {
	envVariableRawKeys := regexnew.
		DollarIdentifierRegex.
		FindAllString(
			stringToCheck,
			constants.MinusOne)

	var envInfos []EnvKeyInfo

	if len(envVariableRawKeys) > 0 {
		envInfos = GetEnvInfosKeyNames(
			envVariableRawKeys)
	}

	envVariableRawKeys2 := regexnew.
		PercentIdentifierRegex.
		FindAllString(
			stringToCheck, constants.MinusOne)

	if len(envVariableRawKeys2) > 0 {
		envInfos2 := GetEnvInfosKeyNames(envVariableRawKeys2)

		for _, envInfo := range envInfos2 {
			envInfos = append(
				envInfos,
				envInfo)
		}
	}

	return envInfos
}
