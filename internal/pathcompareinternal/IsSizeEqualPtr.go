package pathcompareinternal

import "gitlab.com/evatix-go/core/corecomparator"

func IsSizeEqualPtr(
	left, right *int64,
) bool {
	return corecomparator.Integer64Ptr(left, right).IsEqual()
}
