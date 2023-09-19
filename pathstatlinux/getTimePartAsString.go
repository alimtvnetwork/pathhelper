package pathstatlinux

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func getTimePartAsString(s string) string {
	leftRight := stringutil.SplitLeftRightType(s, constants.Colon)

	return strings.TrimSpace(leftRight.Right)
}
