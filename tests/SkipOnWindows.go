package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/coretests"
)

// Skip tests on Windows
func SkipOnWindows(t *testing.T) {
	coretests.SkipOnWindows(t)
}
