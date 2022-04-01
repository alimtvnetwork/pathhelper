package pathjoin

import "gitlab.com/evatix-go/core/osconsts"

// JoinNormalized3If normalized applied auto
func JoinNormalized3If(
	isNormalize bool,
	path1, path2, path3 string,
) string {
	if isNormalize {
		return JoinNormalizedThree(path1, path2, path3)
	}

	finalPath := path1 +
		osconsts.PathSeparator +
		path2 +
		osconsts.PathSeparator +
		path3

	return finalPath
}
