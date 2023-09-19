package pathmodifier

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathchmod"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/chmodcommands"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/chowninsexec"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/namegroup"
)

func ApplySimple(
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	modifier *pathinsfmt.PathModifier,
	flatPaths []string,
) (isSuccess bool) {
	stateTracker := errCollection.StateTracker()

	if modifier == nil || len(flatPaths) == 0 {
		return true
	}

	if modifier.HasChmodCommands() {
		errCollection.AddWrapperPtr(
			chmodcommands.Apply(
				modifier.ChmodCommands),
		)
	}

	if modifier.HasChown() {
		errCollection.AddWrapperPtr(
			chowninsexec.Apply(
				isContinueOnErr,
				modifier.Chown,
				flatPaths),
		)
	}

	if modifier.HasChangeGroup() {
		errCollection.AddWrapperPtr(
			namegroup.ApplyLinuxOnlyGroup(
				modifier.ChangeGroup.IsRecursive,
				isContinueOnErr,
				&modifier.ChangeGroup.BaseGroupName,
				flatPaths...),
		)
	}

	if modifier.HasChangeGroup() {
		errCollection.AddWrapperPtr(
			namegroup.ApplyLinuxOnlyGroup(
				modifier.ChangeGroup.IsRecursive,
				isContinueOnErr,
				&modifier.ChangeGroup.BaseGroupName,
				flatPaths...),
		)
	}

	if modifier.HasRwxInstructions() {
		errCollection.AddWrapperPtr(
			pathchmod.ApplyChmodRwxInstructions(
				&modifier.BaseRwxInstructions,
				flatPaths),
		)
	}

	return stateTracker.IsSuccess()
}
