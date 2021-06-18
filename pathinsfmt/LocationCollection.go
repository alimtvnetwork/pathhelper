package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type LocationCollection struct {
	Locations        []string `json:"Locations,omitempty"`
	IsNormalizeApply bool     `json:"IsNormalizeApply"`
}

func (receiver *LocationCollection) Length() int {
	if receiver == nil {
		return constants.Zero
	}

	return len(receiver.Locations)
}

func (receiver *LocationCollection) IsEmpty() bool {
	return receiver.Length() == constants.Zero
}

func (receiver *LocationCollection) HasAnyItem() bool {
	return receiver.Length() > constants.Zero
}
