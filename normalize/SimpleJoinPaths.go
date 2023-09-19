package normalize

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/osconsts"
)

func SimpleJoinPaths(locations ...string) string {
	if len(locations) == 0 {
		return constants.EmptyString
	}

	return stringslice.NonEmptyJoin(
		locations,
		osconsts.PathSeparator)
}
