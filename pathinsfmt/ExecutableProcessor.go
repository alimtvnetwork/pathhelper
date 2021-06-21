package pathinsfmt

import (
	"log"

	"gitlab.com/evatix-go/core/enums/scripttype"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type ExecutableProcessor struct {
	Name                                        string   `json:"Name"`
	IsEnabled                                   bool     `json:"IsEnabled,omitempty"`
	IsNormalizePath                             bool     `json:"IsNormalizePath,omitempty"`
	FailedMessage                               string   `json:"FailedMessage,omitempty"`
	SuccessMessage                              string   `json:"SuccessMessage,omitempty"`
	OutputToFile                                string   `json:"OutputToFile,omitempty"`
	BinaryPath                                  string   `json:"BinaryPath,omitempty"`
	Args                                        []string `json:"Args,omitempty"`
	IsSecure, IsDisplayToConsole, IsWriteToFile bool
	ScriptType                                  scripttype.Variant `json:"ScriptType"`
	lazyCmdOnce                                 *errcmd.CmdOnce
}

func (e *ExecutableProcessor) GetExecuteOutputByExecuting(
	errWrapperCollection *errwrappers.Collection,
) *ProcessorExecOutput {
	cmdOnce := e.CreateCmdOnce()
	output := ProcessorExecOutput{
		CmdOnce:             cmdOnce,
		ConsoleResult:       cmdOnce.CompiledResult(),
		ExecutableProcessor: e,
		CustomMessage:       cmdOnce.CustomMessage,
	}

	if e.IsDisplayToConsole {
		log.Print(output.ConsoleResult.OutputString())
	}

	if e.IsWriteToFile {
		normalizePath := normalize.PathUsingSeparatorUsingSingleIf(
			e.IsNormalizePath,
			osconsts.PathSeparator,
			e.OutputToFile)

		writeErr := fsinternal.WriteFile(
			normalizePath,
			*output.ConsoleResult.OutputBytes())

		errWrapperCollection.AddWrapperPtr(writeErr)
	}

	errWrapperCollection.AddWrapperPtr(
		cmdOnce.CompiledErrorWrapper())

	return &output
}

func (e *ExecutableProcessor) LazyCmdOnce() *errcmd.CmdOnce {
	if e.lazyCmdOnce != nil {
		return e.lazyCmdOnce
	}

	e.lazyCmdOnce = e.CreateCmdOnce()

	return e.lazyCmdOnce
}

func (e *ExecutableProcessor) CreateCmdOnce() *errcmd.CmdOnce {
	argsCompiled := errcmd.ArgsJoinSlice(e.Args)
	normalizedBinaryPath := normalize.PathUsingSeparatorUsingSingleIf(
		e.IsNormalizePath,
		osconsts.PathSeparator,
		e.BinaryPath)

	script := errcmd.ArgsJoin(
		normalizedBinaryPath,
		argsCompiled)
	hasOutput := e.IsDisplayToConsole || e.IsWriteToFile

	cmdOnce := errcmd.NewCmdOnceUsingScriptType(
		hasOutput,
		e.IsSecure, e.ScriptType,
		script)

	cmdOnce.CustomMessage = &errcmd.CustomMessage{
		Success: e.SuccessMessage,
		Failed:  e.FailedMessage,
	}

	return cmdOnce
}
