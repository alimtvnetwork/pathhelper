package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinNormalizedThree normalized applied auto
func JoinNormalizedThree(first, second, third string) string {
	joined := normalize.SimpleJoinPath3(
		first,
		second,
		third)

	return normalize.Path(joined)
}
