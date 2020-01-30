package data

import (
	"io/ioutil"
	"path/filepath"
)

// fetchFilesFromDir returns a map of all filenames in a directory,
// e.g map{"BAS.DE": "BAS.DE.csv"}.
func fetchFilesFromDir(dir string) (m map[string]string, err error) {
	// read filenames from directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return m, err
	}

	// initialise the map
	m = make(map[string]string)

	// read filenames from directory
	for _, file := range files {
		// file is directory
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		extension := filepath.Ext(filename)
		// file is not CSV
		if extension != ".csv" {
			continue
		}

		name := filename[0 : len(filename)-len(extension)]
		m[name] = filename
	}
	return m, nil
}
