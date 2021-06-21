package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

func main() {
	dir := nginxlinuxpath.GetFullDirStructure(true, "d:\\sample-nginx")
	dir.MkDirAll(filemode.X644)

	fmt.Println(recursiveinternal.GetPaths(osconsts.PathSeparator, dir.Root, true))
	// fmt.Println(dir.AllFilesInSitesEnabled())
}
