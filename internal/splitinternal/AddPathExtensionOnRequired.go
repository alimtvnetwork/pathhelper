package splitinternal

import (
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func AddPathExtensionOnRequired(
	location string,
	dotExtension string,
) (locationFinal string) {
	if stringutil.IsEndsWith(location, dotExtension, true) {
		return location
	}

	return location + dotExtension
}
