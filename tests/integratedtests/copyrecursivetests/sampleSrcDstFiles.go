package copyrecursivetests

import "path/filepath"

func sampleSrcDstFiles(
	srcRoot, dstRoot string,
) (srcFiles, dstFiles []string) {
	srcFiles = []string{
		filepath.Join(srcRoot, "dir1", "a.txt"),
		filepath.Join(srcRoot, "dir2", "b.txt"),
		filepath.Join(srcRoot, "file1.txt"),
		filepath.Join(srcRoot, "file2.txt"),
	}

	dstFiles = []string{
		filepath.Join(dstRoot, "dir1", "a.txt"),
		filepath.Join(dstRoot, "dir2", "b.txt"),
		filepath.Join(dstRoot, "file1.txt"),
		filepath.Join(dstRoot, "file2.txt"),
	}

	return srcFiles, dstFiles
}
