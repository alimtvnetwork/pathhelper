package pathhelper

import (
	"os"
)

func expandEnvironmentVariable(variableForExapanding []string) map[string]string {
	var expandedPath = map[string]string{}

	for _, keyName := range variableForExapanding {
		_, exists := os.LookupEnv(keyName)

		if !exists {
			expandedPath["$"+keyName] = "$" + keyName
		}

		if exists {
			expandedPath["$"+keyName] = os.Getenv(keyName)
		}
	}

	return expandedPath
}
