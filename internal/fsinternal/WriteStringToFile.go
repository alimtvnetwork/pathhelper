package fsinternal

import "gitlab.com/auk-go/errorwrapper"

func WriteStringToFile(
	filePath string, content string,
) *errorwrapper.Wrapper {
	return WriteFileStringDefault(
		filePath,
		content,
	)
}
