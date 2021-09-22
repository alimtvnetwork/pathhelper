package pathcompare

import (
	"gitlab.com/evatix-go/core/corecomparator"
)

func Size(
	left, right *int64,
) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	leftNonPtr := *left
	rightNonPtr := *right

	if leftNonPtr > rightNonPtr {
		return corecomparator.LeftLess
	} else if leftNonPtr < rightNonPtr {
		return corecomparator.LeftLess
	} else if leftNonPtr == rightNonPtr {
		return corecomparator.Equal
	}

	return corecomparator.NotEqual
}
