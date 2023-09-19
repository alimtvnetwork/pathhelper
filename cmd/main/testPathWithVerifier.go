package main

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/pathmodifierverify"
)

func testPathWithVerifier() {
	ins := &pathinsfmt.PathWithVerifier{
		PathWithOptions: pathinsfmt.PathWithOptions{
			Path:          "/home/a/download_test",
			IsNormalize:   true,
			IsRecursive:   true,
			IsSkipInvalid: false,
		},
		Verifier: &pathinsfmt.PathVerifier{
			UserGroupName: pathinsfmt.UserGroupName{
				UserName: "root",
				BaseGroupName: pathinsfmt.BaseGroupName{
					GroupName: "root",
				},
			},
			BaseRwxInstructions: chmodins.BaseRwxInstructions{
				RwxInstructions: []chmodins.RwxInstruction{
					{
						RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
							Owner: "rw-",
							Group: "r--",
							Other: "r--",
						},
					},
				},
			},
		},
	}

	errColl := errwrappers.Empty()
	pathmodifierverify.ApplyPathWithVerifier(true, errColl, ins)
	errColl.HandleError()
}
