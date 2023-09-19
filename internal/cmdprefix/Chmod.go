package cmdprefix

import (
	"gitlab.com/auk-go/core/chmodhelper"
)

// Chmod Format :
//   - Recursive : chmod -R 777 /dir
//   - Non-Recursive : chmod 777 /dir
func Chmod(isRecursive bool, wrapper *chmodhelper.RwxWrapper) string {
	octalModeValueString := wrapper.ToFileModeString()

	return recursiveFormat(
		isRecursive,
		chmodRecursiveFormat,
		chmodNonRecursiveFormat,
		octalModeValueString)
}
