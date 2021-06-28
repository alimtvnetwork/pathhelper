package pathinsfmtexectests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers"
)

func Test_ApplyPathVerifierErrorsUnix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	locations := testwrappers.SetupDefaultPathsUnix()
	verifier := &pathinsfmt.PathVerifier{
		UserGroupName: *testwrappers.DefaultUserNameGroupName,
		BaseRwxInstructions: chmodins.BaseRwxInstructions{
			RwxInstructions: []chmodins.RwxInstruction{
				{
					RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
					Condition: chmodins.Condition{
						IsSkipOnInvalid:   false,
						IsContinueOnError: false,
						IsRecursive:       false,
					},
				},
			},
		},
	}
	errCollection := errwrappers.Empty()

	// Act
	isSuccess := pathmodifierverify.ApplyVerifier(true,
		true,
		false,
		true,
		verifier,
		errCollection,
		locations)

	// Assert
	Convey("Default Paths Create", t, func() {
		So(errCollection.String(), ShouldBeEmpty)
		So(errCollection.IsEmpty(), ShouldBeTrue)
		So(isSuccess, ShouldBeTrue)
	})
}

func Test_ApplyPathVerifierErrorsWindows(t *testing.T) {
	// Arrange
	locations := testwrappers.SetupDefaultPathsUnix()
	verifier := &pathinsfmt.PathVerifier{
		UserGroupName: pathinsfmt.UserGroupName{},
		BaseRwxInstructions: chmodins.BaseRwxInstructions{
			RwxInstructions: []chmodins.RwxInstruction{
				{
					RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
					Condition: chmodins.Condition{
						IsSkipOnInvalid:   false,
						IsContinueOnError: false,
						IsRecursive:       false,
					},
				},
			},
		},
	}
	errCollection := errwrappers.Empty()

	// Act
	isSuccess := pathmodifierverify.ApplyVerifier(true,
		true,
		false,
		true,
		verifier,
		errCollection,
		locations)

	// Assert
	Convey("Default Paths Create", t, func() {
		So(errCollection.String(), ShouldBeEmpty)
		So(errCollection.IsEmpty(), ShouldBeTrue)
		So(isSuccess, ShouldBeTrue)
	})
}
