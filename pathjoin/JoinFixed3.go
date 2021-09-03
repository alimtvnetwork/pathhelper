package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func JoinFixed3(
	path1,
	path2,
	path3 string,
) string {
	joined := JoinSimple3(
		path1,
		path2,
		path3)

	return normalize.Path(joined)
}
