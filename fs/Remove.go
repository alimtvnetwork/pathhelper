package fs

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

// Remove Reference : https://t.ly/xnAe
func Remove(location string) *errorwrapper.Wrapper {
	err := os.RemoveAll(location)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(errtype.RemoveFailed, err, location)
}
