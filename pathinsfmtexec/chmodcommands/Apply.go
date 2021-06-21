package chmodcommands

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(commands *pathinsfmt.ChmodCommands) *errorwrapper.Wrapper {
	if commands == nil || commands.IsEmpty() {
		return errnew.EmptyPtr
	}

	collection := commands.CreateCmdOnceCollection()

	if commands.HasConditions() {
		// conditional execute
		return applyConditionCollection(
			commands.Condition,
			collection)
	}

	return collection.ExecuteAll()
}
