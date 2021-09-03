package normalize

import "gitlab.com/evatix-go/core/osconsts"

// getLongPathFixed
// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix ) or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
func getLongPathFixed(givenAbsolutePath string, isForce bool) string {
	return getLongPathFixedUsingSeparator(
		osconsts.PathSeparator,
		givenAbsolutePath,
		isForce)
}
