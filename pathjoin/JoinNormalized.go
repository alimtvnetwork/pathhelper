package pathjoin

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinNormalized normalized applied auto
func JoinNormalized(
	first, second string,
) string {
	if second == "" {
		return normalize.Path(first)
	}

	finalPath := first +
		osconsts.PathSeparator +
		second

	return normalize.Path(finalPath)
}
