package hashas

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func BytesChecksum(method Variant, inputBytes []byte) *errbyte.Results {
	if inputBytes == nil {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyPointerOrNullPointer,
				"Cannot perform SumOf on Nil Pointer!"))
	}

	hashWriter, errWp := method.NewHash()

	if errWp.HasError() {
		return errbyte.EmptyResultsWithError(
			errWp)
	}

	_, err := hashWriter.Write(inputBytes)
	if err != nil {
		return errbyte.EmptyResultsWithError(
			errnew.ErrorWithMessagesPtr(
				errtype.Hash,
				err,
				"writing hash hashWriter.Write(inputBytes)",
			))
	}

	hashedBytes := hashWriter.Sum(nil)

	return errbyte.EmptyErrorResults(hashedBytes...)
}
