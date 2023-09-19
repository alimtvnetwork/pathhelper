package main

import (
	"fmt"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/pathhelper/hashas"
	"gitlab.com/auk-go/pathhelper/hexchecksum"
	"gitlab.com/auk-go/pathhelper/pathsysinfo"
)

func checkSumCheck() {
	result := hexchecksum.OfFilesContentsAsync(
		false,
		false,
		hashas.Sha256,
		"cmd/main/main.go")

	// result.ErrorWrapper.HandleError()
	result.ErrorWrapper.Log()
	fmt.Println(result.String())

	rs := pathsysinfo.GetPathUserGroupId("cmd/main")

	fmt.Println(converters.Any.ToFullNameValueString(rs))
	fmt.Println(rs.Error)
}
