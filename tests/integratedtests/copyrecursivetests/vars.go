package copyrecursivetests

import (
	"path/filepath"

	"gitlab.com/auk-go/pathhelper/tests/testwrappers/copyrecursivetestwrapper"
)

var (
	TestIsRecursiveDir = filepath.Join(
		copyrecursivetestwrapper.SrcTestDataDir,
		copyrecursivetestwrapper.SourceRecursiveDirName)
	TestSkipOnExistDir = filepath.Join(
		copyrecursivetestwrapper.SrcTestDataDir,
		copyrecursivetestwrapper.SourceIsSkipOnExistDirName)
)
