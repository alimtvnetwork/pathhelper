package pathhelper

import "io/ioutil"

// returns filepaths as []*string. non-lazy execution.
func GetFilesPaths(path string) []*string {
	var fileNames []*string

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		fileNames = append(fileNames, &fileName)
	}

	return fileNames
}
