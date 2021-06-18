package pathstat

import "regexp"

var (
	bracketsMatcherWithContents = regexp.MustCompile(`\(.+\/.+\)`)
)
