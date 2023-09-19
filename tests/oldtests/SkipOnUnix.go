package oldtests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests"
)

// SkipOnUnix Skip on Unix
func SkipOnUnix(t *testing.T) {
	coretests.SkipOnUnix(t)
}
