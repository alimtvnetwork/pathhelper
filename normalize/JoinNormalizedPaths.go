package normalize

import (
	"strings"

	"gitlab.com/auk-go/core/osconsts"
)

func JoinNormalizedPaths(
	baseLocation string,
	locations ...string,
) string {
	if len(locations) == 0 {
		return Path(baseLocation)
	}

	combinedPath := baseLocation +
		osconsts.PathSeparator +
		strings.Join(locations,
			osconsts.PathSeparator)

	return Path(combinedPath)
}
