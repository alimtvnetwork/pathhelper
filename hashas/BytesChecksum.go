package hashas

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func BytesChecksum(method Variant, inputBytes []byte) *errbyte.Results {
	if inputBytes == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Null.WithMessage(
				"cannot perform BytesChecksum on null pointer!",
				inputBytes))
	}

	hashWriter, errWrap := method.NewHash()

	if errWrap.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			errWrap)
	}

	_, err := hashWriter.Write(inputBytes)
	if err != nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Error.TypeMessages(
				errtype.Hash,
				err,
				"writing hash hashWriter.Write(inputBytes)",
			))
	}

	hashedBytes := hashWriter.Sum(nil)

	return errbyte.New.Results.ValuesOnly(hashedBytes)
}
