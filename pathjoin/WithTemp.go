package pathjoin

import (
	"path"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func WithTemp(
	locations ...string,
) string {
	joinedPath := path.Join(locations...)

	return normalize.Path(joinedPath)
}
