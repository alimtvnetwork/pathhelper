package checksummer

import (
	"gitlab.com/auk-go/pathhelper/hashas"
)

func hashAllSync(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	if isRecursive {
		return hashAllSyncRecursive(root, hashType)
	}

	return hashAllSyncNonRecursive(root, hashType)
}
