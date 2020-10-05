package pathhelper

import "strings"

var forbiddenArray = []string{"!", "`", "@", "#", "%", "$", "^", "&", "*", "(", ")", "{", "}", "[", "]", " "}

func GetSlug(path, separatorOfChoice string) string {
	if IsSlug(path) {
		return ""
	}

	for i, _ := range forbiddenArray {
		path = strings.ReplaceAll(path, forbiddenArray[i], separatorOfChoice)
	}

	return RemoveDoubleSeparator(path, separatorOfChoice)
}
