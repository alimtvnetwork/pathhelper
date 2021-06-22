package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec"
)

func main() {
	symLinks := pathinsfmt.SymbolicLinks{
		BaseSpecPlusRequestIds: coreinstruction.BaseSpecPlusRequestIds{},
		IsContinueOnError:      true,
		SymbolicLinks: []pathinsfmt.SymbolicLink{
			// {
			// 	Src:           "d:\\notes.md",
			// 	Dst:           "d:\\alim-sym-link\\something-else\\some-thing3\\main.notes.md",
			// 	IsClearBefore: false,
			// 	IsSkipOnExist: false,
			// 	IsMkDirAll:    true,
			// 	IsSkipOnSrcMissing: true,
			// },
			{
				Src:                "d:\\notes.md",
				Dst:                "d:\\alim-sym-link\\something-else\\main.notes.md",
				IsClearBefore:      false,
				IsSkipOnExist:      false,
				IsMkDirAll:         true,
				IsSkipOnSrcMissing: true,
			},
			// {
			// 	Src:           "d:\\notes.md",
			// 	Dst:           "d:\\alim-sym-link\\mode2.md",
			// 	IsClearBefore: true,
			// 	IsSkipOnExist: true,
			// 	IsMkDirAll:    true,
			// 	IsSkipOnSrcMissing: true,
			// },
			// {
			// 	Src:           "d:\\notes.md",
			// 	Dst:           "d:\\alim-sym-link\\something-else\\main.notes4.md",
			// 	IsClearBefore: true,
			// 	IsSkipOnExist: true,
			// 	IsMkDirAll:    true,
			// 	IsSkipOnSrcMissing: false,
			// },
		},
	}

	fmt.Println(pathinsfmtexec.ApplySymbolicLinks(&symLinks))
}
