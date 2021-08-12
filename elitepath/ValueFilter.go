package elitepath

import "gitlab.com/evatix-go/core/enums/stringcompareas"

type ValueFilter struct {
	Value           string
	IsCaseSensitive bool
	Compare         stringcompareas.Variant
}

func (it *ValueFilter) IsMatch(content string) bool {
	if it == nil {
		return true
	}

	return it.Compare.IsCompareSuccess(
		content,
		it.Value,
		it.IsCaseSensitive)
}
