package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func DbPath(path1, path2 string) string {
	simpleJoin := JoinSimple(path1, path2)

	return normalize.DbPath(simpleJoin)
}
