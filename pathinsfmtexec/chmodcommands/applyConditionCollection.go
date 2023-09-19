package chmodcommands

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
)

func applyConditionCollection(
	condition *chmodins.Condition,
	collection *errcmd.CmdOnceCollection,
) *errorwrapper.Wrapper {
	if condition.IsContinueOnError {
		return collection.ExecuteUntilErr()
	}

	return collection.ExecuteAll()
}
