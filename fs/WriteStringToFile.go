package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFile(
	filePath string, content string,
) *errorwrapper.Wrapper {
	return WriteFile(
		filePath, []byte(content))
}
