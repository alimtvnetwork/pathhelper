package pathjoin

import "gitlab.com/auk-go/pathhelper/normalize"

func DbPaths(locations ...string) string {
	simpleJoin := normalize.SimpleJoinPaths(
		locations...)

	return normalize.DbPath(simpleJoin)
}
