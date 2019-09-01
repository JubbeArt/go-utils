package files

import (
	"io/ioutil"
	"path/filepath"
)

func Files(folder string) []string {
	files := []string{}

	fileObjects, _ := ioutil.ReadDir(folder)

	for _, fileInfo := range fileObjects {
		files = append(files, filepath.Join(folder, fileInfo.Name()))
	}

	return files
}
