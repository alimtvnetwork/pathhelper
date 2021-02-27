package pathhelper

import "gitlab.com/evatix-go/core"

func GetSlugsOf(
	separatorOfChoice string,
	paths ...string,
) *[]string {
	if paths == nil {
		return core.EmptyStringsPtr()
	}

	return GetSlugsOfPtr(
		separatorOfChoice,
		&paths)
}
