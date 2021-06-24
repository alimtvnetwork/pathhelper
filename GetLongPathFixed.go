package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// GetLongPathFixed
//
// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix )
// or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
// if forced then if Windows then fix added anyway.
func GetLongPathFixed(isForce bool, givenAbsolutePath string) string {
	return normalize.GetLongPathFixedUsingSeparator(
		osconsts.PathSeparator,
		givenAbsolutePath,
		isForce)
}
