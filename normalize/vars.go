package normalize

import (
	"os"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/pathhelper/internal/consts"
)

var (
	normalizeMap = map[string]string{
		constants.UriSchemePrefixStandard:   constants.EmptyString,
		constants.UriSchemePrefixTwoSlashes: constants.EmptyString,
	}

	removeAndFixDoubleSeparatorToFinalSeparatorMap = map[string]string{
		constants.ForwardSlash:       constants.BackSlash,
		constants.DoubleForwardSlash: constants.BackSlash,
		constants.DoubleBackSlash:    constants.BackSlash,
	}

	uncPrefixes = []string{
		constants.LongPathUncPrefix,
		constants.LongPathQuestionMarkPrefix,
		consts.BrokenLongPathQuestionMarkPrefix,
		consts.BrokenLongPathUncPrefix,
	}

	PathSeparatorChar = byte(os.PathSeparator)
)
