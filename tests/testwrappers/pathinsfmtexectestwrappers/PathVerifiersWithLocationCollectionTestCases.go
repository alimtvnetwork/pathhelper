package pathinsfmtexectestwrappers

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

var PathVerifiersWithLocationCollectionTestCases = []pathinsfmt.PathVerifiersWithLocationCollection{
	{
		PathVerifiers: &pathinsfmt.PathVerifiers{
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
		},
		LocationCollection: &pathinsfmt.LocationCollection{
			Locations:        nil,
			IsNormalizeApply: false,
		},
	},
}
