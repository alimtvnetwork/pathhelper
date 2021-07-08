package pathregex

import (
	"gitlab.com/evatix-go/core/regexnew"
)

var (
	EachWordsWithDollarSymbol    = regexnew.NewMust(RegExForEachWordsWithDollarSymbolDefinition)
	EachWordsWithinPercentSymbol = regexnew.NewMust(EachWordsWithinPercentSymbolDefinition)
)
