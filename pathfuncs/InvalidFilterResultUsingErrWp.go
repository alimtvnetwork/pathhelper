package pathfuncs

import "gitlab.com/evatix-go/errorwrapper"

func InvalidFilterResultUsingErrWp(
	fullPath string,
	errWp *errorwrapper.Wrapper,
) *FilterResult {
	return &FilterResult{
		FullPath:     fullPath,
		ErrorWrapper: errWp,
	}
}
