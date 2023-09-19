package pathjoin

import (
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

// WithTemp
//
//	LongPath JoinFix, EnvVarExpand, skip on empty path
//
//	Windows
//	    - "%temp%\locations\..."
//	Unix
//	    - "/tmp/locations/..."
func WithTemp(
	locations ...string,
) string {
	return JoinFix(
		pathsconst.TempPermanentDir,
		locations...)
}
