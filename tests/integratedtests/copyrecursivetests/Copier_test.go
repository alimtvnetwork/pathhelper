package copyrecursivetests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errverify"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers/copyrecursivetestwrapper"
)

func TestCopierRecursive(t *testing.T) {
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestIsRecursiveDir)

	errW.HandleError()

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(srcRoot, dstRoot, copyrecursive.Options{
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsClearDestination: false,
		IsUseShellOrCmd:    false,
		IsNormalize:        false,
		IsExpandVar:        false,
	}).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5)
	dstSums := checksummer.NewSync(true, dstRoot, hashas.Md5)

	// remove prefix
	strippedRootDstFiles := make(map[string][]byte, len(dstSums.GetMap()))
	for file, checksum := range dstSums.GetMap() {
		strippedRootDstFiles[strings.TrimPrefix(file, dstRoot)] = checksum
	}

	for file, checksum := range srcSums.GetMap() {
		stripped := strings.TrimPrefix(file, srcRoot)
		dstSum, isFound := strippedRootDstFiles[stripped]
		// Assert
		if !isFound {
			t.Errorf("expected %s in %s, but not found", stripped, dstRoot)
		}

		// Assert
		if !bytes.Equal(checksum, dstSum) {
			t.Errorf("checksum mismatch for %s file", file)
		}
	}
}

func TestCopierNonRecursive(t *testing.T) {
	convey.Convey("Testing copier non recursive", t, func() {

		// Arrange
		root, errW := copyrecursive.CopyToTempDir(TestIsRecursiveDir)
		if errW.HasError() {
			t.Error(errW.Error())
		}

		defer os.RemoveAll(root)

		srcRoot := filepath.Join(root, "src")
		dstRoot := filepath.Join(root, "dst")
		t.Log("src root directory:", srcRoot)
		t.Log("dst root directory:", dstRoot)

		// Act
		errWCopy := copyrecursive.NewCopier(srcRoot, dstRoot, copyrecursive.Options{
			IsSkipOnExist:      false,
			IsRecursive:        false,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        false,
			IsExpandVar:        false,
		}).Copy()

		if errWCopy.HasError() {
			t.Error(errWCopy.Error())
		}

		// Act checksum
		srcSums := checksummer.NewSync(false, srcRoot, hashas.Md5)
		dstSums := checksummer.NewSync(false, dstRoot, hashas.Md5)

		// remove prefix
		strippedRootDstFiles := make(map[string][]byte, len(dstSums.GetMap()))
		for file, checksum := range dstSums.GetMap() {
			strippedRootDstFiles[strings.TrimPrefix(file, dstRoot)] = checksum
		}

		for file, checksum := range srcSums.GetMap() {
			stripped := strings.TrimPrefix(file, srcRoot)
			dstSum, isFound := strippedRootDstFiles[stripped]
			// Assert
			convey.So(isFound, convey.ShouldBeTrue)
			if !isFound {
				t.Errorf("expected %s in %s, but not found", stripped, dstRoot)
			}

			// Assert
			convey.So(checksum, convey.ShouldResemble, dstSum)
			if !bytes.Equal(checksum, dstSum) {
				t.Errorf("checksum mismatch for %s file", file)
			}
		}
	})
}

func TestCopierSkipOnExist(t *testing.T) {
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestSkipOnExistDir)
	if errW.HasError() {
		t.Error(errW.Error())
	}

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(
		srcRoot,
		dstRoot,
		copyrecursive.Options{
			IsSkipOnExist:      true,
			IsRecursive:        true,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        false,
			IsExpandVar:        false,
		},
	).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5).StringHashesMap()
	dstSums := checksummer.NewSync(true, dstRoot, hashas.Md5).StringHashesMap()

	src_dir1_a_txt := filepath.Join(srcRoot, "dir1", "a.txt")
	src_dir2_b_txt := filepath.Join(srcRoot, "dir2", "b.txt")
	src_file1_txt := filepath.Join(srcRoot, "file1.txt")
	src_file2_txt := filepath.Join(srcRoot, "file2.txt")

	dst_dir1_a_txt := filepath.Join(dstRoot, "dir1", "a.txt")
	dst_dir2_b_txt := filepath.Join(dstRoot, "dir2", "b.txt")
	dst_file1_txt := filepath.Join(dstRoot, "file1.txt")
	dst_file2_txt := filepath.Join(dstRoot, "file2.txt")

	// Assert
	// Skip on exist should not copy this file
	if srcSums[src_dir1_a_txt] == dstSums[dst_dir1_a_txt] {
		t.Errorf("expected different checksum, but found equal,"+
			"\nsrc: %s\ndst: %s\n", src_dir1_a_txt, dst_dir1_a_txt)
	}

	if srcSums[src_file1_txt] == dstSums[dst_file1_txt] {
		t.Errorf("expected different checksum, but found equal,"+
			"\nsrc: %s\ndst: %s\n", src_dir1_a_txt, dst_dir1_a_txt)
	}

	if srcSums[src_dir2_b_txt] != dstSums[dst_dir2_b_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_dir2_b_txt, srcSums[src_dir2_b_txt], dst_dir2_b_txt, dstSums[dst_dir2_b_txt])
	}

	if srcSums[src_file2_txt] != dstSums[dst_file2_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_file2_txt, srcSums[src_file2_txt], dst_file2_txt, dstSums[dst_file2_txt])
	}
}

