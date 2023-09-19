package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
)

type RwxWrapperWithError struct {
	RwxWrapper *chmodhelper.RwxWrapper
	errorwrapper.ErrWrapper
}
