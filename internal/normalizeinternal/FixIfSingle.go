package normalizeinternal

import (
	"path"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
)

func FixIfSingle(isFix bool, givenPath string) string {
	if len(givenPath) == 0 {
		return constants.EmptyString
	}

	if !isFix {
		return givenPath
	}

	givenPath2 := strings.ReplaceAll(
		givenPath,
		constants.ForwardSlash,
		osconsts.PathSeparator)

	return path.Clean(givenPath2)
}
