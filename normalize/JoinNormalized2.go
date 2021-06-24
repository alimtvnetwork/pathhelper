package normalize

import "gitlab.com/evatix-go/core/osconsts"

func JoinNormalized2(baseLocation string, nextLocation string) string {
	if nextLocation == "" {
		return Path(baseLocation)
	}

	combinedPath := baseLocation +
		osconsts.PathSeparator +
		nextLocation

	return Path(combinedPath)
}
