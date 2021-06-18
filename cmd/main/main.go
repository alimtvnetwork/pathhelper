package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/pathstat"
)

func main() {
	// eep := pathhelper.GetExecutableEnvironmentPathCollection()
	// fmt.Println(eep)
	//
	// fmt.Println(pathhelper.GetWidowsDirectory())
	// arr := []string{
	// 	"/root/go/bin", "/temp/alimx",
	// }
	//
	// envPath := `PATH="/bin:/home/usrbin:/root/go/bin:/sbin:/snap/bin:/temp/alim:/usr/bin:/usr/games:/usr/local/bin:/usr/local/games:/usr/local/go/bin:/usr/local/sbin:/usr/sbin"`
	// fmt.Println(envpath.ReadEnvPaths())

	// fmt.Println(
	// 	pathchmod.ChmodApply(
	// 		true,
	// 		true,
	// 		chmodhelper.NewUsingFileModePtr(0644),
	// 		"/temp", "/temp/core"))

	fmt.Println(pathstat.Get("/etc"))
	//
	// 	contentLines := `
	// File: /etc/mysql
	//   Size: 4096            Blocks: 8          IO Block: 4096   directory
	// Device: 10302h/66306d   Inode: 6293381     Links: 4
	// Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
	// Access: 2021-06-03 09:18:30.014302041 +0000
	// Modify: 2021-06-03 09:18:32.086923601 +0000
	// Change: 2021-06-03 09:18:32.086923601 +0000
	// `
	// lines := strings.Split(contentLines, constants.NewLineUnix)
	// nonEmptyLines := stringslice.NonWhitespaceSlicePtr(
	// 	&lines)
	// pathInfo := pathstat.ProcessLinesToInfo(
	// 	*nonEmptyLines,
	// 	"/etc/",
	// 	errnew.EmptyPtr)
	//
	// fmt.Println(pathInfo)

	// fmt.Println("Remove", envpath.RemoveEnvPathsPtr(&arr))
	// // fmt.Println(envpath.LinuxAddOrUpdatePtr(&arr, true))
	// // fmt.Println(envpath.AddOrUpdateEnvPathsPtr(&arr))
	// fmt.Println("Read", "\n -\t"+strings.Join(envpath.ReadEnvPaths(), "\n -\t"))

	// fmt.Println(envpath.LinuxAddOrUpdatePtr(&arr, true))
	// fmt.Println(envpath.Linux(&arr, true))
	// collection := recursiveinternal.GetPaths("D:\\github\\Evatix\\text-replace-automation\\SampleFiles", 500, true)
	// configPath := pathhelper.GetExecutableCombinePath("config.json")
	// fmt.Println("Running : " + configPath)
	// allBytes, err := ioutil.ReadFile(configPath)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// var cliConfig pathinsfmt.CliConfig
	// json.Unmarshal(allBytes, &cliConfig)
	// first := (*cliConfig.CliRunner.FilesSelector)[0]
	// query := pathfilter.NewQuery(
	// 	first.Filters,
	// 	first.Extensions)
	//
	// // exceptQuery := pathfilter.NewQuery(
	// // 	&first.SkipFilters,
	// // 	&first.Extensions)
	//
	// pathTranspiler := corestr.NewHashmap(1)
	// pathTranspiler.
	// 	LinuxAddOrUpdate(
	// 		"workdir",
	// 		"D:\\github\\Evatix\\text-replace-automation")
	//
	// collection := pathfilter.GetRecursive(
	// 	osconsts.PathSeparator,
	// 	true,
	// 	pathTranspiler,
	// 	first.Path,
	// 	query)
	//
	// // collection := pathfilter.GetRecursive(
	// // 	osconsts.PathSeparator,
	// // 	true,
	// // 	pathTranspiler,
	// // 	first.Path,
	// // 	query)
	// //
	// // collection := pathfilter.Get(
	// // 	osconsts.PathSeparator,
	// // 	true,
	// // 	pathTranspiler,
	// // 	first.Path,
	// // 	query)
	//
	// collection.ErrorWrappers.HandleError()
	//
	// fmt.Println(strings.Join(*collection.Values, "\n\t"))
	//
	// fmt.Println(dirinfo.New("c:\\windows\\py.exe").IsValidDir)
	//
	// fmt.Println()
	// fmt.Println(
	// 	unipaths.
	// 		New("\\").
	// 		AddPaths("c://windows", "d:\\maindrive\\something/g.go", "hjello/eee").
	// 		StringsPtr())
	//
	// wrappers := unipath.
	// 	New("\\").
	// 	Add("D:\\github\\Evatix\\text-replace-automation\\SampleFiles").
	// 	GetFileInfoWrappers()
	//
	// json2 := wrappers.Json()
	// emptyWrappers := fileinfo.EmptyWrappers()
	//
	// fmt.Println(json2.JsonString())
	// emptyWrappers.ParseInjectUsingJson(json2)
	//
	// fmt.Println(emptyWrappers.PathsCollection().Json().JsonString())
}
