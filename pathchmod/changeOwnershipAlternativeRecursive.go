package pathchmod

import (
	"errors"
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func changeOwnershipWindowsRecursive(path, userName, groupName string) *errorwrapper.Wrapper {
	userInfo := pathsysinfo.GetUserInfo(userName)

	if userInfo.ErrorWrapper.HasError() {
		return userInfo.ErrorWrapper
	}

	groupInfo := pathsysinfo.GetGroupInfo(groupName)

	if groupInfo.ErrorWrapper.HasError() {
		return groupInfo.ErrorWrapper
	}

	uid, gid := userInfo.Id, groupInfo.Id

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
		return errnew.NewPtr(errtype.ChmodApplyFailed, err)
	}

	return errnew.EmptyPtr
}
