package main

import (
	"es-foxes/src"
	"es-foxes/structs"
	"fmt"
)

var Funs map[string]structs.FnString = map[string]structs.FnString{
	"backup":  src.StartBackup,
	"restore": src.StartRestore,
	"copy":    src.StartCopy,
	"clear":   src.ClearPath,
}

var FunsExplain map[string]string = map[string]string{
	"backup": "Backs up an ElasticSearch index to a TXT file located in the backup folder, each line being a record except for the first one which is the name of the index\n\n" +
		"\t\033[33mArguments:\033[0m\n" +
		"\t  - index_name String\n" +
		"\t  - url String [optional]\n\n" +
		"\t\033[33mExample:\033[0m\n" +
		"\t  ./es-foxes backup event-2021-01-01 http://localhost:9200",

	"restore": "Restores ElasticSearch records that are inside the TXT file in the backups folder, it is possible to rename the index passing by argument\n\n" +
		"\t\033[33mArguments:\033[0m\n" +
		"\t  - file_name String\n" +
		"\t  - index_to String\n" +
		"\t  - url String\n\n" +
		"\t\033[33mExample:\033[0m\n" +
		"\t  ./es-foxes restore event-2021-01-01.txt event-2021-01-01 http://localhost:9200",

	"copy": "It performs both the backup and the restore at the same time being the backup for the first ElasticSearch and the restore for the ElasticSearch being\n\n" +
		"\t\033[33mArguments:\033[0m\n" +
		"\t  - url_from String\n" +
		"\t  - index_from String\n\n" +
		"\t  - url_to String\n\n" +
		"\t  - index_to String\n\n" +
		"\t\033[33mExample:\033[0m\n" +
		"\t  ./es-foxes copy http://localhost:9200 event-2021-01-01 http://localhost:9400 event-2021-01-02",

	"clear": "Clean the backups folder following the rules inside foxes.yaml",
}

func ShowFuns() {
	fmt.Println()
	for k := range Funs {
		fmt.Println("\033[32m ♦︎  " + k + "\033[0m")
		fmt.Println("\t" + FunsExplain[k])
		fmt.Println()
	}
}
