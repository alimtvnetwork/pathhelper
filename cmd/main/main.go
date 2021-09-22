package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func main() {
	// downloadChecksumTest()
	// options := normalize.Options{
	// 	IsNormalize:        true,
	// 	IsLongPathFix:      true,
	// 	IsForceLongPathFix: true,
	// }
	//
	// samplePath := options.JoinWithBaseDirPaths(pathjoin.WithTemp(), "basedir", "something")
	// samplePath2 := options.JoinWithBaseDirPaths(samplePath, "basedir", "something")
	//
	// sha1 := hashas.Sha1
	//
	// rs := sha1.HexOfJsonResult(corejson.NewFromAny(samplePath2))
	//
	// fmt.Println(rs.Value)
	//
	// fmt.Println(samplePath)
	// fmt.Println(samplePath2)

	// sha1 := hashas.Sha1
	//
	// slice1 := []string{
	// 	"alim1",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// slice2 := []string{
	// 	"alim1",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// slice3 := []string{
	// 	"alim5",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// rs2 := sha1.HexSumOfAnys(slice1, slice2, slice3)
	//
	// fmt.Println(rs2.String())
	//
	// rs3 := sha1.HexSumOfAnysSingle(slice1, slice2, slice3)
	//
	// fmt.Println(rs3.String())
	//
	// files := recursivepaths.Files(pathsconst.RootDir)
	// fmt.Println("Files", files.String())
	//
	// rs5 := hexchecksum.OfFiles(&hexchecksum.FilesRequest{
	// 	Method:                     sha1,
	// 	IsGenerateContentsChecksum: true,
	// 	Files:                      files.Values,
	// })
	//
	// fmt.Println(converters.AnyToFullNameValueString(rs5))
	// fmt.Println(converters.AnyToFullNameValueString(rs5.CompileToSingle()))

	// checkSumCheck()

	fmt.Println("Hello World")
	a := "/tmp/dbapi/backup-storage//path-backup///alim-key1-dbapi/1/\\"
	b := "/dbmodel/\\webserverstoremodel/\\\\ServerWithSSL.go"
	joined3 := pathjoin.JoinConditionalNormalized3ExpandIf(
		true,
		false,
		a,
		b,
		"")
	fmt.Println(filepath.Clean(joined3))
	fmt.Println(filepath.Join(a, b))

	// joined := filepath.Join(a, b)
	fmt.Println(filepath.Clean(a))

	fmt.Println(normalize.Path("\\\\?\\tmp\\dbapi\\backup-storage\\path-backup\\alim-key1-dbapi\\1\\dbmodel\\webserverstoremodel\\ServerWithSSL.go"))
	fmt.Println(normalize.PathUsingSeparatorIf(true, true, true, osconsts.PathSeparator, "tmp\\dbapi\\backup-storage\\path-backup\\alim-key1-dbapi\\1\\dbmodel\\webserverstoremodel\\ServerWithSSL.go"))

	fmt.Println(normalize.Path("/home/a/../../git-repos"))
	fmt.Println(normalize.Path("\\home\\a\\../../git-repos"))

	// testPathWithVerifier()
}

func checkSumCheck() {
	result := hexchecksum.OfFilesContentsAsync(
		hashas.Sha256,
		"cmd/main/main.go")

	// result.ErrorWrapper.HandleError()
	result.ErrorWrapper.Log()
	fmt.Println(result.String())

	rs := pathsysinfo.GetPathUserGroupId("cmd/main")

	fmt.Println(converters.AnyToFullNameValueString(rs))
	fmt.Println(rs.Error)
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
