package src

import (
	"es-foxes/helpers"
	"es-foxes/structs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func StartCopy(args ...string) {
	log.Println("Start copy")
	println()

	if len(args) < 6 {
		log.Fatalln(
			"Copy Arguments:" +
				"\n\t- url_from String" +
				"\n\t- index_from String" +
				"\n\t- url_to String" +
				"\n\t- index_to String",
		)
		log.Fatalln("Backup args:\n\t - url_from String\n\t - index_from String\n\t - url_to String\n\t - index_to String")
	}

	url_from := helpers.GetArg(2, args, structs.Variables.ElasticUrl)
	index_from := helpers.GetArg(3, args, "")

	url_to := helpers.GetArg(4, args, structs.Variables.ElasticUrl)
	index_to := helpers.GetArg(5, args, "")

	copy(url_from, index_from, url_to, index_to)

	println()
	log.Println("Finished copy")
	println()
}

func copy(url_from string, index_from string, url_to string, index_to string) {
	list_of_files := backup(index_from, url_from)

	for index, i := range list_of_files {
		println()
		log.Println("Processing file " + i + " [" + strconv.Itoa(index+1) + "/" + strconv.Itoa(len(list_of_files)) + "]")

		restore(i, url_to, index_to)

		os.Remove(filepath.Join(helpers.GetPath(), "backup", i))
	}
}
