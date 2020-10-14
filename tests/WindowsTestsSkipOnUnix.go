package tests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper"
)

func WindowsTestsSkipOnUnix(t *testing.T) {
	if !pathhelper.IsWindows() {
		t.Skip(unixIgnoreMessage)
	}
}
