package enums

import "gitlab.com/evatix-go/pathhelper/constants"

type UriSchemes string

const (
	UriUnknown                UriSchemes = "*"
	UriSchemePrefixStandard   UriSchemes = constants.UriSchemePrefixStandard
	UriSchemePrefixTwoSlashes UriSchemes = constants.UriSchemePrefixStandard
)
