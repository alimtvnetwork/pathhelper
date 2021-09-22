package pathcompare

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func LastModified(
	left, right *time.Time,
) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	leftNonPtr := *left
	rightNonPtr := *right

	if leftNonPtr.After(rightNonPtr) {
		return corecomparator.LeftGreater
	} else if leftNonPtr.Before(rightNonPtr) {
		return corecomparator.LeftLess
	} else if leftNonPtr.Equal(rightNonPtr) {
		return corecomparator.Equal
	}

	return corecomparator.NotEqual
}
