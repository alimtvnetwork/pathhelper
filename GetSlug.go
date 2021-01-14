package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var forbiddenArray = []string{"!", "`", "@", "#", "%", "$", "^", "&", "*", "(", ")", "{", "}", "[", "]", " "}

// GetSlug from given path, usages @forbiddenArray to replace with @separatorOfChoice
func GetSlug(path, separatorOfChoice string) string {
	if pathhelpercore.IsEmptyPath(path) {
		return path
	}

	for _, forbidden := range forbiddenArray {
		path = strings.ReplaceAll(path, forbidden, separatorOfChoice)
	}

	return RemoveDoubleUriSeparator(path, separatorOfChoice)
}
