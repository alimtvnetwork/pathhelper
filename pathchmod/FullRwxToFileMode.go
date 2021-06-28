package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func FullRwxToFileMode(rwxFull string) (os.FileMode, *errorwrapper.Wrapper) {
	rwxWrapper, errWp := FullRwxToRwxWrapper(rwxFull)

	if errWp.HasError() {
		return 0, errWp
	}

	return rwxWrapper.
		ToFileMode(), errnew.EmptyPtr
}
