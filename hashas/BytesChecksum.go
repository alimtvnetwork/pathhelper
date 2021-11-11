package hashas

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func BytesChecksum(method Variant, inputBytes []byte) *errbyte.Results {
	if inputBytes == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Messages.Many(
				errtype.EmptyPointerOrNullPointer,
				"Cannot perform SumOf on Nil Pointer!"))
	}

	hashWriter, errWp := method.NewHash()

	if errWp.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			errWp)
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
