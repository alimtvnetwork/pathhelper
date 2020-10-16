package pathhelper

func GetRoot() string {
	if IsWindows() {
		return GetWindowsRoot()
	}

	return GetUnixRoot()
}
