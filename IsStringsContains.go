package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func isStringsContains(array []string, findingItem string) bool {
	if pathhelpercore.IsEmptyArray(array) {
		return false
	}

	for _, arrayItem := range array {
		if strings.Compare(arrayItem, findingItem) == 0 {
			return true
		}
	}

	return false
}
