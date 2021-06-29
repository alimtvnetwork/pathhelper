package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type LineIdentifier struct {
	LineNumber int  `json:"LineNumber,omitempty"`
	IsNewLine  bool `json:"IsNewLine,omitempty"`
}

func (it *LineIdentifier) IsLineInvalid() bool {
	return it == nil || it.LineNumber < 0
}

func (it *LineIdentifier) HasLineNumber() bool {
	return it != nil && it.LineNumber > constants.InvalidValue
}

func (it *LineIdentifier) IsNewLineAdd() bool {
	return it.IsLineInvalid() && it.IsNewLine
}

func (it *LineIdentifier) ToBaseLineIdentifier() *BaseLineIdentifier {
	if it == nil {
		return nil
	}

	return NewBaseLineIdentifier(it.LineNumber)
}

func (it *LineIdentifier) Clone() *LineIdentifier {
	if it == nil {
		return nil
	}

	return &LineIdentifier{
		LineNumber: it.LineNumber,
		IsNewLine:  it.IsNewLine,
	}
}
