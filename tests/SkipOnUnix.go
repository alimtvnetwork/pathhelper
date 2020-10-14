package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

// Skip on Unix
func SkipOnUnix(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip(unixIgnoreMessage)
	}
}