func TestCopierOverwrite(t *testing.T) {
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestSkipOnExistDir)
	if errW.HasError() {
		t.Error(errW.Error())
	}

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(
		srcRoot,
		dstRoot,
		copyrecursive.Options{
			IsSkipOnExist:      false,
			IsRecursive:        true,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        false,
			IsExpandVar:        false,
		},
	).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5).StringHashesMap()
	dstSums := checksummer.NewSync(true, dstRoot, hashas.Md5).StringHashesMap()

	src_dir1_a_txt := filepath.Join(srcRoot, "dir1", "a.txt")
	src_dir2_b_txt := filepath.Join(srcRoot, "dir2", "b.txt")
	src_file1_txt := filepath.Join(srcRoot, "file1.txt")
	src_file2_txt := filepath.Join(srcRoot, "file2.txt")

	dst_dir1_a_txt := filepath.Join(dstRoot, "dir1", "a.txt")
	dst_dir2_b_txt := filepath.Join(dstRoot, "dir2", "b.txt")
	dst_file1_txt := filepath.Join(dstRoot, "file1.txt")
	dst_file2_txt := filepath.Join(dstRoot, "file2.txt")

	// Assert
	if srcSums[src_dir1_a_txt] != dstSums[dst_dir1_a_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_dir1_a_txt, srcSums[src_dir1_a_txt], dst_dir1_a_txt, dstSums[dst_dir1_a_txt])
	}

	if srcSums[src_file1_txt] != dstSums[dst_file1_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_dir1_a_txt, srcSums[src_file1_txt], dst_dir1_a_txt, dstSums[dst_file1_txt])
	}

	if srcSums[src_dir2_b_txt] != dstSums[dst_dir2_b_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_dir2_b_txt, srcSums[src_dir2_b_txt], dst_dir2_b_txt, dstSums[dst_dir2_b_txt])
	}

	if srcSums[src_file2_txt] != dstSums[dst_file2_txt] {
		t.Errorf("expected same checksum, but found different,"+
			"\nsrc: %s sum: %s\ndst: %s sum: %s\n",
			src_file2_txt, srcSums[src_file2_txt], dst_file2_txt, dstSums[dst_file2_txt])
	}
}

func TestCopierMove(t *testing.T) {
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestIsRecursiveDir)
	if errW.HasError() {
		t.Error(errW.Error())
	}

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(srcRoot, dstRoot, copyrecursive.Options{
		IsMove:             true,
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsClearDestination: false,
		IsUseShellOrCmd:    false,
		IsNormalize:        false,
		IsExpandVar:        false,
	}).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5)

	// Assert
	if srcSums.Length() != 0 {
		t.Error("expected move, but didn't occur")
	}
}

func TestCopierRecursiveShell(t *testing.T) {
	coretests.SkipOnWindows(t)
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestIsRecursiveDir)
	if errW.HasError() {
		t.Error(errW.Error())
	}

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(srcRoot, dstRoot, copyrecursive.Options{
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsClearDestination: false,
		IsUseShellOrCmd:    true,
		IsNormalize:        false,
		IsExpandVar:        false,
	}).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5)
	dstSums := checksummer.NewSync(true, dstRoot, hashas.Md5)

	// remove prefix
	strippedRootDstFiles := make(map[string][]byte, len(dstSums.GetMap()))
	for file, checksum := range dstSums.GetMap() {
		strippedRootDstFiles[strings.TrimPrefix(file, dstRoot)] = checksum
	}

	for file, checksum := range srcSums.GetMap() {
		stripped := strings.TrimPrefix(file, srcRoot)
		dstSum, isFound := strippedRootDstFiles[stripped]
		// Assert
		if !isFound {
			t.Errorf("expected %s in %s, but not found", stripped, dstRoot)
		}

		// Assert
		if !bytes.Equal(checksum, dstSum) {
			t.Errorf("checksum mismatch for %s file", file)
		}
	}
}

