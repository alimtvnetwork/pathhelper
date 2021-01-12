package enums

import "gitlab.com/evatix-go/core/constants"

type UriSchemes string

const (
	UriUnknown                UriSchemes = "*"
	UriSchemePrefixStandard   UriSchemes = constants.UriSchemePrefixStandard
	UriSchemePrefixTwoSlashes UriSchemes = constants.UriSchemePrefixStandard
)
