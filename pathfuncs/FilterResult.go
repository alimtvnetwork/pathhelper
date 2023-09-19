package pathfuncs

import (
	"gitlab.com/auk-go/errorwrapper"
)

type FilterResult struct {
	FullPath        string
	IsTake, IsBreak bool
	ErrorWrapper    *errorwrapper.Wrapper
}
