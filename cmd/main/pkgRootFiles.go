package main

import (
	"gitlab.com/auk-go/pathhelper/pathsconst"
	"gitlab.com/auk-go/pathhelper/recursivepaths"
)

func pkgRootFiles() []string {
	return recursivepaths.Files(pathsconst.RootDir).Values
}
