package pathfixer

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type PathOptions struct {
	IsNormalize     bool `json:"IsNormalize,omitempty"`
	IsExpandEnvVar  bool `json:"IsExpandEnvVariable,omitempty"` // can expand environment variables with %{Name} $Name %{Java_home} ${Java_HOME}
	IsRecursive     bool `json:"IsRecursive,omitempty"`
	IsSkipOnInvalid bool `json:"IsSkipOnInvalid,omitempty"`
	IsSkipOnExist   bool `json:"IsSkipOnExist,omitempty"`
	IsSkipOnEmpty   bool `json:"IsSkipOnEmpty,omitempty"`
	IsRelative      bool `json:"IsRelative,omitempty"`
}

func (it *PathOptions) GetFixedPath(location string) string {
	return pathjoin.FixPath(
		it.IsNormalize,
		it.IsExpandEnvVar,
		location)
}

func (it *PathOptions) GetFixedPathJoined(location1, location2 string) string {
	return pathjoin.JoinConditionalNormalizedExpandIf(
		it.IsNormalize,
		it.IsExpandEnvVar,
		location1,
		location2)
}

func (it *PathOptions) GetFixedPathJoinedMany(
	locations ...string,
) string {
	return pathjoin.JoinWithSep(
		it.IsSkipOnEmpty,
		it.IsExpandEnvVar,
		it.IsNormalize,
		osconsts.PathSeparator,
		locations...)
}

func (it *PathOptions) ClonePathOptions() *PathOptions {
	if it == nil {
		return nil
	}

	return &PathOptions{
		IsNormalize:     it.IsNormalize,
		IsRelative:      it.IsRelative,
		IsExpandEnvVar:  it.IsExpandEnvVar,
		IsSkipOnInvalid: it.IsSkipOnInvalid,
		IsSkipOnExist:   it.IsSkipOnExist,
		IsSkipOnEmpty:   it.IsSkipOnEmpty,
	}
}
