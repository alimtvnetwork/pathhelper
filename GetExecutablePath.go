package pathhelper

import "os"

func GetExecutablePath() string {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return exe
}
