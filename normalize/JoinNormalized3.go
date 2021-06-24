package normalize

import "gitlab.com/evatix-go/core/osconsts"

func JoinNormalized3(baseLocation string, next1, next2 string) string {
	if next1 == "" && next2 == "" {
		return Path(baseLocation)
	}

	if next1 != "" && next2 == "" {
		return JoinNormalized2(baseLocation, next1)
	}

	if next1 == "" && next2 != "" {
		return JoinNormalized2(baseLocation, next2)
	}

	combinedPath := baseLocation +
		osconsts.PathSeparator +
		next1 +
		osconsts.PathSeparator +
		next2

	return Path(combinedPath)
}
