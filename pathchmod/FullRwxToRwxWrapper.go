package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func FullRwxToRwxWrapper(rwxFull string) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	varWrapper, err := chmodhelper.NewRwxVariableWrapper(rwxFull)

	if err != nil {
		return nil, errnew.RefsWithMessage(
			errtype.ChmodInvalid,
			err.Error(),
			ref.Value{
				Variable: "rwx",
				Value:    rwxFull,
			})
	}

	if varWrapper == nil || !varWrapper.IsFixedType() {
		return nil, errnew.RefsWithMessage(
			errtype.ChmodInvalid,
			"Rwx must be be fixed without wildcard to receive file mode.",
			ref.Value{
				Variable: "rwx",
				Value:    rwxFull,
			})
	}

	return varWrapper.
		ToCompileFixedPtr(), errnew.EmptyPtr
}
