package fileinfopath

type (
	FilterFunc        func(info *Instance) (isTake, isBreak bool)
	TakeAllFilterFunc func(info *Instance) (isTake bool)
	StringerFmtFunc   func(index int, info *Instance) string
	HasFilterFunc     func(index int, info *Instance) (isSuccess bool)
	HasKeyFilterFunc  func(key string, info *Instance) (isSuccess bool)
)
