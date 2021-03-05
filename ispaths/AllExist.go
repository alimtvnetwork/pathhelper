package ispaths

func AllExist(paths ...string) bool {
	if paths == nil {
		return false
	}

	allExistResults := ExistPtr(&paths)

	for _, isExist := range *allExistResults {
		if !isExist {
			return false
		}
	}

	return true
}
