package pathscreateinsexec

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func SimpleCreate(
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

	file, errWp := createpath.CreateSingleUsingFileMode(
		mode,
		filePath,
	)

	if errWp.HasError() {
		return file, errWp
	}

	if userGroupName != nil {
		return file, pathchmod.ChangeOwnership(
			filePath,
			userGroupName.UserName,
			userGroupName.GroupName)
	}

	return file, errnew.EmptyPtr
}
