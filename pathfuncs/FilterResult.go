package pathfuncs

import "gitlab.com/evatix-go/errorwrapper"

type FilterResult struct {
	FullPath        string
	IsKeep, IsBreak bool
	ErrorWrapper    *errorwrapper.Wrapper
}
