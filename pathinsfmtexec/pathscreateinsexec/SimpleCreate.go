package pathscreateinsexec

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/createpath"
	"gitlab.com/auk-go/pathhelper/deletepaths"
	"gitlab.com/auk-go/pathhelper/pathchmod"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func SimpleCreate(
	isLock,
	isRemoveBeforeCreate bool,
	filePath string,
	mode os.FileMode,
	userGroupName *pathinsfmt.UserGroupName,
) (*os.File, *errorwrapper.Wrapper) {
	removeErr := deletepaths.SingleOnExistIf(
		isRemoveBeforeCreate,
		filePath)

	if removeErr.IsFailed() {
		return nil, removeErr
	}

	file, errWrap := createpath.CreateSingleUsingFileMode(
		isLock,
		mode,
		filePath,
	)

	if errWrap.HasError() {
		return file, errWrap
	}

	if userGroupName != nil {
		return file, pathchmod.ChangeOwnership(
			filePath,
			userGroupName.UserName,
			userGroupName.GroupName)
	}

	return file, nil
}
