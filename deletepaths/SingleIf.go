package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SingleIf(
	isRemove bool,
	location string,
) *errorwrapper.Wrapper {
	if !isRemove {
		return errnew.EmptyPtr
	}

	err := os.Remove(location)

	return errnew.Path(errtype.DeletePathFailed, err, location)
}
