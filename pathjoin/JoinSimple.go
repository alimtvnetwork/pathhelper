package pathjoin

import "gitlab.com/evatix-go/core/osconsts"

// JoinSimple doesn't apply normalize
func JoinSimple(path1, path2 string) string {
	return path1 + osconsts.PathSeparator + path2
}
