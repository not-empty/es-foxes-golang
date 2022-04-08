package src

import (
	"bufio"
	"context"
	"encoding/json"
	"es-foxes/es"
	"es-foxes/helpers"
	"es-foxes/structs"
	"log"
	"os"
	"path/filepath"

	"github.com/olivere/elastic/v7"
)

func StartRestore(args ...string) {
	log.Println("Start restore")
	println()

	if len(args) < 5 {
		log.Fatalln(
			"Restore Arguments:" +
				"\n\t- file_name String" +
				"\n\t- index_to String" +
				"\n\t- url String",
		)
	}

	url := args[4]
	index_name := args[3]

	restore(args[2], url, index_name)

	println()
	log.Println("Finished restore")
}

func restore(filename string, url string, index_name string) {
	file := helpers.OpenFileRead(filepath.Join("backup", filename))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var index structs.IndexName

	if index_name != "" {
		scanner.Scan()
		index = getEventIndexFromParam(index_name)
	} else {
		index = getEventIndex(file, scanner)
	}

	insertOnES(file, scanner, index, url)
}

func getEventIndex(file *os.File, scanner *bufio.Scanner) structs.IndexName {
	scanner.Scan()

	var index_name structs.IndexName

	err := json.Unmarshal([]byte(scanner.Text()), &index_name)

	if err != nil {
		log.Fatalln("Index format error")
	}

	return index_name
}

func getEventIndexFromParam(index_name_param string) structs.IndexName {
	index_name := structs.IndexName{
		Index: index_name_param,
	}

	return index_name
}

func getEvent(eventJson string) (structs.Event, error) {
	var event structs.Event

	err := json.Unmarshal([]byte(eventJson), &event)

	if err != nil {
		return structs.Event{}, err
	}

	return event, nil
}

func insertOnES(file *os.File, scanner *bufio.Scanner, index structs.IndexName, url string) {
	client, err := es.GetESClient(url)

	if err != nil {
		log.Fatalln("Client ES not connected")
	}

	line := 0
	processor := getBulkProcessor(client)

	defer processor.Close()

	for scanner.Scan() {
		line = line + 1
		event, err := getEvent(scanner.Text())

		if err != nil {
			log.Println("Evento can not be parsed on line", line)
			continue
		}

		if line%1000 == 0 {
			log.Println(line, "Events Inserts")
		}

		request := elastic.NewBulkIndexRequest()
		request = request.Index(index.Index)
		request = request.Type("_doc")
		request = request.Id(event.Id)
		request = request.Doc(event.Source)

		processor.Add(request)
	}

	processor.Flush()
	log.Println(line, "Events Inserts")
}

func afterFunction(executionId int64, requests []elastic.BulkableRequest, response *elastic.BulkResponse, err error) {
	if response != nil && response.Errors {
		file := helpers.OpenFileWrite(filepath.Join("backup", "errors.txt"))
		defer file.Close()

		log.Println("Errors to send event to ES, please see on backup/errors.txt")

		for _, i := range requests {
			s, err := i.Source()

			if err != nil {
				continue
			}

			out, err := json.Marshal(s)

			if err != nil {
				continue
			}

			helpers.AddRowToFile(file, string(out))
		}

		for _, erro := range response.Failed() {
			log.Printf("Error %v", erro.Error)
		}
	}
}

func getBulkProcessor(client *elastic.Client) *elastic.BulkProcessor {
	bulk := client.BulkProcessor()
	bulk = bulk.Name("ES-FOXES")
	bulk = bulk.Workers(structs.Variables.BulkWorkers)
	bulk = bulk.Stats(true)
	bulk = bulk.BulkSize(structs.Variables.BulkSizeMB << 20)
	bulk = bulk.After(afterFunction)

	processor, err := bulk.Do(context.Background())

	if err != nil {
		log.Fatalln("Error to create bulker")
	}

	return processor
}
