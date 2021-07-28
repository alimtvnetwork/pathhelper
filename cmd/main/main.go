package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/hashas"
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
	// CopierTest()
	// CopierTest2()

	// TestHashSumSync()
	// TestHashSumAsync()
	CopierTest()
}

func CopierTest() {
	tmpDir, _ := ioutil.TempDir("", "ttt")
	_, b, _, _ := runtime.Caller(0)
	srcDir := filepath.Join(filepath.Dir(b), "..", "..")
	fmt.Println("to", tmpDir)
	fmt.Println("from", srcDir)
	errW := copyrecursive.NewCopier(
		srcDir, tmpDir, copyrecursive.Options{
			IsSkipOnExist:      false,
			IsRecursive:        false,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        true,
		},
	).Copy()

	errW.HandleError()
}

const v = false

func prettyPrint(m map[string][]byte) {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, key := range names {
		fmt.Printf("%x: %s\n", m[key], key)
	}
}

func TestHashSumSync() {
	start := time.Now()
	c := checksummer.NewSync(true, "D:\\vm", hashas.Md5)
	elapsed := time.Since(start)
	if v {
		prettyPrint(c.GetMap())
	}

	fmt.Printf("Elapsed Sync: %s\n", elapsed)
}

func TestHashSumAsync() {
	start := time.Now()
	c := checksummer.NewAsync(true, "D:\\vm", hashas.Md5)
	elapsed := time.Since(start)
	if v {
		prettyPrint(c.GetMap())
	}

	fmt.Printf("Elapsed Async: %s\n", elapsed)
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
