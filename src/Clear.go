package src

import (
	"es-foxes/helpers"
	"es-foxes/structs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ClearPath(_ ...string) {
	clear("backup")
}

func clear(dir string) {
	files, err := ioutil.ReadDir(filepath.Join(helpers.GetPath(), dir))

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if !validadeFile(f.Name()) {
			continue
		}

		os.Remove(filepath.Join(helpers.GetPath(), dir, f.Name()))
		log.Println(filepath.Join(dir, f.Name()), "Removed")
	}
}

func validadeFile(filename string) bool {
	filename_split := strings.Split(filename, ".")
	ext := filename_split[len(filename_split)-1]

	for _, i := range structs.Variables.FoxesConfig.Clear.IgnoreExtensions {
		if i == strings.ToLower(ext) {
			return false
		}
	}

	for _, i := range structs.Variables.FoxesConfig.Clear.IgnoreFiles {
		if i == strings.ToLower(filename) {
			return false
		}
	}

	return true
}
