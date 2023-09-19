package fs

import "gitlab.com/auk-go/errorwrapper"

type (
	CopierOrMoverFunc = func(source, destination string) *errorwrapper.Wrapper
)
