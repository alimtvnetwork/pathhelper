package copyrecursive

import "gitlab.com/auk-go/core/filemode"

const (
	defaultFileMode  = filemode.X755
	mvCommand        = "mv"
	cpCommand        = "cp"
	cpArgSkipOnExist = "-n"
	cpArgRecursive   = "-r"
	cpArgRemoveDst   = "--remove-destination"
	mvArgSkipOnExist = "-n"
)
