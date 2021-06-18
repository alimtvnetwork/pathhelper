package cmdprefix

import (
	"fmt"

	"gitlab.com/evatix-go/core/chmodhelper"
)

// recursiveFormat (recursive format, non recursive format)
func recursiveFormat(
	isRecursive bool,
	formatRecursive string,
	nonRecursiveFormat string,
	args ...interface{},
) string {
	if isRecursive {
		return fmt.Sprintf(formatRecursive, args...)
	}

	return fmt.Sprintf(nonRecursiveFormat, args...)
}

// Chmod Format :
//  - Recursive : chmod -R 777 /dir
//  - Non Recursive : chmod 777 /dir
func Chmod(isRecursive bool, wrapper *chmodhelper.RwxWrapper) string {
	octalModeValueString := wrapper.ToFileModeString()

	return recursiveFormat(
		isRecursive,
		chmodRecursiveFormat,
		chmodNonRecursiveFormat,
		octalModeValueString)
}

// ChownUser Format: chown -R $user:$group /dir
//  - Recursive : chown -R $user:$group /dir
//  - Non Recursive : chown $user:$group /dir
func ChownUser(isRecursive bool, userName, groupName string) string {
	return recursiveFormat(
		isRecursive,
		chownRecursiveFormat,
		chownNonRecursiveFormat,
		userName,
		groupName)
}

// ChangeGroup Format: chgrp -R $group /dir
//  - Recursive (changeGroupRecursiveFormat) : chgrp -R $group
//  - Non Recursive (changeGroupNonRecursiveFormat) : chgrp $group
func ChangeGroup(isRecursive bool, groupName string) string {
	return recursiveFormat(
		isRecursive,
		changeGroupRecursiveFormat,
		changeGroupNonRecursiveFormat,
		groupName)
}
