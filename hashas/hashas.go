package hashas

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

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
			it.Name()+"(HashMethod/Variant) is expected to be not defined. Thus nil hasher.",
		)
	case Md5:
		return md5.New(), nil
	case Sha1:
		return sha1.New(), nil
	case Sha256:
		return sha256.New(), nil
	case Sha512:
		return sha512.New(), nil
	default:
		return nil, errnew.MessagesPtr(
			errtype.InvalidOption,
			it.Name()+" invalid option.",
			BasicEnumImpl.RangesInvalidMessage(),
		)
	}
}

func (it Variant) NewHashError() (hash.Hash, error) {
	switch it {
	case Undefined:
		return nil, errtype.UnexpectedDefinition.ReferencesCsvError(
			"(HashMethod/Variant) is expected to be not defined. Thus nil hasher.",
			it.Name(),
		)
	case Md5:
		return md5.New(), nil
	case Sha1:
		return sha1.New(), nil
	case Sha256:
		return sha256.New(), nil
	case Sha512:
		return sha512.New(), nil
	default:
		return nil, errtype.InvalidOption.ReferencesCsvError(
			BasicEnumImpl.RangesInvalidMessage(),
			it.Name(),
		)
	}
}

func (it Variant) HexSumOf(
	inputBytes []byte,
) *errstr.Result {
	return HexChecksumOfRawBytes(it, inputBytes)
}

func (it Variant) HexSumOfFile(
	fileName string,
) *errstr.Result {
	return HexChecksumOfFilePath(it, fileName)
}

func (it Variant) HexSumOfFileNoError(
	fullPath string,
) string {
	return HexChecksumOfFilePathNoError(
		false,
		it,
		fullPath)
}

func (it Variant) HexSumOfFileNoErrorIf(
	isSkipGenerate bool,
	fullPath string,
) string {
	if isSkipGenerate {
		return constants.EmptyString
	}

	return HexChecksumOfFilePathNoError(
		false,
		it,
		fullPath)
}

func (it Variant) HexSumOfFileIf(
	isSkipGenerate bool,
	fullPath string,
) *errstr.Result {
	if isSkipGenerate {
		return errstr.Empty()
	}

	return HexChecksumOfFilePath(
		it,
		fullPath)
}

func (it Variant) SumOfFile(
	filePath string,
) *errbyte.Results {
	return SumOfFile(it, filePath)
}

func (it Variant) SumOf(
	inputBytes []byte,
) *errbyte.Results {
	return BytesChecksum(it, inputBytes)
}

func (it Variant) SumOfErrorBytes(
	errBytes *errbyte.Results,
) *errbyte.Results {
	return ErrorWrapperWithBytesChecksum(it, errBytes)
}

func (it *Variant) HexSumOfAny(
	item interface{},
) *errstr.Result {
	jsonResult := corejson.NewFromAnyPtr(item)

	return it.HexOfJsonResult(jsonResult)
}

func (it *Variant) HexSumOfAnyIf(
	isGenerate bool,
	item interface{},
) *errstr.Result {
	if isGenerate {
		jsonResult := corejson.NewFromAnyPtr(item)

		return it.HexOfJsonResult(jsonResult)
	}

	return errstr.Empty()
}

func (it Variant) HexSumOfAnyItemsToCombinedSingleString(
	isSkipOnNil bool,
	items ...interface{},
) *errstr.Result {
	return HexChecksumOfAnyItemsToCombinedSingleString(
		isSkipOnNil,
		it,
		items...)
}

func (it Variant) HexSumOfAnyItems(
	isSkipOnNil bool,
	items ...interface{},
) *errstr.Results {
	return HexChecksumOfAnyItems(
		isSkipOnNil,
		it,
		items...)
}

func (it *Variant) SumOfJsonResult(
	result *corejson.Result,
) *errbyte.Results {
	if result == nil || result.Bytes == nil {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.EmptyPointerOrNullPointer,
				"cannot hash nil json result or nil bytes values!",
			))
	}

	if result.HasError() {
		return errbyte.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.JsonSyntaxIssue,
				"cannot hash on error json results!",
				result.MeaningfulError().Error()))
	}

	return it.SumOf(result.ValueMust())
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

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
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
