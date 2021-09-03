package hashas

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type Variant byte

const (
	Undefined Variant = iota
	Md5
	Sha1
	Sha256
	Sha512
)

func (it Variant) NewHash() (hash.Hash, *errorwrapper.Wrapper) {
	switch it {
	case Undefined:
		return nil, errnew.MessagesPtr(
			errtype.UnexpectedDefinition,
			it.Name()+" is expected to be not defined. Thus nil hasher.",
		)
	case Md5:
		return md5.New(), errnew.EmptyPtr
	case Sha1:
		return sha1.New(), errnew.EmptyPtr
	case Sha256:
		return sha256.New(), errnew.EmptyPtr
	case Sha512:
		return sha512.New(), errnew.EmptyPtr
	default:
		return nil, errnew.MessagesPtr(
			errtype.InvalidOption,
			it.Name()+" invalid option.",
			BasicEnumImpl.RangesInvalidMessage(),
		)
	}
}

func (it *Variant) StringSumOf(
	inputBytes []byte,
) *errstr.Result {
	outputBytesResults := it.SumOf(inputBytes)
	toString := outputBytesResults.NonEmptyString(convertBytesResultsToEncodedHexString)

	return &errstr.Result{
		Value:        toString,
		ErrorWrapper: outputBytesResults.ErrorWrapper,
	}
}

func (it *Variant) StringSumOfFile(
	fileName string,
) *errstr.Result {
	byteResults := it.SumOfFile(fileName)
	if byteResults.HasError() {
		return &errstr.Result{
			ErrorWrapper: byteResults.ErrorWrapper,
		}
	}

	toString := byteResults.NonEmptyString(
		convertBytesResultsToEncodedHexString)

	return errstr.EmptyErrorResult(toString)
}

func (it *Variant) SumOfFile(
	fileName string,
) *errbyte.Results {
	if fileName == constants.EmptyString {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyString,
				"File name is empty"))
	}

	hashWriter, errWp := it.NewHash()

	if errWp.HasError() {
		return errbyte.EmptyResultsWithError(errWp)
	}

	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		return errbyte.EmptyResultsWithError(
			errnew.Path(
				errtype.FileRead,
				errOpen,
				"Error opening file : "+fileName,
			))
	}

	defer file.Close()

	_, errCopy := io.Copy(hashWriter, file)
	if errCopy != nil {
		return errbyte.EmptyResultsWithError(
			errnew.Path(
				errtype.Copy,
				errOpen,
				"Error copying to  file : "+fileName,
			))
	}

	hashedBytes := hashWriter.Sum(nil)

	return errbyte.EmptyErrorResults(hashedBytes...)
}

func (it *Variant) SumOf(
	inputBytes []byte,
) *errbyte.Results {
	if inputBytes == nil {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyPointerOrNullPointer,
				"Cannot perform SumOf on Nil Pointer!"))
	}

	hashWriter, errWp := it.NewHash()

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

func (it *Variant) SumOfErrorBytes(
	errBytes *errbyte.Results,
) *errbyte.Results {
	if errBytes == nil || errBytes.Values == nil {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyPointerOrNullPointer,
				"Cannot perform SumOfErrorBytes on Nil Pointer!"))
	}

	if errBytes.HasError() {
		return errBytes
	}

	hashWriter, errWp := it.NewHash()

	if errWp.HasError() {
		return errbyte.EmptyResultsWithError(errWp)
	}

	hashedBytes := hashWriter.Sum(
		errBytes.Values)

	return errbyte.EmptyErrorResults(hashedBytes...)
}

func (it *Variant) SumOfJsonResult(
	result *corejson.Result,
) *errbyte.Results {
	if result.HasError() {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.JsonSyntaxIssue,
				"cannot hash on error json results!",
				result.MeaningfulError().Error()))
	}

	hashWriter, errWp := it.NewHash()

	if errWp.HasError() {
		return errbyte.EmptyResultsWithError(errWp)
	}

	hashedBytes := hashWriter.Sum(
		result.ValueMust())

	return errbyte.EmptyErrorResults(hashedBytes...)
}

func (it *Variant) HexOfJsonResult(
	result *corejson.Result,
) *errstr.Result {
	bytesResult := it.SumOfJsonResult(result)

	if bytesResult.HasError() {
		return errstr.ErrorWrapper(bytesResult.ErrorWrapper)
	}

	toString := bytesResult.NonEmptyString(
		convertBytesResultsToEncodedHexString)

	return errstr.EmptyErrorResult(toString)
}

func (it Variant) IsUndefined() bool {
	return it == Undefined
}

func (it Variant) IsMd5() bool {
	return it == Md5
}

func (it Variant) IsSha1() bool {
	return it == Sha1
}

func (it Variant) IsSha256() bool {
	return it == Sha256
}

func (it Variant) IsSha512() bool {
	return it == Sha512
}

func (it *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it *Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	byteVal, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(byteVal)
	}

	return err
}

func (it *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
