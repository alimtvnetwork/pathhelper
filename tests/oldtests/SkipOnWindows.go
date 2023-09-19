package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests"
)

// SkipOnWindows Skip tests on Windows
func SkipOnWindows(t *testing.T) {
	coretests.SkipOnWindows(t)
}
