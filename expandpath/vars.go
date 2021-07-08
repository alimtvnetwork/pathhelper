package expandpath

import "gitlab.com/evatix-go/pathhelper/pathregex"

var (
	regexEachWordPercent = pathregex.EachWordsWithinPercentSymbol
	regexEachWordDollar  = pathregex.EachWordsWithDollarSymbol
)
