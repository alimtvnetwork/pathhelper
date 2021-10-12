package pathcompareinternal

import (
	"gitlab.com/evatix-go/core/corecomparator"
)

func Size(
	left, right int64,
) corecomparator.Compare {
	return corecomparator.Integer64(left, right)
}
