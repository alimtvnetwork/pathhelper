package expandpath

import (
	"os"
)

// ExpandEnvironmentVariable function takes an array of environment variables (string) as input
// and outputs a map of expanded path of those variables if the paths exist.
func ExpandEnvironmentVariable(envInfos *[]EnvKeyInfo) *map[string]string {
	var expandedPathMap = make(
		map[string]string,
		len(*envInfos))

	for _, envInfo := range *envInfos {
		name := envInfo.SimplifiedName
		_, isExist := os.LookupEnv(name)

		if isExist {
			expandedPathMap[envInfo.GivenAs] = os.Getenv(name)
		}
	}

	return &expandedPathMap
}
