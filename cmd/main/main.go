package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/cmd/config/datamodel"
	"gitlab.com/evatix-go/pathhelper/pathfilter"
)

func main() {
	// eep := pathhelper.GetExecutableEnvironmentPathCollection()
	// fmt.Println(eep)
	//
	// fmt.Println(pathhelper.GetWidowsDirectory())

	// collection := recursiveinternal.GetPaths("D:\\github\\Evatix\\text-replace-automation\\SampleFiles", 500, true)
	configPath := pathhelper.GetExecutableCombinePath("config.json")
	fmt.Println("Running : " + configPath)
	allBytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		return
	}

	var cliConfig datamodel.CliConfig
	json.Unmarshal(allBytes, &cliConfig)
	first := (cliConfig.CliRunner.FilesSelector)[0]
	query := pathfilter.NewQuery(
		&first.Filters,
		&first.Extensions)

	exceptQuery := pathfilter.NewQuery(
		&first.SkipFilters,
		&first.Extensions)

	pathTranspiler := corestr.NewHashmap(1)
	pathTranspiler.AddOrUpdate("workdir", "D:\\github\\Evatix\\text-replace-automation")

	collection := pathfilter.GetRecursiveExcept(
		osconsts.PathSeparator,
		true,
		pathTranspiler,
		first.Path,
		query,
		exceptQuery)

	collection.ErrorWrappers.HandleError()

	fmt.Println(strings.Join(*collection.Values, "\n\t"))
}
