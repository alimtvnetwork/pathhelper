package expandpath

import "gitlab.com/evatix-go/core/coreutils/stringutil"

func GetCompiledPath(
	pathTemplate string,
	compilingMap map[string]string,
) string {
	return stringutil.ReplaceByMap(
		pathTemplate,
		compilingMap)
}
