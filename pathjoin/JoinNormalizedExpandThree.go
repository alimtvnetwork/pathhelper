package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinNormalizedExpandThree normalized and expand applied auto
func JoinNormalizedExpandThree(
	first, second, third string,
) string {
	joined := JoinSimpleThree(first, second, third)
	expand := expandpath.ExpandVariables(joined)

	return normalize.Path(expand)
}
