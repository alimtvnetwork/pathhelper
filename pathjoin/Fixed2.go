package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func Fixed2(
	path1,
	path2 string,
) string {
	joined := JoinSimpleConditionalNonEmpty(
		path1,
		path2)

	return normalize.Path(joined)
}
