package pathjoin

import "gitlab.com/evatix-go/core/osconsts"

// Doesn't apply normalize
func JoinSimple3(path1, path2, path3 string) string {
	return path1 +
		osconsts.PathSeparator +
		path2 +
		osconsts.PathSeparator +
		path3
}
