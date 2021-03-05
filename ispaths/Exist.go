package ispaths

func Exist(paths ...string) *[]bool {
	if paths == nil {
		return &[]bool{}
	}

	return ExistPtr(&paths)
}
