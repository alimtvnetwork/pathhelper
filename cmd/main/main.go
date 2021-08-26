package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
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
	// CopierTest()
	//
	// instruction := pathrecurseinfo.Instruction{
	// 	Root:                   pathsconst.RootDir,
	// 	ExcludingRootNames:     []string{".git"},
	// 	ExcludingPaths:         []string{"D:\\others-git\\gitlabs\\pathhelper\\apachelinuxpath"},
	// 	IsIncludeFilesOnly:     false,
	// 	IsRelativePath:         true,
	// 	IsIncludeDirsOnly:      true,
	// 	IsIncludeAll:           false,
	// 	IsExcludeRoot:          false,
	// 	IsRecursive:            true,
	// 	IsExpandEnvironmentVar: false,
	// 	IsNormalize:            false,
	// }
	//
	// result := instruction.Result()
	// // fmt.Println(result.PathsString())
	// fmt.Println(result.PathsResult.JoinWithRoot(true, false, "d:\\a//").String())
	// fmt.Println(pathhelper.GetLocationInfo("ab/c\\d.tx").String())
	// fmt.Println(pathhelper.GetLocationInfo("ab/c/d").String())
	// fmt.Println(pathhelper.GetLocationInfo("d.tx").String())
	// //
	// collection := elitepath.NewPathCollectionDirect(nil, pathsconst.RootDir)
	// fmt.Println(collection.First().BothExt())

	// filter := &elitepath.Filter{
	// 	// PathFilter: &elitepath.ValueFilter{
	// 	// 	Value:           "checksummer",
	// 	// 	IsCaseSensitive: false,
	// 	// 	Compare:         stringcompareas.EndsWith,
	// 	// },
	// 	ExistFilter: &elitepath.ExistFilter{
	// 		HasSafeItems: true,
	// 		IsDir:   true,
	// 	},
	// 	// NameRegexFilter: "",
	// 	// PathRegexFilter: `\\.+checksummer$`,
	// }
	// fmt.Println(collection.Skip(2).Take(5).FilterPathCollection(filter).String())
	// DownloadTest()

	// downloadChecksumTest()

	testPathWithVerifier()
}

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

func downloadChecksumTest() {
	download := &pathinsfmt.Download{
		// Url:                  "https://github.com/aria2/aria2/releases/download/release-1.35.0/aria2-1.35.0.tar.xz",
		Url:                  "https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh",
		Destination:          "/home/a/checksum2",
		FileName:             "install.sh",
		IsSkipOnExist:        false,
		IsCreateDir:          true,
		FileModeDir:          consts.DefaultDirectoryFileMode,
		ChecksumVerifyMethod: hashas.Sha256,
		// ChecksumVerify:       "1e2b7fd08d6af228856e51c07173cfcf987528f1ac97e04c5af4a47642617dfd",
		ChecksumVerify: "b6af836b2662f21081091e0bd851d92b2507abb94ece340b663db7e4019f8c7c",
	}

	errW := downloadinsexec.Apply(download)
	errW.HandleError()
}

func DownloadTest() {
	ins := &pathinsfmt.Download{
		Url:              "https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh",
		Destination:      "/home/a/dtestxxxx",
		FileName:         "installx.sh",
		IsCreateDir:      true,
		IsClearDir:       true,
		ParallelRequests: 4,
		MaxRetries:       5,
		FileModeDir:      filemode.X666,
	}

	errW := downloadinsexec.Apply(ins)
	fmt.Println(errW)
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
