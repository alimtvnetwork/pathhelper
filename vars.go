package pathhelper

import "gitlab.com/auk-go/core/constants"

var (
	uriRemovePrefixes = []string{
		constants.UriSchemePrefixStandard,
		constants.UriSchemePrefixTwoSlashes,
	}
)
