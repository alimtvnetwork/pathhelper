package pathhelper

// Represents the directory where the application is running from.
func GetExecutableDirectory() string {
	exePath := GetExecutablePath()
	exeDir, _ := Split(exePath)

	return exeDir
}
