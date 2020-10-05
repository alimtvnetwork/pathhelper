package pathhelper

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecutableDirectory() string {
	return filepath.Dir(strings.Join(os.Args, "|"))
}
