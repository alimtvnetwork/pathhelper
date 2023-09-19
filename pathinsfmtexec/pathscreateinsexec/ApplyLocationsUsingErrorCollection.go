package pathscreateinsexec

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/createpath"
	"gitlab.com/auk-go/pathhelper/deletepaths"
	"gitlab.com/auk-go/pathhelper/pathchmod"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/namegroup"
)

func ApplyLocationsUsingErrorCollection(
	isLock,
	isDeleteAllBeforeCreate,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	applyNameGroup *pathinsfmt.BaseUserNameGroupName,
	applyRwx *chmodins.RwxOwnerGroupOther,
	locations []string,
) (
	isSuccess bool,
) {
	if len(locations) == 0 {
		return true
	}

	errCount := errorCollection.Length()

	if isDeleteAllBeforeCreate {
		errorCollection.AddWrapperPtr(
			deletepaths.All(locations...))
	}

	// paths create
	// create without chmod
	_, filesCreateErr := createpath.CreateMany(
		isLock,
		isIgnoreOnExist,
		locations)

	errorCollection.AddWrapperPtr(filesCreateErr)

	if applyRwx != nil {
		chmodErr := pathchmod.ApplyChmodRwxOwnerGroupOther(
			false,
			false,
			false,
			applyRwx,
			locations)

		errorCollection.AddWrapperPtr(chmodErr)
	}

	// apply groups
	if applyNameGroup != nil && osconsts.IsUnixGroup {
		errWrap := namegroup.Apply(
			true,
			false,
			&applyNameGroup.UserGroupName,
			locations...)

		errorCollection.AddWrapperPtr(errWrap)
	}

	return errCount == errorCollection.Length()
}
