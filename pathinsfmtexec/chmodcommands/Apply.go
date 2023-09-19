package chmodcommands

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func Apply(commands *pathinsfmt.ChmodCommands) *errorwrapper.Wrapper {
	if commands == nil || commands.IsEmpty() {
		return nil
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
