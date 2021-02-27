package normalize

import "gitlab.com/evatix-go/core/constants"

var (
	normalizeMap = map[string]string{
		constants.UriSchemePrefixStandard:   "",
		constants.UriSchemePrefixTwoSlashes: "",
	}
	removeAndFixDoubleSeparatorToFinalSeparatorMap = map[string]string{
		constants.ForwardSlash:       constants.BackSlash,
		constants.DoubleForwardSlash: constants.BackSlash,
		constants.DoubleBackSlash:    constants.BackSlash,
	}
)
