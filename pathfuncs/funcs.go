package pathfuncs

type (
	FilterResult struct {
		FullPath        string
		IsKeep, IsBreak bool
	}

	Filter func(arg *FilterArg) *FilterResult
)
