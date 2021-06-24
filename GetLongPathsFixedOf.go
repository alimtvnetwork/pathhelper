package pathhelper

import "gitlab.com/evatix-go/core"

// GetLongPathsFixedOf
//
// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix )
// or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
// if forced then if Windows then fix added anyway.
func GetLongPathsFixedOf(
	isForce bool,
	pathSeparator string,
	givenAbsolutePaths ...string,
) *[]string {
	if givenAbsolutePaths == nil {
		return core.EmptyStringsPtr()
	}

	return GetLongPathsFixedOfPtr(
		isForce,
		pathSeparator,
		&givenAbsolutePaths)
}
