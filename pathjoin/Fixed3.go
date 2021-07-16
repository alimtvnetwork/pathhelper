package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func Fixed3(
	path1,
	path2,
	path3 string,
) string {
	joined := JoinSimple3ConditionalNonEmpty(
		path1,
		path2,
		path3)

	return normalize.Path(joined)
}
