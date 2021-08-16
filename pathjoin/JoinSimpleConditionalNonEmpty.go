package pathjoin

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

// JoinSimpleConditionalNonEmpty Don't apply normalize or not join if empty
func JoinSimpleConditionalNonEmpty(
	path1,
	path2 string,
) string {
	if path1 == constants.EmptyString && path2 == constants.EmptyString {
		return constants.EmptyString
	}

	if path2 == "" {
		return path1
	}

	if path1 == "" {
		return path2
	}

	return path1 +
		osconsts.PathSeparator +
		path2
}
