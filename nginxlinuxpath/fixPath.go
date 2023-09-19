package nginxlinuxpath

import "gitlab.com/auk-go/pathhelper/internal/normalizeinternal"

func fixPath(root, next string) string {
	return normalizeinternal.JoinFixIf(
		true,
		root,
		next)
}
