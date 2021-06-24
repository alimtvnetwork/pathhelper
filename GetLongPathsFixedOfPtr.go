package pathhelper

import "gitlab.com/evatix-go/pathhelper/normalize"

// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix )
// or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
// if forced then if Windows then fix added anyway.
func GetLongPathsFixedOfPtr(
	isForce bool,
	pathSeparator string,
	givenAbsolutePaths *[]string,
) *[]string {
	return GetAsyncProcessed(
		givenAbsolutePaths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return normalize.GetLongPathFixedUsingSeparator(
				pathSeparator,
				currentPath,
				isForce)
		})
}
