package pathjoin

import "gitlab.com/auk-go/pathhelper/normalize"

func JoinFixedThree(
	first,
	second,
	third string,
) string {
	joined := JoinSimpleThree(
		first,
		second,
		third)

	return normalize.Path(joined)
}
