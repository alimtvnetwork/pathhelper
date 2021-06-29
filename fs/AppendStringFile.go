package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendStringFile(
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return AppendFile(
		filePath,
		[]byte(content))
}
