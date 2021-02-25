package envpath

import normalize "gitlab.com/evatix-go/pathhelper/expandpath"

// todo change filename accordingly
func GetExecutableEnvironmentPathCollection() ExecutableEnvironmentPathCollection {
	rawPaths := GetRawExecutableEnvironmentPathCollection()

	pathsCollection := NewExecutableEnvironmentPathCollection(len(rawPaths))

	for _, rawPath := range rawPaths {
		expandedPath := normalize.EnvironmentVarExpand(rawPath)
		executableEnvironmentPath := ExecutableEnvironmentPath{
			Variable: rawPath,
			Expanded: expandedPath,
		}

		pathsCollection.AddPtr(&executableEnvironmentPath)
	}

	return pathsCollection
}
