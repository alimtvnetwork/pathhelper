package pathcompareinternal

import "gitlab.com/evatix-go/core/corecomparator"

func SizePtr(
	left, right *int64,
) corecomparator.Compare {
	return corecomparator.Integer64Ptr(left, right)
}
