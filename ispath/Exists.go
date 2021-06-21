package ispath

import "os"

func Exists(location string) bool {
	_, err := os.Stat(location)

	return !os.IsNotExist(err)
}
