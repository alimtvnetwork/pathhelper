package urischemes

import "gitlab.com/auk-go/core/constants"

type Type string

const (
	UriUnknown                Type = "*"
	UriSchemePrefixStandard   Type = constants.UriSchemePrefixStandard
	UriSchemePrefixTwoSlashes Type = constants.UriSchemePrefixTwoSlashes
)
