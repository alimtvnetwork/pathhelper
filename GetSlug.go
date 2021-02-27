package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

var slugForbiddenArray = []string{
	" ",
	"!",
	"`",
	"@",
	"#",
	"%",
	"$",
	"^",
	"&",
	"*",
	"(",
	")",
	"{",
	"}",
	"[",
	"]",
}

// GetSlug from given path, usages @slugForbiddenArray to replace with @separatorOfChoice
func GetSlug(path, separatorOfChoice string) string {
	if ispath.Empty(path) {
		return path
	}

	for _, forbidden := range slugForbiddenArray {
		path = strings.ReplaceAll(path, forbidden, separatorOfChoice)
	}

	return RemoveDoubleUriSeparator(path, separatorOfChoice)
}
