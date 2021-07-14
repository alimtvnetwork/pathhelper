package pathjoin

import "gitlab.com/evatix-go/core/osconsts"

// JoinSimple3ConditionalNonEmpty Doesn't apply normalize or doesn't join if empty
func JoinSimple3ConditionalNonEmpty(
	path1,
	path2,
	path3 string,
) string {
	if path3 == "" && path2 == "" {
		return path1
	}

	if path3 == "" {
		return JoinSimpleConditionalNonEmpty(
			path1,
			path2)
	}

	if path2 == "" {
		return JoinSimpleConditionalNonEmpty(
			path1,
			path3)
	}

	return path1 +
		osconsts.PathSeparator +
		path2 +
		osconsts.PathSeparator +
		path3
}
