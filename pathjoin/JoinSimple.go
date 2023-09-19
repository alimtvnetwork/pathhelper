package pathjoin

import (
	"gitlab.com/auk-go/core/osconsts"
)

// JoinSimple doesn't apply normalize
func JoinSimple(first, second string) string {
	return first + osconsts.PathSeparator + second
}
