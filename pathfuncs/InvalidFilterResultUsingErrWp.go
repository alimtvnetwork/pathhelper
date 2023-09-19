package pathfuncs

import "gitlab.com/auk-go/errorwrapper"

func InvalidFilterResultUsingErrWp(
	fullPath string,
	errWrap *errorwrapper.Wrapper,
) *FilterResult {
	return &FilterResult{
		FullPath:     fullPath,
		ErrorWrapper: errWrap,
	}
}
