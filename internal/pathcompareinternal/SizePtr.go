package pathcompareinternal

import (
	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
)

func SizePtr(
	left, right *int64,
) corecomparator.Compare {
	return corecmp.Integer64Ptr(left, right)
}
