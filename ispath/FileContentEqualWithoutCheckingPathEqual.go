package ispath

import "gitlab.com/auk-go/pathhelper/internal/ispathinternal"

func FileContentEqualWithoutCheckingPathEqual(
	leftFullPath string,
	rightFullPath string,
) bool {
	return ispathinternal.FileContentEqualWithoutCheckingPathEqual(
		leftFullPath,
		rightFullPath)
}
