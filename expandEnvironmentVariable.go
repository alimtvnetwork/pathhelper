package pathhelper

import (
	"os"
)

func expandEnvironmentVariable(variableToExpand []string) map[string]string {
	var expandedPath = map[string]string{}

	for _, keyName := range variableToExpand {
		_, exists := os.LookupEnv(keyName)

		if !exists {
			expandedPath["$"+keyName] = ""
		}

		if exists {
			expandedPath["$"+keyName] = os.Getenv(keyName)
		}
	}

	return expandedPath
}
