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
		IsContinueOnError:      false,
		SymbolicLinks: []pathinsfmt.SymbolicLink{
			{
				Src:           "d:\\notes.md",
				Dst:           "d:\\alim-sym-link\\something-else\\some-thing3\\main.notes.md\\wdw",
				IsClearBefore: false,
				IsSkipOnExist: false,
				IsMkDirAll:    true,
				IsSkipOnSrcMissing: true,
			},
			{
				Src:           "d:\\notes.mdx",
				Dst:           "d:\\alim-sym-link\\something-else\\main.notes.md",
				IsClearBefore: true,
				IsSkipOnExist: true,
				IsMkDirAll:    true,
				IsSkipOnSrcMissing: true,
			},
			{
				Src:           "d:\\notes.md",
				Dst:           "d:\\alim-sym-link\\something-else\\main.notes3.md",
				IsClearBefore: false,
				IsSkipOnExist: true,
				IsMkDirAll:    true,
				IsSkipOnSrcMissing: false,
			},
			{
				Src:           "d:\\notes.md",
				Dst:           "d:\\alim-sym-link\\something-else\\main.notes4.md",
				IsClearBefore: true,
				IsSkipOnExist: true,
				IsMkDirAll:    true,
				IsSkipOnSrcMissing: false,
			},
		},
	}

	fmt.Println(pathinsfmtexec.ApplySymbolicLinks(&symLinks))
}
