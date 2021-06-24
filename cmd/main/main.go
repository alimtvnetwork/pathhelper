package main

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
)

func main() {
	verifiers := pathinsfmt.PathVerifiers{
		BaseSpecPlusRequestIds: coreinstruction.BaseSpecPlusRequestIds{},
		PathVerifiers: []pathinsfmt.PathVerifier{
			{
				BaseUserNamePlusGroupName: *pathinsfmt.NewBaseUserNamePlusGroupName(
					"", ""),
				BaseRwxInstructions: chmodins.BaseRwxInstructions{
					RwxInstructions: &[]*chmodins.RwxInstruction{
						{
							RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
								Owner: "rw-",
								Group: "rw-",
								Other: "rw-",
							},
							Condition: chmodins.Condition{},
						},
					},
				},
				IsSkipCheckingOnNonExist: false,
				IsNormalize:              true,
				IsRecursiveCheck:         true,
			},
		},
	}

	fmt.Println(pathmodifierverify.ApplyVerifierDirect(
		&verifiers.PathVerifiers[0],
		os.TempDir()))
}
