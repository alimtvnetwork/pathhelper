package pathfixer

import (
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func FixIf(isNormalize, isExpand bool, path string) string {
	if path == "" {
		return path
	}

	expand := expandpath.ExpandVariablesIf(
		isExpand, path)

	return normalize.PathUsingSingleIf(
		isNormalize,
		expand)
}
