package pathjoin

import "gitlab.com/auk-go/pathhelper/normalize"

func DbPath(first, second string) string {
	simpleJoin := JoinSimple(first, second)

	return normalize.DbPath(simpleJoin)
}
