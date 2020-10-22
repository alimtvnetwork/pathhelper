package pathhelper

// todo change filename accordingly
func GetExecutableEnvironmentPathCollection() []ExecutableEnvironmentPath {
	rawPaths := GetRawExecutableEnvironmentPathCollection()

	var outputs = []ExecutableEnvironmentPath{}

	for _, rawPath := range rawPaths {
		expandedPath := PathFromEnvVariable(rawPath)
		outputPerPath := ExecutableEnvironmentPath{
			Variable: rawPath,
			Expanded: expandedPath,
		}

		outputs = append(outputs, outputPerPath)
	}

	return outputs
}
