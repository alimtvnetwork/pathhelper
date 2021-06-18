package pathstat

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func getRwxSimpleFromPathStatLines(splits []string, filePath string) *RwxSimple {
	// Access: (0755/drwxr-xr-x)
	leftRight := getLeftRightOfBracketsPatternData(splits[pathStatRwxIndex])

	if !leftRight.IsValid {
		return InvalidRwxSimple()
	}

	rwxWrapper, err := chmodhelper.NewUsingHyphenedRwxFullString(
		leftRight.Right)

	if err == nil {
		return &RwxSimple{
			HyphenedRwxValue: leftRight.Right,
			RwxWrapper:       &rwxWrapper,
			ErrorWrapper:     errorwrapper.StaticEmptyPtr,
			IsValid:          true,
		}
	}

	return &RwxSimple{
		HyphenedRwxValue: leftRight.Right,
		RwxWrapper:       &rwxWrapper,
		ErrorWrapper: errorwrapper.NewPath(
			errtype.ChmodInvalid,
			err,
			filePath),
		IsValid: false,
	}
}
