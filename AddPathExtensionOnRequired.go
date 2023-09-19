package pathhelper

import (
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
)

func AddPathExtensionOnRequired(location, dotExt string) string {
	return splitinternal.AddPathExtensionOnRequired(location, dotExt)
}
