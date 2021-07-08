package main

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
)

func main() {
	verifiers := pathinsfmt.PathVerifiers{
		BaseSpecPlusRequestIds: coreinstruction.BaseSpecPlusRequestIds{},
		PathVerifiers: []pathinsfmt.PathVerifier{
			{
				UserGroupName: *pathinsfmt.NewUserGroupName(
					"alim", ""),
				BaseRwxInstructions: chmodins.BaseRwxInstructions{
					RwxInstructions: []chmodins.RwxInstruction{
						{
							RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
								Owner: "rwx",
								Group: "rw-",
								Other: "rw-",
							},
							Condition: chmodins.Condition{},
						},
					},
				},
			},
		},
		IsSkipCheckingOnInvalid: false,
		IsNormalize:             false,
		IsRecursiveCheck:        false,
	}

	locations := []string{
		os.TempDir(),
	}

	errorCollection := errwrappers.Empty()
	_ = pathmodifierverify.ApplyVerifier(
		true,
		true,
		true,
		true,
		&verifiers.PathVerifiers[0],
		errorCollection,
		locations)

	wkPath := os.TempDir() + "/main.json"
	errCollection2 := &errwrappers.Collection{}

	readWp := fs.
		ReadJsonParseSelfInjector(wkPath, errCollection2)
	fmt.Println(readWp)

	fmt.Println(errCollection2.IsSuccess(), errCollection2)

	// wr := fs.
	// 	WriteJsonResult(false, errorCollection.Json(),wkPath)

	// fmt.Println(wr)
}
