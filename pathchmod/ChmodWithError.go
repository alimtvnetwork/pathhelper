package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper/errinf"
)

type ChmodWithError struct {
	Chmod os.FileMode
	errinf.ErrWrapper
}
