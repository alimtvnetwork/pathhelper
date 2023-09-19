package pathchmod

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
)

type ChmodWithError struct {
	Chmod os.FileMode
	errorwrapper.ErrWrapper
}
