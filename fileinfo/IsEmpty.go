package fileinfo

func IsEmpty(paths []*Wrapper) bool {
	return paths == nil || &paths == nil || len(paths) == 0
}
