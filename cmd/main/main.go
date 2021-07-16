package main

import (
	"fmt"
	"io/ioutil"

	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func main() {
	// verifiers := pathinsfmt.PathVerifiers{
	// 	BaseSpecPlusRequestIds: coreinstruction.BaseSpecPlusRequestIds{},
	// 	PathVerifiers: []pathinsfmt.PathVerifier{
	// 		{
	// 			UserGroupName: *pathinsfmt.NewUserGroupName(
	// 				"alim", ""),
	// 			BaseRwxInstructions: chmodins.BaseRwxInstructions{
	// 				RwxInstructions: []chmodins.RwxInstruction{
	// 					{
	// 						RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
	// 							Owner: "rwx",
	// 							Group: "rw-",
	// 							Other: "rw-",
	// 						},
	// 						Condition: chmodins.Condition{},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	IsSkipCheckingOnInvalid: false,
	// 	IsNormalize:             false,
	// 	IsRecursiveCheck:        false,
	// }
	//
	// locations := []string{
	// 	os.TempDir(),
	// }
	//
	// errorCollection := errwrappers.Empty()
	// _ = pathmodifierverify.ApplyVerifier(
	// 	true,
	// 	true,
	// 	true,
	// 	true,
	// 	&verifiers.PathVerifiers[0],
	// 	errorCollection,
	// 	locations)
	//
	// wkPath := os.TempDir() + "/main.json"
	// errCollection2 := &errwrappers.Collection{}
	//
	// readWp := fs.
	// 	ReadJsonParseSelfInjector(wkPath, errCollection2)
	// fmt.Println(readWp)
	//
	// fmt.Println(errCollection2.IsSuccess(), errCollection2)
	//
	// fmt.Println(fs.CopyFileContents(wkPath, wkPath+"2.json"))
	//
	// pathFinal := pathjoin.WithTempPlusDefaults(
	// 	"alim",
	// 	"loca1",
	// 	"loc2")
	//
	// fmt.Println(pathFinal)
	//
	// // wr := fs.
	// // 	WriteJsonResult(false, errorCollection.Json(),wkPath)
	//
	// // fmt.Println(wr)
	CopierTest()
	// CopierTest2()

}

func CopierTest() {
	tmpDir, _ := ioutil.TempDir(pathsconst.DefaultTempTestDir, "ttt")
	tmpDir = tmpDir + "/a/something"
	fmt.Println("to", tmpDir)
	fmt.Println("from", pathsconst.RootDir)
	errW := copyrecursive.NewCopier(
		pathsconst.RootDir, tmpDir, copyrecursive.Options{
			IsSkipOnExist:      false,
			IsRecursive:        true,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        true,
		},
	).Copy()

	errW.HandleError()
}

func CopierTest2() {
	tmpDir, _ := ioutil.TempDir("", "ttt")

	fmt.Println(pathsconst.RootDir)
	errW := copyrecursive.Do(
		false,
		pathsconst.RootDir,
		tmpDir)

	errW.HandleError()
}
