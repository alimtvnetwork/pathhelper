package pathregex

import "gitlab.com/evatix-go/core/regconsts"

const (
	identifer                                   = regconsts.AnyIdentifier
	RegExForEachWordsWithDollarSymbolDefinition = `(\$\{` + identifer + `\}|\$` + identifer + `)+` // Selects a full word that starts with a "$" symbol
	EachWordsWithinPercentSymbolDefinition      = `(\%\{` + identifer + `\}|\%` + identifer + `)+`  // Selects a full word that is within two "%" symbol
)
