package nginxlinuxpath

import "gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"

func fixPath(root, next string) string {
	return normalizeinternal.FixIf(
		true,
		root,
		next)
}

func fixPathIf(isFix bool, root, next string) string {
	return normalizeinternal.FixIf(
		isFix,
		root,
		next)
}
