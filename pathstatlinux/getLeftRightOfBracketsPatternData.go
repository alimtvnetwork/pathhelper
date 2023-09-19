package pathstatlinux

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/pathhelper/internal/strremove"
)

// parentThesisWrappedSlashLeftRight : `(0755/drwxr-xr-x)  Uid` or `(left/right) whatever` data to left : 0755, right : drwxr-xr-x
func getLeftRightOfBracketsPatternData(parentThesisWrappedSlashLeftRight string) corestr.LeftRight {
	findSubStringsByRegex := bracketsMatcherWithContents.FindStringSubmatch(
		parentThesisWrappedSlashLeftRight)

	parentThesisPatternedData := stringslice.FirstOrDefault(findSubStringsByRegex)
	splitItems := strremove.SimpleManySplitsBy(
		parentThesisPatternedData,
		constants.ForwardSlash,
		constants.ParenthesisStart,
		constants.ParenthesisEnd)

	length := len(splitItems)

	if length == 2 {
		return corestr.LeftRight{
			Left:    strings.TrimSpace(splitItems[coreindexes.First]),
			Right:   strings.TrimSpace(splitItems[coreindexes.Second]),
			IsValid: true,
			Message: "",
		}
	}

	first, last := stringslice.FirstLastDefault(splitItems)

	return corestr.LeftRight{
		Left:    strings.TrimSpace(first),
		Right:   strings.TrimSpace(last),
		IsValid: false,
		Message: errcore.Expecting("Expected length", 2, length),
	}
}
