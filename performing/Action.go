package performing

type Action int8

//goland:noinspection ALL
const (
	CreateAction Action = iota
	ReadAction
	DeleteAction
	EmptyDirectoryResult
)
