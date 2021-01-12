package pathhelper

// todo change filename accordingly
func GetExecutableEnvironmentPathCollection() ExecutableEnvironmentPathCollection {
	rawPaths := GetRawExecutableEnvironmentPathCollection()

	pathsCollection := NewExecutableEnvironmentPathCollection(len(rawPaths))

	for _, rawPath := range rawPaths {
		expandedPath := PathFromEnvVariable(rawPath)
		executableEnvironmentPath := ExecutableEnvironmentPath{
			Variable: rawPath,
			Expanded: expandedPath,
		}

		pathsCollection.AddPtr(&executableEnvironmentPath)
	}

	return pathsCollection
}
