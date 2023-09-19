package fs

import (
	"path/filepath"

	"gitlab.com/auk-go/core/constants"
)

func ParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return filepath.Dir(location)
}
