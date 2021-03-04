package envpath

type ExecutableEnvironmentPathCollection struct {
	pathsMap *map[string]ExecutableEnvironmentPath
	paths    *[]ExecutableEnvironmentPath
}

func NewExecutableEnvironmentPathCollection(capacity int) ExecutableEnvironmentPathCollection {
	pathsMap := make(map[string]ExecutableEnvironmentPath, capacity)
	paths := make([]ExecutableEnvironmentPath, 0, capacity)

	return ExecutableEnvironmentPathCollection{
		pathsMap: &pathsMap,
		paths:    &paths,
	}
}

func NewExecutableEnvironmentPathCollectionPtr(capacity int) *ExecutableEnvironmentPathCollection {
	pathsMap := make(map[string]ExecutableEnvironmentPath, capacity)
	paths := make([]ExecutableEnvironmentPath, 0, capacity)

	return &ExecutableEnvironmentPathCollection{
		pathsMap: &pathsMap,
		paths:    &paths,
	}
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) Add(exeEnvPath ExecutableEnvironmentPath) {
	(*executableEnvironmentPathCollection.pathsMap)[exeEnvPath.Variable] = exeEnvPath
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) AddPtr(exeEnvPath *ExecutableEnvironmentPath) {
	if exeEnvPath != nil {
		(*executableEnvironmentPathCollection.pathsMap)[exeEnvPath.Variable] = *exeEnvPath
	}
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) IsExistPtr(exeEnvPath *ExecutableEnvironmentPath) bool {
	_, has := (*executableEnvironmentPathCollection.pathsMap)[exeEnvPath.Variable]

	return has
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) IsExist(exeEnvPath ExecutableEnvironmentPath) bool {
	_, has := (*executableEnvironmentPathCollection.pathsMap)[exeEnvPath.Variable]

	return has
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) List() *[]ExecutableEnvironmentPath {
	return executableEnvironmentPathCollection.paths
}

func (executableEnvironmentPathCollection *ExecutableEnvironmentPathCollection) OnlyNamesCollection() *[]ExecutableEnvironmentPath {
	return executableEnvironmentPathCollection.paths
}
