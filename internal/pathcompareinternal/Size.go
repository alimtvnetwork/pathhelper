package pathcompareinternal

import (
	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
)

func Size(
	left, right int64,
) corecomparator.Compare {
	return corecmp.Integer64(left, right)
}
