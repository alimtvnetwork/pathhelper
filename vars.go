package pathhelper

import "gitlab.com/evatix-go/core/constants"

var (
	slugForbiddenArray = []string{
		" ",
		"!",
		"`",
		"@",
		"#",
		"%",
		"$",
		"^",
		"&",
		"*",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
	}

	uriRemovePrefixes = []string{
		constants.UriSchemePrefixStandard,
		constants.UriSchemePrefixTwoSlashes,
	}
)
