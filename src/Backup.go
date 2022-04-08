package src

import (
	"context"
	"encoding/json"
	"es-foxes/es"
	"es-foxes/helpers"
	"es-foxes/structs"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/oklog/ulid"
	"github.com/olivere/elastic/v7"
)

func StartBackup(args ...string) {
	log.Println("Start backup")
	println()

	if len(args) < 3 {
		log.Fatalln(
			"Backup Arguments:" +
				"\n\t- index_name String" +
				"\n\t- url String",
		)
	}

	url := helpers.GetArg(3, args, structs.Variables.ElasticUrl)

	backup(args[2], url)

	println()
	log.Println("Finished backup")
	println()
}

func backup(index_name string, url string) []string {
	q := 0
	next_print := 1000
	next_file := structs.Variables.EventsPerFile
	scroller := getScroller(index_name, url)
	name_list := []string{}
	event_index_bytes := getEventIndexBytes(index_name)

	file := createNewFile(index_name, &name_list, event_index_bytes)

	for {
		response, err := scroller.Do(context.Background())

		if err == io.EOF {
			log.Println(response.Hits.TotalHits.Value, "Events processed")
			file.Close()
			break
		}

		if err != nil {
			log.Fatalln("Error to get events from elastic")
			file.Close()
			break
		}

		q = q + len(response.Hits.Hits)

		if q >= next_print {
			log.Println(q, "Events processed")
			next_print = next_print + 1000
		}

		addHitsToFile(file, response.Hits.Hits)

		if q >= next_file {
			next_file = next_file + structs.Variables.EventsPerFile
			file.Close()
			file = createNewFile(index_name, &name_list, event_index_bytes)
		}
	}

	println()
	log.Println("Finished compression")

	return name_list
}

func createNewFile(index_name string, name_list *[]string, event_index_bytes []byte) *os.File {
	name := genBackupName(index_name)
	*name_list = append(*name_list, name)

	file := helpers.OpenFileWrite(filepath.Join("backup", name))
	helpers.AddRowToFile(file, string(event_index_bytes))

	return file
}

func getScroller(index_name string, url string) *elastic.ScrollService {
	client, err := es.GetESClient(url)

	if err != nil {
		log.Fatalln("Client ES not connected")
	}

	scroller := client.Scroll(index_name)
	scroller = scroller.Scroll("1m")
	scroller = scroller.Size(structs.Variables.EventsPerPage)

	return scroller
}

func getEventIndexBytes(index_name string) []byte {
	event_index := structs.IndexName{
		Index: index_name,
	}

	event_index_bytes, err := json.Marshal(event_index)

	if err != nil {
		log.Fatal("Error to parse index name")
	}

	return event_index_bytes
}

func genBackupName(index_name string) string {
	now := time.Now()

	t := time.Unix(1000000, 0)
	ulidz := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	return index_name + "_" + now.Format("20060201_150405") + "_" + ulid.MustNew(ulid.Timestamp(now), ulidz).String() + ".txt"
}

func addHitsToFile(file *os.File, hits []*elastic.SearchHit) {
	for _, i := range hits {
		var js map[string]interface{}

		outjson, err := i.Source.MarshalJSON()

		if err != nil {
			log.Fatalln("Error to parse json event")
		}

		json.Unmarshal(outjson, &js)
		js["id"] = i.Id

		event := structs.Event{
			Id:     i.Id,
			Source: js,
		}

		out, err := json.Marshal(event)

		if err != nil {
			log.Fatalln("Error to parse event")
		}

		helpers.AddRowToFile(file, string(out))
	}
}
