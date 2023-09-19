package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/normalize"
)

// JoinSimpleThree Doesn't apply normalize
func JoinSimpleThree(first, second, third string) string {
	return normalize.SimpleJoinPath3(first, second, third)
}
