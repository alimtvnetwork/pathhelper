package pathhelpercore

func IsEmptyArray(paths []string) bool {
	return &paths == nil || paths == nil || len(paths) == 0
}

func IsEmptyArrayPtr(paths []*string) bool {
	return paths == nil || &paths == nil || len(paths) == 0
}
