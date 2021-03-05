package ispath

import "os"

func NotExists(path string) bool {
	_, err := os.Stat(path)

	return err != nil
}
