package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

// SafeRemove Reference : https://t.ly/xnAe
func SafeRemove(location string) *errorwrapper.Wrapper {
	if IsPathExists(location) {
		return Remove(location)
	}

	return errnew.EmptyPtr
}
