package enums

type PerformingAction int8

const (
	CreateAction PerformingAction = iota
	ReadAction
	DeleteAction
	NothingToPeform
)
