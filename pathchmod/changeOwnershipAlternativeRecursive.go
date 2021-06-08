package pathchmod

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func changeOwnershipWindowsRecursive(path, userName, groupName string) *errorwrapper.Wrapper {
	userObj, errLookup := user.Lookup(userName)
	if errLookup != nil {
		return errnew.MessagesPtr(
			errtype.SearchFailed,
			errLookup.Error(),
			" user name:",
			userName)
	}

	uid, errUidConvert := strconv.Atoi(userObj.Uid)
	if errUidConvert != nil {
		return errnew.NewPtr(errtype.ConversionValueToInteger, errUidConvert)
	}

	groupObj, errLookupGroup := user.LookupGroup(groupName)
	if errLookupGroup != nil {
		return errnew.MessagesPtr(
			errtype.SearchFailed,
			errLookupGroup.Error(),
			" groupName:",
			groupName)
	}

	gid, errGidConvert := strconv.Atoi(groupObj.Gid)
	if errGidConvert != nil {
		return errnew.MessagesPtr(
			errtype.SearchFailed,
			errGidConvert.Error(),
			" groupObj.Gid:",
			groupObj.Gid)
	}

	// https://github.com/gutengo/fil/blob/6109b2e0b5cfdefdef3a254cc1a3eaa35bc89284/file.go#L27-L34
	err := filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}

		if err == filepath.SkipDir {
			return nil
		}

		if err != nil {
			compiledErrMsg := err.Error() +
				", \nfailed path:" +
				path +
				", \nfailed for chown (name, uid, gid): " +
				name +
				constants.CommaSpace +
				converters.AnyToString(uid) +
				constants.CommaSpace +
				converters.AnyToString(gid)

			return errors.New(compiledErrMsg)
		}

		return nil
	})

	if err != nil {
		return errnew.NewPtr(errtype.FileOrDirectoryRelatedExecution, err)
	}

	return errnew.EmptyPtr
}
