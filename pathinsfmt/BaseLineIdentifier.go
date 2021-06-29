package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type BaseLineIdentifier struct {
	LineIdentifier
}

func NewBaseLineIdentifier(lineNumber int) *BaseLineIdentifier {
	return &BaseLineIdentifier{
		LineIdentifier{
			LineNumber: lineNumber,
			IsNewLine:  lineNumber < constants.Zero,
		},
	}
}

func (it *BaseLineIdentifier) ToNewLineIdentifier() *LineIdentifier {
	if it == nil {
		return nil
	}

	return &LineIdentifier{
		LineNumber: it.LineNumber,
		IsNewLine:  it.IsNewLine,
	}
}

func (it *BaseLineIdentifier) Clone() *BaseLineIdentifier {
	if it == nil {
		return nil
	}

	return &BaseLineIdentifier{
		LineIdentifier{
			LineNumber: it.LineNumber,
			IsNewLine:  it.IsNewLine,
		},
	}
}
