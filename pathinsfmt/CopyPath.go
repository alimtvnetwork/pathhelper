package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type CopyPath struct {
	Source, Destination string
	IsRecursive         bool
	IsClearBeforeCopy   bool
	IsNormalize         bool
	IsExpand            bool
}

func (it CopyPath) DestinationFixedPath() string {
	return pathjoin.JoinConditionalNormalizedExpandIf(
		it.IsNormalize,
		it.IsExpand,
		it.Destination,
		constants.EmptyString)
}

func (it CopyPath) SourceFixedPath() string {
	return pathjoin.JoinConditionalNormalizedExpandIf(
		it.IsNormalize,
		it.IsExpand,
		it.Source,
		constants.EmptyString)
}
