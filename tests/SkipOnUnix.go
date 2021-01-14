package tests

import (
	"testing"

	"gitlab.com/evatix-go/core/coretests"
)

// Skip on Unix
func SkipOnUnix(t *testing.T) {
	coretests.SkipOnUnix(t)
}
