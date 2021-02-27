package expandpath

import (
	"regexp"

	"gitlab.com/evatix-go/core/regconsts"
)

var (
	regularExpressionForEachWordsWithDollarSymbol    = regexp.MustCompile(regconsts.RegExForEachWordsWithDollarSymbol)
	regularExpressionForEachWordsWithinPercentSymbol = regexp.MustCompile(regconsts.EachWordsWithinPercentSymbol)
)
