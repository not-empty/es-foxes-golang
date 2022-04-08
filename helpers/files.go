package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func GetPath() string {
	ex, err := os.Executable()

	if err != nil {
		return ""
	}

	path := filepath.Dir(ex)

	return path
}

func OpenFileRead(filename string) *os.File {
	path := filepath.Join(GetPath(), filename)

	file, err := os.Open(path)

	if err != nil {
		log.Fatalln("File " + path + " not found")
	}

	return file
}

func OpenFileWrite(filename string) *os.File {
	path := filepath.Join(GetPath(), filename)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalln("File " + path + " not found")
	}

	return file
}

func AddRowToFile(file *os.File, row string) {
	_, err := file.WriteString(row + "\n")

	if err != nil {
		log.Fatalln("Erro on write file")
	}
}
