package pathinsfmt

import "gitlab.com/evatix-go/pathhelper/pathjoin"

type PathWithVerifier struct {
	IsNormalize    bool          `json:"IsNormalize,omitempty"`
	IsRecursive    bool          `json:"IsRecursive,omitempty"`
	IsExpandEnvVar bool          `json:"IsExpandEnvVar,omitempty"`
	IsSkipInvalid  bool          `json:"IsSkipInvalid,omitempty"`
	Verifier       *PathVerifier `json:"Verifier,omitempty"`
	Path           string        `json:"Path"` // warning : use CompiledPath
	compiledPath   *string
}

func (it *PathWithVerifier) IsVerifierUndefined() bool {
	return it == nil || it.Verifier == nil
}

func (it *PathWithVerifier) IsVerifierDefined() bool {
	return it != nil && it.Verifier != nil
}

func (it *PathWithVerifier) IsPathWithVerifierDefined() bool {
	return it != nil && it.Path != "" && it.Verifier != nil
}

func (it *PathWithVerifier) CompiledPath() string {
	if it.compiledPath != nil {
		return *it.compiledPath
	}

	compiledPath := pathjoin.FixPath(
		it.IsNormalize,
		it.IsExpandEnvVar,
		it.Path)

	it.compiledPath = &compiledPath

	return *it.compiledPath
}
