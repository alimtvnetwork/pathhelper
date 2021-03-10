package consts

import "gitlab.com/evatix-go/core/coredata/corestr"

var (
	SlugForbiddenArray = []string{
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

	SlugHashset = corestr.NewHashsetUsingStrings(
		&SlugForbiddenArray,
		0,
		false)
)
