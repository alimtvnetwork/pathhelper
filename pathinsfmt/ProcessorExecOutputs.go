package pathinsfmt

import "gitlab.com/auk-go/errorwrapper"

type ProcessorExecOutputs struct {
	CompiledError *errorwrapper.Wrapper
	Outputs       []ProcessorExecOutput
}
