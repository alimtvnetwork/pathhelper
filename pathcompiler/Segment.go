package pathcompiler

import "gitlab.com/auk-go/core/coredata/corestr"

type Segment struct {
	Name, Format   string
	compiledFormat corestr.SimpleStringOnce
}
