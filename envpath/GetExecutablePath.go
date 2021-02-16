package envpath

import "os"

// Represents the exact path to the executable file
func GetExecutablePath() string {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return exe
}
