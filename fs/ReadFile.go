package fs

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ReadFile(filePath string) *errbyte.Results {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &errbyte.Results{
			Values: &[]byte{},
			ErrorWrapper: errnew.MessagesPtr(
				errtype.ReadRequestFailed,
				"fs.ReadFile",
				filePath,
				err.Error()),
		}
	}

	if data == nil {
		return &errbyte.Results{
			Values: &[]byte{},
			ErrorWrapper: errnew.MessagesPtr(
				errtype.EmptyContent,
				"fs.ReadFile",
				filePath,
				"Path doesn't contain any valid data but nil."),
		}
	}

	return &errbyte.Results{
		Values:       &data,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
