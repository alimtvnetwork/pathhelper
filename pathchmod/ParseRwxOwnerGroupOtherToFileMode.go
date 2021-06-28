package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ParseRwxOwnerGroupOtherToFileMode(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (os.FileMode, *errorwrapper.Wrapper) {
	rwxWrapper, errWp := ParseRwxOwnerGroupOtherToRwxWrapper(rwxOwnerGroupOther)

	if errWp.HasError() {
		return 0, errWp
	}

	if rwxWrapper == nil {
		return 0, errnew.MessagesPtr(
			errtype.ChmodInvalid,
			"Cannot process wildcard rwx to convert to fixed file mode")
	}

	return rwxWrapper.ToFileMode(), errnew.EmptyPtr
}
