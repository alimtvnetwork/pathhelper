package enums

type PerformingAction int8

//goland:noinspection ALL
const (
	CreateAction PerformingAction = iota
	ReadAction
	DeleteAction
	EmptyDirectoryResult
)
