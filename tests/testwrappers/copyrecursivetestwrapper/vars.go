package copyrecursivetestwrapper

import (
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

var (
	SrcTestDataDir = pathjoin.JoinNormalized(
		currentDir(), SourceDataDirName)
	SourceRecursivePath = pathjoin.JoinNormalized(
		SrcTestDataDir,
		SourceRecursiveDirName)
	Destination = pathjoin.WithTempTest(PkgDirName)
)

func currentDir() string {
	_, f, _, _ := runtime.Caller(1)

	return filepath.Dir(f)
}
