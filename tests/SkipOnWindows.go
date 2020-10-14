package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

const (
	unixIgnoreMessage = "Windows tests ignored in Unix."
	windowsIgnoreMessage = "Unix tests ignored in Windows."
)

// Skip tests on Windows
func SkipOnWindows(t *testing.T)  {
	if pathhelper.IsWindows() {
		t.Skip(windowsIgnoreMessage)
	}
}

