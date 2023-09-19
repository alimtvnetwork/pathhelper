package fsinternal

import "gitlab.com/auk-go/pathhelper/internal/splitinternal"

func GetFileName(location string) string {
	return splitinternal.GetName(location)
}
