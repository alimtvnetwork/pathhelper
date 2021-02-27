package pathhelper

import "gitlab.com/evatix-go/core"

func GetPathAsUrisOf(
	isNormalizePath bool,
	paths ...string,
) *[]string {
	if paths == nil {
		return core.EmptyStringsPtr()
	}

	return GetPathAsUrisOfPtr(
		isNormalizePath,
		&paths)
}
