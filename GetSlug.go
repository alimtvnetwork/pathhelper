package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var forbiddenArray = []string{"!", "`", "@", "#", "%", "$", "^", "&", "*", "(", ")", "{", "}", "[", "]", " "}

func GetSlug(path, separatorOfChoice string) string {
	if pathhelpercore.IsEmptyPath(path){
		return path
	}

	if IsSlug(path) {
		return path
	}

	for i, _ := range forbiddenArray {
		path = strings.ReplaceAll(path, forbiddenArray[i], separatorOfChoice)
	}

	return RemoveDoubleUriSeparator(path, separatorOfChoice)
}