func TestCopierMoveShell(t *testing.T) {
	coretests.SkipOnWindows(t)
	// Arrange
	root, errW := copyrecursive.CopyToTempDir(TestIsRecursiveDir)
	if errW.HasError() {
		t.Error(errW.Error())
	}

	defer os.RemoveAll(root)

	srcRoot := filepath.Join(root, "src")
	dstRoot := filepath.Join(root, "dst")
	t.Log("src root directory:", srcRoot)
	t.Log("dst root directory:", dstRoot)

	// Act
	errWCopy := copyrecursive.NewCopier(srcRoot, dstRoot, copyrecursive.Options{
		IsMove:             true,
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsClearDestination: false,
		IsUseShellOrCmd:    true,
		IsNormalize:        false,
		IsExpandVar:        false,
	}).Copy()

	if errWCopy.HasError() {
		t.Error(errWCopy.Error())
	}

	// Act checksum
	srcSums := checksummer.NewSync(true, srcRoot, hashas.Md5)

	// Assert
	if srcSums.Length() != 0 {
		t.Error("expected move, but didn't occur")
	}
}

func TestCopierRecursive_2(t *testing.T) {
	for _, testCase := range copyrecursivetestwrapper.CopyRecursiveTestCases {
		// Arrange
		header := testCase.Header +
			constants.Hyphen +
			"Src : " +
			testCase.CopyInstruction.Source +
			"Dest : " +
			testCase.CopyInstruction.Destination

		removeErr := deletepaths.RecursiveIf(
			testCase.IsClearDestinationPre,
			testCase.CopyInstruction.Destination)

		removeErr.HandleError()

		// Act
		errWCopy := testCase.
			CopyInstruction.
			Run()

		// Assert
		convey.Convey(header,
			t, func() {
				convey.So(
					errWCopy.IsSuccess(),
					convey.ShouldBeTrue,
				)
			})

		// Clear
		clearErr := deletepaths.RecursiveIf(
			testCase.IsClearDestinationPost,
			testCase.CopyInstruction.Destination)

		clearErr.HasError()
	}
}

func TestCopierRecursive_3(t *testing.T) {
	for i, testCase := range copyrecursivetestwrapper.CopyRecursiveTestCases {
		// Arrange
		isRecursive := testCase.
			CopyInstruction.
			IsRecursive
		src := testCase.CopyInstruction.Source
		dst := testCase.CopyInstruction.Destination
		header := testCase.Header +
			constants.Hyphen +
			"Src : " +
			src +
			"Dest : " +
			dst

		removeErr := deletepaths.RecursiveIf(
			testCase.IsClearDestinationPre,
			dst)

		removeErr.HandleError()

		// Act
		errWCopy := testCase.
			CopyInstruction.
			Run()

		errWCopy.HandleError()

		srcSums := checksummer.NewSync(isRecursive, src, hashas.Md5)
		dstSums := checksummer.NewSync(isRecursive, dst, hashas.Md5)
		checkSumErr := srcSums.VerifyError(
			true, // isContinue on error
			true, // is trim root
			dstSums)

		errVerifyParams := &errverify.VerifyCollectionParams{
			CaseIndex:       i,
			FuncName:        testCase.ErrorCollectionVerifier.FunctionName,
			TestCaseName:    "CopyRecursiveTestCases",
			ErrorCollection: checkSumErr,
		}

		isSuccess := testCase.
			ErrorCollectionVerifier.
			IsMatch(errVerifyParams)

		// Assert
		convey.Convey(header+testCase.ErrorCollectionVerifier.Header,
			t, func() {
				convey.So(
					isSuccess,
					convey.ShouldBeTrue,
				)
			})

		// Clear
		clearErr := deletepaths.RecursiveIf(
			testCase.IsClearDestinationPost,
			dst)

		clearErr.HasError()
	}
}
